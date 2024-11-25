/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/gofiber/fiber/v3"
)

func NewReq[REQ any](c fiber.Ctx) (*REQ, error) {
	req := new(REQ)
	err := Bind(c, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func BindBody(r fiber.Ctx, obj interface{}) error {
	return MustBindWith(r, obj, CustomBody)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func BindQuery(c fiber.Ctx, obj interface{}) error {
	return MustBindWith(c, obj, Query)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// BindUri binds the passed struct pointer using binding.Uri.
// It will abort the request with HTTP 400 if any error occurs.
func BindUri(c fiber.Ctx, obj interface{}) error {
	return ShouldBindUri(c, obj)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func MustBindWith(c fiber.Ctx, obj interface{}, b Binding) error {
	return ShouldBindWith(c, obj, b)
}

// ShouldBind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//
//	"application/json" --> JSON binding
//	"application/xml"  --> XML binding
//
// otherwise --> returns an error
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.GinBind() but this method does not set the response status code to 400 and abort if the json is not valid.
func ShouldBind(c fiber.Ctx, obj interface{}) error {
	b := Default(c.Method(), c.Request().Header.ContentType())
	return ShouldBindWith(c, obj, b)
}

func ShouldBindBody(c fiber.Ctx, obj interface{}) error {
	return ShouldBindWith(c, obj, CustomBody)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func ShouldBindQuery(c fiber.Ctx, obj interface{}) error {
	return ShouldBindWith(c, obj, Query)
}

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func ShouldBindUri(c fiber.Ctx, obj interface{}) error {
	return Uri.Bind(c, obj)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(c fiber.Ctx, obj interface{}, b Binding) error {
	return b.Bind(c, obj)
}
