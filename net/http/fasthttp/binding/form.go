package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	"github.com/valyala/fasthttp"
)

type formPostBinding struct{}
type formMultipartBinding struct{}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*MultipartRequest)(&req.Request), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}
func (formPostBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*ArgsSource)(req.PostArgs()), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
