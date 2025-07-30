package fiber

import (
	"github.com/hopeio/gox/strings"
	"github.com/valyala/fasthttp"
)

type ResponseHeader struct {
	*fasthttp.ResponseHeader
}

func (h ResponseHeader) Add(key, value string) {
	h.ResponseHeader.Add(key, value)
}

func (h ResponseHeader) Set(key, value string) {
	h.ResponseHeader.Set(key, value)
}

func (h ResponseHeader) Get(key string) string {
	return strings.BytesToString(h.ResponseHeader.Peek(key))
}

func (h ResponseHeader) Values(key string) []string {
	byteValues := h.ResponseHeader.PeekAll(key)
	values := make([]string, len(byteValues))
	for i := range byteValues {
		values[i] = strings.BytesToString(byteValues[i])
	}
	return values
}

func (h ResponseHeader) Range(f func(key, value string)) {
	h.ResponseHeader.VisitAll(func(key, value []byte) {
		f(strings.BytesToString(key), strings.BytesToString(value))
	})
}

type RequestHeader struct {
	*fasthttp.RequestHeader
}

func (h RequestHeader) Add(key, value string) {
	h.RequestHeader.Add(key, value)
}

func (h RequestHeader) Set(key, value string) {
	h.RequestHeader.Set(key, value)
}

func (h RequestHeader) Get(key string) string {
	return strings.BytesToString(h.RequestHeader.Peek(key))
}

func (h RequestHeader) Values(key string) []string {
	byteValues := h.RequestHeader.PeekAll(key)
	values := make([]string, len(byteValues))
	for i := range byteValues {
		values[i] = strings.BytesToString(byteValues[i])
	}
	return values
}

func (h RequestHeader) Range(f func(key, value string)) {
	h.RequestHeader.VisitAll(func(key, value []byte) {
		f(strings.BytesToString(key), strings.BytesToString(value))
	})
}
