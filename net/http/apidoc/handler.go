package apidoc

import (
	"bytes"
	"github.com/go-openapi/runtime/middleware"
	httpi "github.com/hopeio/utils/net/http"
	"net/http"
	"os"
	"path"
)

// 目录结构 ./api/mod/mod.swagger.json ./api/mod/mod.apidoc.md
// 请求路由 /apidoc /apidoc/swagger/mod/mod.swagger.json /apidoc/markdown/mod/mod.apidoc.md
var UriPrefix = "/apidoc"
var ApiDocDir = "./apidoc/"

const TypeSwagger = "swagger"
const TypeMarkdown = "markdown"
const SwaggerEXT = ".swagger.json"
const MarkDownEXT = ".apidoc.md"

func Swagger(w http.ResponseWriter, r *http.Request) {
	prefixUri := UriPrefix + "/" + TypeSwagger + "/"
	if r.RequestURI[len(r.RequestURI)-5:] == ".json" {
		b, err := os.ReadFile(ApiDocDir + r.RequestURI[len(prefixUri):])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set(httpi.HeaderContentType, "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	}
	mod := r.RequestURI[len(prefixUri):]
	middleware.Redoc(middleware.RedocOpts{
		BasePath: prefixUri,
		SpecURL:  path.Join(prefixUri+mod, mod+SwaggerEXT),
		Path:     mod,
	}, http.NotFoundHandler()).ServeHTTP(w, r)
}

func Markdown(w http.ResponseWriter, r *http.Request) {
	prefixUri := UriPrefix + "/" + TypeMarkdown + "/"
	mod := r.RequestURI[len(prefixUri):]
	b, err := os.ReadFile(ApiDocDir + mod + "/" + mod + MarkDownEXT)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func DocList(w http.ResponseWriter, r *http.Request) {
	fileInfos, err := os.ReadDir(ApiDocDir)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var buff bytes.Buffer
	for i := range fileInfos {
		if fileInfos[i].IsDir() {
			buff.Write([]byte(`<a href="` + r.RequestURI + "/swagger/" + fileInfos[i].Name() + `"> swagger: ` + fileInfos[i].Name() + `</a><br>`))
			buff.Write([]byte(`<a href="` + r.RequestURI + "/markdown/" + fileInfos[i].Name() + `"> markdown: ` + fileInfos[i].Name() + `</a><br>`))
		}
	}
	w.Write(buff.Bytes())
}

func OpenApi(mux *http.ServeMux, filePath string) {
	ApiDocDir = filePath
	mux.Handle(UriPrefix, http.HandlerFunc(DocList))
	mux.Handle(UriPrefix+"/markdown/", http.HandlerFunc(Markdown))
	mux.Handle(UriPrefix+"/swagger/", http.HandlerFunc(Swagger))
}
