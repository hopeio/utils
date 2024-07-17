package binding

import (
	"fmt"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/valyala/fasthttp"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (j jsonBinding) Bind(req *fasthttp.RequestCtx, obj interface{}) error {
	body := req.Request.Body()
	if req == nil || body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.DecodeJsonData(body, obj)
}
