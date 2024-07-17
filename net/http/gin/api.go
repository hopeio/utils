package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/apidoc"
	_ "github.com/ugorji/go"
)

func OpenApi(mux *gin.Engine, filePath string) {
	apidoc.ApiDocDir = filePath
	mux.GET(apidoc.UriPrefix+"/markdown/*file", Wrap(apidoc.Markdown))
	mux.GET(apidoc.UriPrefix, Wrap(apidoc.DocList))
	mux.GET(apidoc.UriPrefix+"/swagger/*file", Wrap(apidoc.Swagger))
}
