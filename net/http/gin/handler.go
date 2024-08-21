package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/errors/errcode"
	httpi "github.com/hopeio/utils/net/http"
	"github.com/hopeio/utils/net/http/gin/binding"
	"github.com/hopeio/utils/types/funcs"
	"net/http"
)

// only example
func Handler[REQ, RES any](service funcs.GrpcServiceMethod[*REQ, *RES]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(REQ)
		err := binding.Bind(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, errcode.InvalidArgument.Wrap(err))
			return
		}
		res, err := service(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, err)
			return
		}
		ctx.JSON(http.StatusOK, httpi.NewSuccessResData(res))
	}
}
