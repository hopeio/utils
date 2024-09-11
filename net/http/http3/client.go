package http3

import (
	"github.com/quic-go/quic-go/http3"
	"net/http"
)

func NewClient() *http.Client {
	return &http.Client{
		Transport: &http3.RoundTripper{},
	}
}
