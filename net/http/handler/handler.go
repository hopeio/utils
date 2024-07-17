package handler

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/types/funcs"
	"net/http"
)

// TODO
func commonHandler[REQ, RES any](method funcs.GrpcServiceMethod[*REQ, *RES]) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(REQ)
		err := binding.Bind(r, req)
		if err != nil {
			return
		}
	})
}
