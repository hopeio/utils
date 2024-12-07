/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package middleware

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/apidoc"
	"net/http"
)

type ModName string

func (m ModName) ApiDocMiddle(ctx *gin.Context) {
	currentRouteName := ctx.Request.RequestURI[len(ctx.Request.Method):]

	var pathItem *openapi3.PathItem

	doc := apidoc.GetDoc(apidoc.Dir, string(m))

	if doc.Paths != nil {
		if path := doc.Paths.Value(currentRouteName); path != nil {
			pathItem = path
		} else {
			pathItem = new(openapi3.PathItem)
		}
	} else {
		doc.Paths = openapi3.NewPaths()
		pathItem = new(openapi3.PathItem)
	}

	parameters := make([]*openapi3.ParameterRef, len(ctx.Params), len(ctx.Params))

	params := ctx.Params

	for i := range params {
		key := params[i].Key

		//val := params[i].ValueRaw
		parameters[i] = &openapi3.ParameterRef{
			Value: &openapi3.Parameter{
				Name:        key,
				In:          "path",
				Description: "Description",
			},
		}
	}

	if stop, _ := ctx.GetQuery("apidoc"); stop == "stop" {
		defer apidoc.WriteToFile(apidoc.Dir, string(m))
	}

	ress := openapi3.NewResponses()
	op := openapi3.Operation{
		Description: "Description",
		Tags:        []string{"Tags"},
		Summary:     "Summary",
		OperationID: "currentRouteName" + ctx.Request.Method,
		Parameters:  parameters,
		Responses:   ress,
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
	doc.Paths.Set(currentRouteName, pathItem)
	ctx.Next()
}
