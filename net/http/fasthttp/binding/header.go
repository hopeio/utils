package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	"github.com/valyala/fasthttp"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {

	if err := mtos.MapFormByTag(obj, (*HeaderSource)(&req.Request.Header), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}
