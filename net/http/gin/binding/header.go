package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"net/http"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {
	if err := binding.MapHeader(obj, req.Header); err != nil {
		return err
	}
	return Validate(obj)
}
