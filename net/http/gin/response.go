package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/errors/errcode"
	httpi "github.com/hopeio/utils/net/http"
)

func RespErrcode(ctx *gin.Context, code errcode.ErrCode) {
	httpi.RespErrcode(ctx.Writer, code)
}

func SuccessRespMsg(ctx *gin.Context, msg string) {
	httpi.SuccessRespMsg(ctx.Writer, msg)
}

func RespErrRep(ctx *gin.Context, rep *errcode.ErrRep) {
	httpi.RespErrRep(ctx.Writer, rep)
}

func Response(ctx *gin.Context, code errcode.ErrCode, msg string, data interface{}) {
	httpi.Response(ctx.Writer, code, msg, data)
}

func SuccessResponse[T any](ctx *gin.Context, msg string, data T) {
	httpi.SuccessResponse(ctx.Writer, msg, data)
}
