package client

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andybalholm/brotli"
	httpi "github.com/hopeio/utils/net/http"
	url2 "github.com/hopeio/utils/net/url"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/hopeio/utils/strings/ascii"
	"github.com/hopeio/utils/strings/unicode"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	DefaultClient = New().DisableLog()
	bufPool       = &sync.Pool{New: func() interface{} { return &bytes.Buffer{} }}
)

type Request struct {
	ctx         context.Context
	Method, Url string
	contentType ContentType
	header      httpi.Header //请求级请求头
	client      *Client
}

func NewRequest(method, url string) *Request {
	return &Request{
		ctx:    context.Background(),
		Method: method,
		Url:    url,
		client: DefaultClient,
	}
}

func (req *Request) WithClient(c *Client) *Request {
	req.client = c
	return req
}

func (req *Request) SetClient(set func(c *Client)) *Request {
	req.client = New()
	set(req.client)
	return req
}

func (req *Request) Header(header httpi.Header) *Request {
	req.header = header
	return req
}

func (req *Request) AddHeader(k, v string) *Request {
	req.header.Set(k, v)
	return req
}

func (req *Request) ContentType(contentType ContentType) *Request {
	req.contentType = contentType
	return req
}

func (req *Request) Context(ctx context.Context) *Request {
	req.ctx = ctx
	return req
}

func (req *Request) DoEmpty() error {
	return req.Do(nil, nil)
}

func (req *Request) DoNoParam(response any) error {
	return req.Do(nil, response)
}

func (req *Request) DoNoResponse(param any) error {
	return req.Do(param, nil)
}

func (req *Request) DoRaw(param any) (RawBytes, error) {
	var raw RawBytes
	err := req.Do(param, &raw)
	if err != nil {
		return raw, err
	}
	return raw, nil
}

func (req *Request) DoStream(param any) (io.ReadCloser, error) {
	var resp *http.Response
	err := req.Do(param, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (req *Request) addHeader(request *http.Request, c *Client) string {
	var auth string
	nv := 0
	for _, vv := range c.header {
		nv += len(vv)
	}
	sv := make([]string, nv) // shared backing array for header' values

	for k, vv := range c.header {
		if k == httpi.HeaderAuthorization {
			auth = vv[0]
		}
		if vv == nil {
			continue
		}
		n := copy(sv, vv)
		request.Header[k] = sv[:n:n]
		sv = sv[n:]
	}

	for i := 0; i+1 < len(req.header); i += 2 {
		request.Header.Set(req.header[i], req.header[i+1])
		if req.header[i] == httpi.HeaderAuthorization {
			auth = req.header[i+1]
		}
	}

	request.Header.Set(httpi.HeaderContentType, req.contentType.String())
	return auth
}

// Do create a HTTP request
// param: 请求参数 目前只支持编码为json 或 Url-encoded
func (req *Request) Do(param, response any, opts ...RequestOption) error {
	for _, opt := range opts {
		opt(req)
	}
	if req.Method == "" {
		return errors.New("not set method")
	}

	if req.Url == "" {
		return errors.New("not set url")
	}
	if req.ctx == nil {
		req.ctx = context.Background()
	}
	if req.client == nil {
		req.client = DefaultClient
	}
	c := req.client

	var reqBody, respBody *Body
	var statusCode, reqTimes int
	var err error
	var auth string
	reqTime := time.Now()
	// 日志记录
	defer func(now time.Time) {
		if c.logLevel == LogLevelInfo || (err != nil && c.logLevel == LogLevelError) {
			c.logger(req.Method, req.Url, auth, reqBody, respBody, statusCode, time.Since(now), err)
		}
	}(reqTime)

	if req.Method == http.MethodGet {
		req.Url = url2.AppendQueryParam(req.Url, param)
	} else {
		reqBody = &Body{}
		if param != nil {
			switch paramType := param.(type) {
			case string:
				reqBody.Data = stringsi.ToBytes(paramType)
			case []byte:
				reqBody.Data = paramType
			case io.Reader:
				var reqBytes []byte
				reqBytes, err = io.ReadAll(paramType)
				reqBody.Data = reqBytes
			default:
				switch req.contentType {
				case ContentTypeForm:
					params := url2.QueryParam(param)
					reqBody.Data = stringsi.ToBytes(params)
				case ContentTypeXml:

				default:
					var reqBytes []byte
					reqBytes, err = json.Marshal(param)
					if err != nil {
						return err
					}
					reqBody.Data = reqBytes
					reqBody.ContentType = ContentTypeJson
				}
			}
		}
	}
	var request *http.Request
	reqBytes := reqBody.Data
	if c.reqDataHandler != nil {
		reqBytes, err = c.reqDataHandler(reqBody.Data)
	}
	request, err = http.NewRequestWithContext(req.ctx, req.Method, req.Url, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}

	auth = req.addHeader(request, c)

	var resp *http.Response
Retry:
	if reqTimes > 0 {
		if c.retryInterval != 0 {
			time.Sleep(c.retryInterval)
		}
		reqTime = time.Now()
		if reqBytes != nil {
			request.Body = io.NopCloser(bytes.NewReader(reqBytes))
		}
		if c.retryHandler != nil {
			c.retryHandler(request)
		}
	}
	resp, err = c.httpClient.Do(request)
	reqTimes++
	if err != nil {
		if c.retryTimes == 0 || reqTimes == c.retryTimes {
			return err
		} else {
			if c.logLevel > LogLevelSilent {
				c.logger(req.Method, req.Url, auth, reqBody, respBody, statusCode, time.Since(reqTime), errors.New(err.Error()+";will retry"))
			}
			goto Retry
		}
	}

	respBody = &Body{}
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		respBody.ContentType = ContentTypeText
		if resp.StatusCode == http.StatusNotFound {
			err = ErrNotFound
		} else {
			var msg []byte
			msg, err = io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return err
			}
			err = errors.New("status:" + resp.Status + " " + unicode.ToUtf8(msg))
		}
		return err
	}

	if httpresp, ok := response.(*http.Response); ok {
		*httpresp = *resp
		return err
	}

	if httpresp, ok := response.(**http.Response); ok {
		*httpresp = resp
		return err
	}

	var reader io.Reader
	// net/http会自动处理gzip
	// go1.22 发现没有处理
	if resp.Header.Get(httpi.HeaderContentEncoding) == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		resp.Header.Del("Content-Encoding")
		resp.Header.Del("Content-Length")
		resp.ContentLength = -1
		resp.Uncompressed = true
		if err != nil {
			resp.Body.Close()
			return err
		}
	} else if ascii.EqualFold(resp.Header.Get("Content-Encoding"), "br") {
		reader = brotli.NewReader(resp.Body)
		resp.Header.Del("Content-Encoding")
		resp.Header.Del("Content-Length")
		resp.ContentLength = -1
		resp.Uncompressed = true
	} else if ascii.EqualFold(resp.Header.Get("Content-Encoding"), "deflate") {
		reader = flate.NewReader(resp.Body)
		resp.Header.Del("Content-Encoding")
		resp.Header.Del("Content-Length")
		resp.ContentLength = -1
		resp.Uncompressed = true
	} else {
		reader = resp.Body
	}

	if httpresp, ok := response.(*io.Reader); ok {
		*httpresp = reader
		return err
	}
	statusCode = resp.StatusCode

	var respBytes []byte
	if c.responseHandler != nil {
		var retry bool
		retry, respBytes, err = c.responseHandler(resp)
		resp.Body.Close()

		if retry {
			if c.logLevel > LogLevelSilent {
				c.logger(req.Method, req.Url, auth, reqBody, respBody, statusCode, time.Since(reqTime), err)
			}
			goto Retry
		} else if err != nil {
			return err
		}
	} else {
		respBytes, err = io.ReadAll(reader)
		resp.Body.Close()
		if err != nil {
			return err
		}
	}
	respBody.Data = respBytes
	if len(respBytes) > 0 && response != nil {
		contentType := resp.Header.Get(httpi.HeaderContentType)
		respBody.ContentType.Decode(contentType)

		if raw, ok := response.(*RawBytes); ok {
			*raw = respBytes
			return nil
		}
		if respBody.ContentType == ContentTypeForm {
			// TODO
		} else {
			// 默认json
			err = json.Unmarshal(respBytes, response)
			if err != nil {
				return fmt.Errorf("json.Unmarshal error: %v", err)
			}
		}

		if v, ok := response.(ResponseBodyCheck); ok {
			err = v.CheckError()
		}
	}

	return err
}
