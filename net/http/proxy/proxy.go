package proxy

import (
	"io"
	"net/http"
	"time"
)

func Proxy(host string) http.Handler {
	var proxyClient = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100, // 提升目标服务器连接复用率
			IdleConnTimeout:     90 * time.Second,
		},
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Host = host
		req, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
		if err != nil {
			return
		}
		req.Header = r.Header
		resp, err := proxyClient.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		resHeader := w.Header()
		for key, values := range resp.Header {
			for _, value := range values {
				resHeader.Add(key, value)
			}
		}
		w.WriteHeader(resp.StatusCode)

		io.Copy(w, resp.Body)
	})
}
