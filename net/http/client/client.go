package client

import (
	"errors"
	httpi "github.com/hopeio/utils/net/http"
	"io"
	"net"
	"net/http"
	stdurl "net/url"
	"time"
)

// 不是并发安全的

var (
	DefaultHttpClient = newHttpClient()
	DefaultLogLevel   = LogLevelError
)

const timeout = time.Minute

func newHttpClient() *http.Client {
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

func SetTimeout(timeout time.Duration) {
	DefaultHttpClient.Timeout = timeout
}

func DisableLog() {
	DefaultLogLevel = LogLevelSilent
}

func SetAccessLog(log AccessLog) {
	defaultLog = log
}

func SetProxy(url string) {
	purl, _ := stdurl.Parse(url)
	setProxy(DefaultHttpClient, http.ProxyURL(purl))
}

func ResetProxy() {
	DefaultHttpClient.Transport.(*http.Transport).Proxy = http.ProxyFromEnvironment
}

func SetHttpClient(client *http.Client) {
	DefaultHttpClient = client
}

// Client ...
type Client struct {

	// httpClient settings
	httpClient    *http.Client
	newHttpClient bool

	parseTag string // 默认json

	// request
	httpRequestOptions []HttpRequestOption
	header             http.Header //公共请求头

	// response
	responseHandler func(response *http.Response) (retry bool, data []byte, err error)
	resDataHandler  func(data []byte) ([]byte, error)

	// logger
	logger   AccessLog
	logLevel LogLevel

	// retry
	retryTimes    int
	retryInterval time.Duration
	retryHandler  func(*Client)

	req *Request
}

func New() *Client {
	return &Client{httpClient: DefaultHttpClient, logger: defaultLog, logLevel: DefaultLogLevel, retryInterval: 200 * time.Millisecond}
}

func (c *Client) Header(header http.Header) *Client {
	if c.header == nil {
		c.header = make(http.Header)
	}
	httpi.CopyHttpHeader(header, c.header)
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
func (c *Client) ResponseHandler(handler func(response *http.Response) (retry bool, data []byte, err error)) *Client {
	c.responseHandler = handler
	return c
}

func (c *Client) ResDataHandler(handler func(response []byte) ([]byte, error)) *Client {
	c.resDataHandler = handler
	return c
}

// 设置过期时间,仅对单次请求有效
func (c *Client) Timeout(timeout time.Duration) *Client {
	if !c.newHttpClient {
		c.httpClient = newHttpClient()
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
		c.httpClient = newHttpClient()
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

func (c *Client) RetryHandler(handle func(*Client)) *Client {
	c.retryHandler = handle
	return c
}

func (c *Client) Proxy(proxyUrl string) *Client {
	if !c.newHttpClient {
		c.httpClient = newHttpClient()
		c.newHttpClient = true
	}
	if proxyUrl != "" {
		purl, _ := stdurl.Parse(proxyUrl)
		setProxy(c.httpClient, http.ProxyURL(purl))
	}
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
	c.req = r
	return r
}

func (c *Client) RequestDo(param, response any) error {
	if c.req == nil {
		return errors.New("request is nil")
	}
	req := c.req
	c.req = nil
	return req.Do(param, response)
}

func (c *Client) Do(r *Request, param, response any) error {
	return r.WithClient(c).Do(param, response)
}

func (c *Client) Get(url string, param, response any) error {
	return NewRequest(http.MethodGet, url).WithClient(c).Do(param, response)
}

func (c *Client) GetRequest(url string) *Request {
	return NewRequest(http.MethodGet, url).WithClient(c)
}

func (c *Client) Post(url string, param, response any) error {
	return NewRequest(http.MethodPost, url).WithClient(c).Do(param, response)
}

func (c *Client) PostRequest(url string) *Request {
	return NewRequest(http.MethodPost, url).WithClient(c)
}

func (c *Client) Put(url string, param, response any) error {
	return NewRequest(http.MethodPut, url).WithClient(c).Do(param, response)
}

func (c *Client) PutRequest(url string) *Request {
	return NewRequest(http.MethodPut, url).WithClient(c)
}

func (c *Client) Delete(url string, param, response any) error {
	return NewRequest(http.MethodDelete, url).WithClient(c).Do(param, response)
}

func (c *Client) DeleteRequest(url string) *Request {
	return NewRequest(http.MethodDelete, url).WithClient(c)
}

func (c *Client) GetX(url string, response any) error {
	return NewRequest(http.MethodGet, url).WithClient(c).Do(nil, response)
}

func (c *Client) GetStream(url string, param any) (io.ReadCloser, error) {
	return NewRequest(http.MethodGet, url).DoStream(param)
}

func (c *Client) GetStreamX(url string) (io.ReadCloser, error) {
	return NewRequest(http.MethodGet, url).DoStream(nil)
}

type ResponseBodyCheck interface {
	CheckError() error
}

type RawBytes = []byte
