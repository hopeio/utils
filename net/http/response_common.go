package http

import (
	"io"
	"net/http"
)

type ICommonResponseWriter interface {
	Status(code int)
	Set(k, v string)
	io.Writer
}

func CommonResponseWrite(w ICommonResponseWriter, httpres IHttpResponse) (int, error) {
	w.Status(httpres.StatusCode())
	for k, v := range httpres.RespHeader() {
		w.Set(k, v)
	}
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
func (w *CommonResponseWriter) Set(k, v string) {
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
