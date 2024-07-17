//go:build !nomsgpack
// +build !nomsgpack

package binding

import (
	"bytes"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/valyala/fasthttp"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (m msgpackBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	return m.BindBody(req.Request.Body(), obj)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}) error {
	return binding.DecodeMsgPack(bytes.NewReader(body), obj)
}
