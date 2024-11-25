/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/spec"
	"github.com/hopeio/utils/net/http/apidoc"
	"net/http"
)

type ModName string

func (m ModName) ApiDocMiddle(ctx *gin.Context) {
	currentRouteName := ctx.Request.RequestURI[len(ctx.Request.Method):]

	var pathItem *spec.PathItem

	doc := apidoc.GetDoc(apidoc.Dir, string(m))

	if doc.Paths != nil && doc.Paths.Paths != nil {
		if path, ok := doc.Paths.Paths[currentRouteName]; ok {
			pathItem = &path
		} else {
			pathItem = new(spec.PathItem)
		}
	} else {
		doc.Paths = &spec.Paths{Paths: map[string]spec.PathItem{}}
		pathItem = new(spec.PathItem)
	}

	parameters := make([]spec.Parameter, len(ctx.Params), len(ctx.Params))

	params := ctx.Params

	for i := range params {
		key := params[i].Key

		//val := params[i].ValueRaw
		parameters[i] = spec.Parameter{
			ParamProps: spec.ParamProps{
				Name:        key,
				In:          "path",
				Description: "Description",
			},
		}
	}

	if stop, _ := ctx.GetQuery("apidoc"); stop == "stop" {
		defer apidoc.WriteToFile(apidoc.Dir, string(m))
	}

	var res spec.Responses
	op := spec.Operation{
		OperationProps: spec.OperationProps{
			Description: "Description",
			Consumes:    []string{"application/x-www-form-urlencoded"},
			Tags:        []string{"Tags"},
			Summary:     "Summary",
			ID:          "currentRouteName" + ctx.Request.Method,
			Parameters:  parameters,
			Responses:   &res,
		},
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
	doc.Paths.Paths[currentRouteName] = *pathItem
	ctx.Next()
}
