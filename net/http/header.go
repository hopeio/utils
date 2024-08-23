package http

import (
	"fmt"
	"mime"
	"net/http"
	"strconv"
	"strings"
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

func (h Header) IntoHttpHeader(header http.Header) {
	hlen := len(h)
	for i := 0; i < hlen && i+1 < hlen; i += 2 {
		header.Set(h[i], h[i+1])
	}
}

func (h Header) Clone() Header {
	newh := make(Header, len(h))
	copy(newh, h)
	return newh
}

func CopyHttpHeader(dst, src http.Header) {
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

func copyHeader(src, dst http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func ParseDisposition(disposition string) (mediatype string, params map[string]string, err error) {
	return mime.ParseMediaType(disposition)
}

func ParseRange(rangeHeader string) (start int64, end int64, total int64, err error) {
	// 提取Range值，格式为"bytes unit-unit/*"
	parts := strings.Split(rangeHeader, " ")
	if len(parts) != 2 || parts[0] != "bytes" {
		err = fmt.Errorf("invalid Content-Range format")
		return
	}
	rangeSpec := parts[1]
	info := strings.Split(rangeSpec, "/")
	bounds := strings.Split(info[0], "-")
	start, err = strconv.ParseInt(bounds[0], 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid range start %w", err)
		return
	}

	if len(bounds) > 1 {
		end, err = strconv.ParseInt(bounds[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("invalid range end %w", err)
			return
		}
	} else {
		// 如果只有开始位置，结束位置默认为文件末尾
		end = -1
	}

	if len(info) == 2 && info[1] != "*" {
		total, err = strconv.ParseInt(info[1], 10, 64)
		if err != nil {
			err = fmt.Errorf("invalid range total %w", err)
			return
		}
	} else {
		total = -1
	}
	return
}

func FormatRange(start, end, total int64) string {
	if end <= 0 {
		return fmt.Sprintf("bytes=%d-", start)
	}
	if total <= 0 {
		return fmt.Sprintf("bytes=%d-%d/*", start, end)
	}
	return fmt.Sprintf("bytes=%d-%d/%d", start, end, total)
}
