package client

import (
	"context"
	httpi "github.com/hopeio/utils/net/http"
)

type UploadReq struct {
	Url      string
	uploader *Uploader
	ctx      context.Context
	header   httpi.Header //请求级请求头
	Boundary string
	Files    []string
	Fields   map[string]string
}

func (r *UploadReq) Upload() {

}
