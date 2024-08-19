package client

import (
	"bytes"
	"context"
	"fmt"
	httpi "github.com/hopeio/utils/net/http"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
)

var (
	ContentTypeKey = http.CanonicalHeaderKey("Content-Type")
)

type UploadReq struct {
	Url      string
	uploader *Uploader
	ctx      context.Context
	header   httpi.Header //请求级请求头
	Boundary string
	FormData map[string]string
	Files    []*File
}

type File struct {
	Name        string
	Param       string
	ContentType string
	io.Reader
}

func NewUploadReq(url string) *UploadReq {
	return &UploadReq{
		ctx:      context.Background(),
		Url:      url,
		uploader: DefaultUploader,
	}
}

func (r *UploadReq) WithDownloader(u *Uploader) *UploadReq {
	r.uploader = u
	return r
}

func (r *UploadReq) UploadMultipart() error {
	body := bufPool.Get().(*bytes.Buffer)
	w := multipart.NewWriter(body)

	for k, v := range r.FormData {
		if err := w.WriteField(k, v); err != nil {
			return err
		}
	}
	cbuf := make([]byte, 512)
	for _, file := range r.Files {
		if file.ContentType == "" {
			size, err := file.Reader.Read(cbuf)
			if err != nil && err != io.EOF {
				return err
			}
			partWriter, err := w.CreatePart(createMultipartHeader(file.Param, file.Name,
				http.DetectContentType(cbuf[:size])))
			if err != nil {
				return err
			}
			if _, err = partWriter.Write(cbuf[:size]); err != nil {
				return err
			}

			_, err = io.Copy(partWriter, file.Reader)
			if err != nil {
				return err
			}
		} else {
			partWriter, err := w.CreatePart(createMultipartHeader(file.Param, file.Name, file.ContentType))
			if err != nil {
				return err
			}

			_, err = io.Copy(partWriter, file.Reader)
			if err != nil {
				return err
			}
		}
	}
	r.header.Set(httpi.HeaderContentType, w.FormDataContentType())
	err := w.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, r.Url, body)
	if err != nil {
		return err
	}
	d := r.uploader
	httpi.CopyHttpHeader(d.header, req.Header)
	for i := 0; i+1 < len(r.header); i += 2 {
		req.Header.Set(r.header[i], r.header[i+1])
	}
	for _, opt := range d.httpRequestOptions {
		opt(req)
	}
	_, err = r.uploader.httpClient.Do(req)
	if err != nil {
		return err
	}
	// TODO: error handler, retry
	return nil
}

func createMultipartHeader(param, fileName, contentType string) textproto.MIMEHeader {
	header := make(textproto.MIMEHeader)

	var contentDispositionValue string
	if IsStringEmpty(fileName) {
		contentDispositionValue = fmt.Sprintf(`form-data; name="%s"`, param)
	} else {
		contentDispositionValue = fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			param, escapeQuotes(fileName))
	}
	header.Set("Content-Disposition", contentDispositionValue)

	if !IsStringEmpty(contentType) {
		header.Set(ContentTypeKey, contentType)
	}
	return header
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func IsStringEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
