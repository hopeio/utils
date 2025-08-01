package fs

import (
	"fmt"
	httpi "github.com/hopeio/gox/net/http"
	"github.com/hopeio/gox/net/http/consts"
	"io"
	"net/http"
)

type ResponseFile struct {
	Name string        `json:"name"`
	Body io.ReadCloser `json:"body,omitempty"`
}

func (res *ResponseFile) Response(w http.ResponseWriter) (int, error) {
	return res.CommonResponse(httpi.CommonResponseWriter{ResponseWriter: w})
}

func (res *ResponseFile) CommonResponse(w httpi.ICommonResponseWriter) (int, error) {
	header := w.Header()
	header.Set(consts.HeaderContentType, consts.ContentTypeOctetStream)
	header.Set(consts.HeaderContentDisposition, fmt.Sprintf(consts.AttachmentTmpl, res.Name))
	n, err := io.Copy(w, res.Body)
	res.Body.Close()
	return int(n), err
}

type ResponseFileWriteTo struct {
	Name string               `json:"name"`
	Body httpi.WriterToCloser `json:"body,omitempty"`
}

func (res *ResponseFileWriteTo) Response(w http.ResponseWriter) (int, error) {
	return res.CommonResponse(httpi.CommonResponseWriter{ResponseWriter: w})
}

func (res *ResponseFileWriteTo) CommonResponse(w httpi.ICommonResponseWriter) (int, error) {
	header := w.Header()
	header.Set(consts.HeaderContentType, consts.ContentTypeOctetStream)
	header.Set(consts.HeaderContentDisposition, fmt.Sprintf(consts.AttachmentTmpl, res.Name))
	n, err := res.Body.WriteTo(w)
	res.Body.Close()
	return int(n), err
}
