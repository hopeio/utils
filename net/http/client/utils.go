package client

import (
	"fmt"
	urli "github.com/hopeio/utils/net/url"
	"net/http"
	"net/url"
	"os"
	"time"
)

func SetTag(t string) {
	urli.SetTag(t)
}

func SetProxyEnv(url string) {
	os.Setenv("HTTP_PROXY", url)
	os.Setenv("HTTPS_PROXY", url)
}

func setTimeout(client *http.Client, timeout time.Duration) {
	if client == nil {
		client = DefaultHttpClient
	}
	if timeout < time.Second {
		timeout = timeout * time.Second
	}
	client.Timeout = timeout
}

func setProxy(client *http.Client, proxy func(*http.Request) (*url.URL, error)) {
	client.Transport.(*http.Transport).Proxy = proxy
}

func CloseReaderWrap(err error) error {
	return fmt.Errorf("close reader error: %w", err)
}
