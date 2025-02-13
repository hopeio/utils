/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	"github.com/hopeio/utils/os/fs"
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

func NewDownloader(options ...Option) *Downloader {
	downloader := &Downloader{
		typ:           ClientTypeDownload,
		httpClient:    DefaultDownloadHttpClient,
		retryTimes:    3,
		retryInterval: time.Second,
		logger:        nil,
		logLevel:      LogLevelSilent,
	}
	for _, opt := range options {
		opt(downloader)
	}
	return downloader
}

func (c *Downloader) Download(filepath string, r *DownloadReq) error {
	return r.Downloader(c).Download(filepath)
}

func (c *Downloader) DownloadReq(url string) *DownloadReq {
	return NewDownloadReq(url).Downloader(c)
}

const DownloadKey = fs.DownloadKey
