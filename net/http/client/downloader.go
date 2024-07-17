package client

import (
	"errors"
	"github.com/hopeio/utils/io/fs"
	httpi "github.com/hopeio/utils/net/http"
	"io"
	"net/http"
	stdurl "net/url"
	"time"
)

var DefaultDownloadHttpClient = newDownloadHttpClient()

func newDownloadHttpClient() *http.Client {
	return &http.Client{
		//Timeout: timeout * 2,
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment, // 代理使用
			ForceAttemptHTTP2: true,
		},
	}
}

// TODO: Range Status(206) PartialContent 下载
type Downloader struct {
	httpClient         *http.Client
	newHttpClient      bool
	header             http.Header //公共请求头
	responseHandler    func(response *http.Response) (retry bool, data []byte, err error)
	resDataHandler     func(data []byte) ([]byte, error)
	retryTimes         int
	retryInterval      time.Duration
	httpRequestOptions []HttpRequestOption

	req *DownloadReq
}

func NewDownloader() *Downloader {
	return &Downloader{
		httpClient:    DefaultDownloadHttpClient,
		retryTimes:    3,
		retryInterval: time.Second,
	}
}

func (c *Downloader) HttpClient(client *http.Client) *Downloader {
	c.httpClient = client
	c.newHttpClient = true
	return c
}

func (c *Downloader) SetHttpClient(opt HttpClientOption) *Downloader {
	if !c.newHttpClient {
		c.httpClient = newDownloadHttpClient()
		c.newHttpClient = true
	}
	opt(c.httpClient)
	return c
}

func (c *Downloader) Options(opts ...HttpRequestOption) *Downloader {
	c.httpRequestOptions = append(c.httpRequestOptions, opts...)
	return c
}

func (c *Downloader) ResponseHandler(handler func(response *http.Response) (retry bool, data []byte, err error)) *Downloader {
	c.responseHandler = handler
	return c
}

func (c *Downloader) ResDataHandler(handler func(response []byte) ([]byte, error)) *Downloader {
	c.resDataHandler = handler
	return c
}

func (c *Downloader) Header(header http.Header) *Downloader {
	if c.header == nil {
		c.header = make(http.Header)
	}
	httpi.CopyHttpHeader(header, c.header)
	return c
}

func (c *Downloader) AddHeader(header, value string) *Downloader {
	if c.header == nil {
		c.header = make(http.Header)
	}
	c.header.Add(header, value)
	return c
}

func (c *Downloader) BasicAuth(authUser, authPass string) *Downloader {
	c.httpRequestOptions = append(c.httpRequestOptions, func(request *http.Request) {
		request.SetBasicAuth(authUser, authPass)
	})
	return c
}

func (c *Downloader) Clone() *Downloader {
	return &(*c)
}

func (c *Downloader) Timeout(timeout time.Duration) *Downloader {
	if !c.newHttpClient {
		c.httpClient = newHttpClient()
		c.newHttpClient = true
	}
	setTimeout(c.httpClient, timeout)
	return c
}

func (c *Downloader) Proxy(proxyUrl string) *Downloader {
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

func (c *Downloader) GetResponse(r *DownloadReq) (*http.Response, error) {
	return r.WithDownloader(c).GetResponse()
}

func (c *Downloader) GetReader(r *DownloadReq) (io.ReadCloser, error) {
	return r.WithDownloader(c).GetReader()
}

func (c *Downloader) Download(filepath string, r *DownloadReq) error {
	return r.WithDownloader(c).Download(filepath)
}

func (c *Downloader) DownloadReq(url string) *DownloadReq {
	return NewDownloadReq(url).WithDownloader(c)
}

func (c *Downloader) ReqDownload(filepath string) error {
	if c.req == nil {
		return errors.New("request is nil")
	}
	req := c.req
	c.req = nil
	return req.Download(filepath)
}

const DownloadKey = fs.DownloadKey
