package fs

import (
	"fmt"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/consts"
	"io"
	"net/http"
)

type ResponseFile struct {
	Name string        `json:"name"`
	Body io.ReadCloser `json:"body,omitempty"`
}

func (res *ResponseFile) Header() httpi.Header {
	return &httpi.SliceHeader{consts.HeaderContentType, consts.ContentTypeOctetStream, consts.HeaderContentDisposition, fmt.Sprintf(consts.AttachmentTmpl, res.Name)}
}

func (res *ResponseFile) WriteTo(writer io.Writer) (int64, error) {
	return io.Copy(writer, res.Body)
}

func (res *ResponseFile) Close() error {
	return res.Body.Close()
}
func (res *ResponseFile) StatusCode() int {
	return http.StatusOK
}

type ResponseFileWriteTo struct {
	Name string               `json:"name"`
	Body httpi.WriterToCloser `json:"body,omitempty"`
}

func (res *ResponseFileWriteTo) Header() httpi.Header {
	return &httpi.SliceHeader{consts.HeaderContentType, consts.ContentTypeOctetStream, consts.HeaderContentDisposition, fmt.Sprintf(consts.AttachmentTmpl, res.Name)}
}

func (res *ResponseFileWriteTo) WriteTo(writer io.Writer) (int64, error) {
	return res.Body.WriteTo(writer)
}

func (res *ResponseFileWriteTo) Close() error {
	return res.Body.Close()
}

func (res *ResponseFileWriteTo) StatusCode() int {
	return http.StatusOK
}
