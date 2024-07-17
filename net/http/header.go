package http

import (
	"net/http"
)

type Header []string

func NewHeader() *Header {
	h := make(Header, 0, 6)
	return &h
}

func (h *Header) Add(k, v string) *Header {
	*h = append(*h, k, v)
	return h
}

func (h *Header) Set(k, v string) *Header {
	header := *h
	for i, s := range header {
		if s == k {
			header[i+1] = v
			return h
		}
	}
	return h.Add(k, v)
}

func (h *Header) IntoHttpHeader(header http.Header) {
	res := *h
	hlen := len(res)
	for i := 0; i < hlen && i+1 < hlen; i += 2 {
		header.Set(res[i], res[i+1])
	}
}

func (h Header) Clone() Header {
	newh := make(Header, len(h))
	copy(newh, h)
	return newh
}

func CopyHttpHeader(src, dst http.Header) {
	if src == nil {
		return
	}

	// Find total number of values.
	nv := 0
	for _, vv := range src {
		nv += len(vv)
	}
	sv := make([]string, nv) // shared backing array for headers' values

	for k, vv := range src {
		if vv == nil {
			continue
		}
		n := copy(sv, vv)
		dst[k] = sv[:n:n]
		sv = sv[n:]
	}
}
