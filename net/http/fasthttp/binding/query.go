package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/valyala/fasthttp"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	values := req.URI().QueryArgs()
	if err := binding.MapForm(obj, (*ArgsSource)(values)); err != nil {
		return err
	}
	return binding.Validate(obj)
}
