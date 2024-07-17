package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/debug"
)

func Debug(r *gin.Engine) {
	r.Any("/debug/*path", Wrap(debug.Handler()))
}
