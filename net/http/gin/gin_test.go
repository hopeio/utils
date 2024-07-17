package gin

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRoute(t *testing.T) {
	i := gin.New()
	i.GET("/:id/:name/:path", func(context *gin.Context) { context.Writer.WriteString("/:id/:name/:path") })
	i.GET("/id/name/path", func(context *gin.Context) { context.Writer.WriteString("/id/name/path") })
	i.Run(":8080")
}
