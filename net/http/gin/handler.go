package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/types/funcs"
)

// TODO:
func commonHandler[REQ, RES any](method funcs.GrpcServiceMethod[*REQ, *RES]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(REQ)
		err := ctx.Bind(req)
		if err != nil {
			return
		}
	}
}
