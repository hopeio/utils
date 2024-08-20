package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/errors/errcode"
	"github.com/hopeio/utils/types/funcs"
	"net/http"
)

// only example
func Handler[REQ, RES any](service funcs.GrpcServiceMethod[*REQ, *RES]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(REQ)
		if len(ctx.Params) > 0 {
			err := ctx.ShouldBindUri(req)
			if err != nil {
				ctx.JSON(http.StatusOK, errcode.InvalidArgument.Wrap(err))
				return
			}
		}
		if len(ctx.Request.URL.RawQuery) > 0 {
			err := ctx.ShouldBindQuery(req)
			if err != nil {
				ctx.JSON(http.StatusOK, errcode.InvalidArgument.Wrap(err))
				return
			}
		}
		if ctx.Request.Body != nil && ctx.Request.ContentLength != 0 {
			err := ctx.ShouldBindJSON(req)
			if err != nil {
				ctx.JSON(http.StatusOK, errcode.InvalidArgument.Wrap(err))
				return
			}
		}
		res, err := service(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusOK, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
