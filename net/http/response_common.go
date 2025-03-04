package http

import (
	"io"
	"net/http"
)

type ICommonResponseWriter interface {
	Status(code int)
	Header() Header
	io.Writer
}

func CommonResponseWrite(w ICommonResponseWriter, httpres IHttpResponse) (int, error) {
	w.Status(httpres.StatusCode())
	header := w.Header()
	httpres.Header().Range(header.Set)
	i, err := httpres.WriteTo(w)
	if err != nil {
		return int(i), err
	}
	err = httpres.Close()
	if err != nil {
		return int(i), err
	}
	return int(i), err
}

type CommonResponseWriter struct {
	http.ResponseWriter
}

func (w *CommonResponseWriter) Status(code int) {
	w.WriteHeader(code)
}
func (w *CommonResponseWriter) SetHeader(k, v string) {
	w.Header().Set(k, v)
}
func (w *CommonResponseWriter) Write(p []byte) (int, error) {
	return w.ResponseWriter.Write(p)
}

type ICommonHttpResponseTo interface {
	Response(w CommonResponseWriter) (int, error)
}

type CommonHttpResponseTo struct {
	IHttpResponseTo
}

type ICommonResponseTo interface {
	Response(w ICommonResponseWriter) (int, error)
}
