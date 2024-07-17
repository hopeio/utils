package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/valyala/fasthttp"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (x xmlBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	return binding.DecodeXmlData(req.Request.Body(), obj)
}
