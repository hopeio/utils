package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/valyala/fasthttp"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	return binding.DecodeProtobuf(req.Request.Body(), obj)
}
