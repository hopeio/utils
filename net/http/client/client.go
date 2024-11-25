/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	httpi "github.com/hopeio/utils/net/http"
	"io"
	"net"
	"net/http"
	stdurl "net/url"
	"time"
)

// github.com/go-resty/resty 是个不错的选择,但是缺少一些我需要的功能，例如brotli解码，以及自定义处理body data，用于解决一些参数和返回body的AES加密或其他
// 不是并发安全的

var (
	DefaultHttpClient = newHttpClient(ClientTypeApi)
	DefaultLogLevel   = LogLevelError
)

const timeout = time.Minute

type ClientType uint8

const (
	ClientTypeApi      ClientType = iota
	ClientTypeDownload ClientType = iota
	ClientTypeUpload   ClientType = iota
)

func newHttpClient(typ ClientType) *http.Client {
	if typ == ClientTypeApi {
		return &http.Client{
			//Timeout: timeout * 2,
			Transport: &http.Transport{
				Proxy:             http.ProxyFromEnvironment, // 代理使用
				ForceAttemptHTTP2: true,
				DialContext: (&net.Dialer{
					Timeout:   timeout,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				//DisableKeepAlives: true,
				TLSHandshakeTimeout: timeout,
			},
		}
	}
	return newDownloadHttpClient()
}

// Client ...
type Client struct {
	typ ClientType
	// httpClient settings
	httpClient    *http.Client
	newHttpClient bool

	parseTag string // 默认json

	// request
	httpRequestOptions []HttpRequestOption
	header             http.Header //公共请求头
	reqDataHandler     func(data []byte) ([]byte, error)

	// response
	responseHandler func(response *http.Response) (retry bool, reader io.Reader, err error)
	resDataHandler  func(data []byte) ([]byte, error)

	// logger
	logger   AccessLog
	logLevel LogLevel

	// retry
	retryTimes    int
	retryInterval time.Duration
	retryHandler  func(*http.Request)
}

func New() *Client {
	return &Client{httpClient: DefaultHttpClient, logger: DefaultLogger, logLevel: DefaultLogLevel, retryInterval: 200 * time.Millisecond}
}

func (c *Client) Header(header httpi.Header) *Client {
	if c.header == nil {
		c.header = make(http.Header)
	}
	header.IntoHttpHeader(c.header)
	return c
}

func (c *Client) AddHeader(k, v string) *Client {
	if c.header == nil {
		c.header = make(http.Header)
	}
	c.header.Add(k, v)
	return c
}

func (c *Client) Logger(logger AccessLog) *Client {
	if logger == nil {
		return c
	}
	c.logger = logger
	return c
}

func (c *Client) DisableLog() *Client {
	c.logLevel = LogLevelSilent
	return c
}

func (c *Client) LogLevel(lvl LogLevel) *Client {
	c.logLevel = lvl
	return c
}

func (c *Client) ParseTag(tag string) *Client {
	c.parseTag = tag
	return c
}

// handler 返回值:是否重试,返回数据,错误
func (c *Client) ResponseHandler(handler func(response *http.Response) (retry bool, reader io.Reader, err error)) *Client {
	c.responseHandler = handler
	return c
}

func (c *Client) ResDataHandler(handler func(data []byte) ([]byte, error)) *Client {
	c.resDataHandler = handler
	return c
}

func (c *Client) ReqDataHandler(handler func(data []byte) ([]byte, error)) *Client {
	c.reqDataHandler = handler
	return c
}

func (c *Client) HttpRequestOption(opts ...HttpRequestOption) *Downloader {
	c.httpRequestOptions = append(c.httpRequestOptions, opts...)
	return c
}

// 设置过期时间,仅对单次请求有效
func (c *Client) Timeout(timeout time.Duration) *Client {
	if !c.newHttpClient {
		c.httpClient = newHttpClient(c.typ)
		c.newHttpClient = true
	}
	setTimeout(c.httpClient, timeout)
	return c
}

func (c *Client) HttpClient(client *http.Client) *Client {
	c.httpClient = client
	c.newHttpClient = true
	return c
}

func (c *Client) SetHttpClient(opt HttpClientOption) *Client {
	if !c.newHttpClient {
		c.httpClient = newHttpClient(c.typ)
		c.newHttpClient = true
	}
	opt(c.httpClient)
	return c
}

func (c *Client) RetryTimes(retryTimes int) *Client {
	c.retryTimes = retryTimes
	return c
}

func (c *Client) RetryTimesWithInterval(retryTimes int, retryInterval time.Duration) *Client {
	c.retryTimes = retryTimes
	c.retryInterval = retryInterval
	return c
}

func (c *Client) RetryHandler(handle func(r *http.Request)) *Client {
	c.retryHandler = handle
	return c
}

func (c *Client) Proxy(proxyUrl string) *Client {
	if !c.newHttpClient {
		c.httpClient = newHttpClient(c.typ)
		c.newHttpClient = true
	}
	if proxyUrl != "" {
		purl, _ := stdurl.Parse(proxyUrl)
		setProxy(c.httpClient, http.ProxyURL(purl))
	}
	return c
}

func (c *Client) ResetProxy() *Client {
	if !c.newHttpClient {
		return c
	}
	c.httpClient.Transport.(*http.Transport).Proxy = http.ProxyFromEnvironment
	return c
}

func (c *Client) BasicAuth(authUser, authPass string) *Client {
	c.httpRequestOptions = append(c.httpRequestOptions, func(request *http.Request) {
		request.SetBasicAuth(authUser, authPass)
	})
	return c
}

func (c *Client) Clone() *Client {
	return &(*c)
}

func (c *Client) Request(method, url string) *Request {
	r := &Request{
		Method: method, Url: url, client: c,
	}
	return r
}

func (c *Client) Do(r *Request, param, response any) error {
	return r.Client(c).Do(param, response)
}

func (c *Client) Get(url string, param, response any) error {
	return NewRequest(http.MethodGet, url).Client(c).Do(param, response)
}

func (c *Client) GetRequest(url string) *Request {
	return NewRequest(http.MethodGet, url).Client(c)
}

func (c *Client) Post(url string, param, response any) error {
	return NewRequest(http.MethodPost, url).Client(c).Do(param, response)
}

func (c *Client) PostRequest(url string) *Request {
	return NewRequest(http.MethodPost, url).Client(c)
}

func (c *Client) Put(url string, param, response any) error {
	return NewRequest(http.MethodPut, url).Client(c).Do(param, response)
}

func (c *Client) PutRequest(url string) *Request {
	return NewRequest(http.MethodPut, url).Client(c)
}

func (c *Client) Delete(url string, param, response any) error {
	return NewRequest(http.MethodDelete, url).Client(c).Do(param, response)
}

func (c *Client) DeleteRequest(url string) *Request {
	return NewRequest(http.MethodDelete, url).Client(c)
}

func (c *Client) GetX(url string, response any) error {
	return NewRequest(http.MethodGet, url).Client(c).Do(nil, response)
}

func (c *Client) GetRaw(url string, param any) (RawBytes, error) {
	return NewRequest(http.MethodGet, url).Client(c).DoRaw(param)
}

func (c *Client) GetRawX(url string) (RawBytes, error) {
	return NewRequest(http.MethodGet, url).Client(c).DoRaw(nil)
}

func (c *Client) GetStream(url string, param any) (io.ReadCloser, error) {
	return NewRequest(http.MethodGet, url).Client(c).DoStream(param)
}

func (c *Client) GetStreamX(url string) (io.ReadCloser, error) {
	return NewRequest(http.MethodGet, url).Client(c).DoStream(nil)
}

type ResponseBodyCheck interface {
	CheckError() error
}

type RawBytes = []byte
