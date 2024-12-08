/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/apidoc"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
	"net/http"
)

type ModName string

func (m ModName) ApiDocMiddle(ctx *gin.Context) {
	currentRouteName := ctx.Request.RequestURI[len(ctx.Request.Method):]

	var pathItem *v3high.PathItem

	doc := apidoc.GetDoc(apidoc.Dir, string(m))

	if path, ok := doc.Paths.PathItems.Load(currentRouteName); ok {
		pathItem = path
	} else {
		pathItem = new(v3high.PathItem)
	}

	parameters := make([]*v3high.Parameter, len(ctx.Params), len(ctx.Params))

	params := ctx.Params

	for i := range params {
		key := params[i].Key

		//val := params[i].ValueRaw
		parameters[i] = &v3high.Parameter{
			Name:        key,
			In:          "path",
			Description: "Description",
		}
	}

	if stop, _ := ctx.GetQuery("apidoc"); stop == "stop" {
		defer apidoc.WriteToFile(apidoc.Dir, string(m))
	}

	op := v3high.Operation{
		Description: "Description",
		Tags:        []string{"Tags"},
		Summary:     "Summary",
		OperationId: "currentRouteName" + ctx.Request.Method,
		Parameters:  parameters,
	}

	switch ctx.Request.Method {
	case http.MethodGet:
		pathItem.Get = &op
	case http.MethodPost:
		pathItem.Post = &op
	case http.MethodPut:
		pathItem.Put = &op
	case http.MethodDelete:
		pathItem.Delete = &op
	case http.MethodOptions:
		pathItem.Options = &op
	case http.MethodPatch:
		pathItem.Patch = &op
	case http.MethodHead:
		pathItem.Head = &op
	}
	doc.Paths.PathItems.Set(currentRouteName, pathItem)
	ctx.Next()
}
