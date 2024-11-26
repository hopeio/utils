/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package fiber

import (
	"bytes"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/apidoc"
	"github.com/hopeio/utils/os/fs"
	"html/template"
	"net/http"
	"os"
	"path"
)

func Swagger(ctx fiber.Ctx) error {
	prefixUri := apidoc.UriPrefix + "/" + apidoc.TypeSwagger + "/"
	requestURI := string(ctx.Request().URI().RequestURI())
	if requestURI[len(requestURI)-5:] == ".json" {
		b, err := os.ReadFile(apidoc.Dir + requestURI[len(prefixUri):])
		if err != nil {
			return err
		}
		ctx.Response().Header.Set("Content-Type", "application/json; charset=utf-8")
		ctx.Response().SetStatusCode(http.StatusOK)
		ctx.Write(b)
		return nil
	}
	mod := requestURI[len(prefixUri):]

	opts := middleware.RedocOpts{
		BasePath: prefixUri,
		SpecURL:  path.Join(prefixUri+mod, mod+apidoc.SwaggerEXT),
		Path:     mod,
	}
	opts.EnsureDefaults()
	pth := path.Join(opts.BasePath, opts.Path)
	tmpl := template.Must(template.New("redoc").Parse(opts.Template))
	assets := bytes.NewBuffer(nil)
	if err := tmpl.Execute(assets, opts); err != nil {
		panic(fmt.Errorf("cannot execute template: %w", err))
	}
	if path.Clean(ctx.Path()) == pth {
		ctx.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
		ctx.Response().SetStatusCode(http.StatusOK)
		_, _ = ctx.Write(assets.Bytes())
	}
	return nil
}

func Markdown(ctx fiber.Ctx) error {
	prefixUri := apidoc.UriPrefix + "/" + apidoc.TypeMarkdown + "/"
	mod := string(ctx.Request().URI().RequestURI()[len(prefixUri):])
	b, err := os.ReadFile(apidoc.Dir + mod + "/" + mod + apidoc.MarkDownEXT)
	if err != nil {
		return err
	}
	ctx.Response().Header.Set("Content-Type", "text/markdown; charset=utf-8")
	ctx.Response().SetStatusCode(http.StatusOK)
	ctx.Write(b)
	return nil
}

func DocList(ctx fiber.Ctx) error {
	fileInfos, err := os.ReadDir(apidoc.Dir)
	if err != nil {
		return err
	}
	requestURI := string(ctx.Request().URI().RequestURI())
	var buff bytes.Buffer
	for i := range fileInfos {
		if fileInfos[i].IsDir() {
			buff.Write([]byte(`<a href="` + requestURI + "/swagger/" + fileInfos[i].Name() + `"> swagger: ` + fileInfos[i].Name() + `</a><br>`))
			buff.Write([]byte(`<a href="` + requestURI + "/markdown/" + fileInfos[i].Name() + `"> markdown: ` + fileInfos[i].Name() + `</a><br>`))
		}
	}
	ctx.Write(buff.Bytes())
	return nil
}

func OpenApi(mux *fiber.App, uriPrefix, dir string) {
	if uriPrefix != "" {
		apidoc.UriPrefix = uriPrefix
	}
	if dir != "" {
		if b := dir[len(dir)-1:]; b == "/" || b == "\\" {
			apidoc.Dir = dir
		} else {
			apidoc.Dir = dir + fs.PathSeparator
		}
	}
	mux.Get(apidoc.UriPrefix+"/markdown/", Markdown)
	mux.Get(apidoc.UriPrefix, DocList)
	mux.Get(apidoc.UriPrefix+"/swagger/", Swagger)
}
