package client

import (
	"github.com/hopeio/utils/fs"
	"net/http"
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
type Downloader = Client

func NewDownloader() *Downloader {
	return &Downloader{
		typ:           ClientTypeDownload,
		httpClient:    DefaultDownloadHttpClient,
		retryTimes:    3,
		retryInterval: time.Second,
		logger:        nil,
		logLevel:      LogLevelSilent,
	}
}

func (c *Downloader) Download(filepath string, r *DownloadReq) error {
	return r.WithDownloader(c).Download(filepath)
}

func (c *Downloader) DownloadReq(url string) *DownloadReq {
	return NewDownloadReq(url).WithDownloader(c)
}

const DownloadKey = fs.DownloadKey
