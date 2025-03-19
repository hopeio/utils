/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/reflect/mtos"
	stringsi "github.com/hopeio/utils/strings"
	"net/http"
)

type Binding interface {
	Name() string

	Bind(fiber.Ctx, interface{}) error
}

type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

var (
	CustomBody    = bodyBinding{name: "json", unmarshaller: json.Unmarshal}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
)

func Default(method string, contentType []byte) Binding {
	if method == http.MethodGet {
		return Query
	}

	return Body(contentType)
}

func Body(contentType []byte) Binding {
	switch stringsi.BytesToString(contentType) {
	case consts.ContentTypeForm:
		return FormPost
	case consts.ContentTypeMultipart:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return CustomBody
	}
}

func Bind(c fiber.Ctx, obj interface{}) error {
	tag := binding.Tag
	if data := c.Body(); len(data) > 0 {
		b := Body(c.Request().Header.ContentType())
		err := b.Bind(c, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = b.Name()
	}

	var args mtos.PeekVsSource

	args = append(args, (*uriSource)(c.(*fiber.DefaultCtx)))

	if query := c.Queries(); len(query) > 0 {
		args = append(args, QuerySource(query))
	}
	if headers := c.GetReqHeaders(); len(headers) > 0 {
		args = append(args, binding.HeaderSource(headers))
	}
	err := mtos.MappingByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}

func SetTag(tag string) {
	binding.SetTag(tag)
}

func NewReq[REQ any](c fiber.Ctx) (*REQ, error) {
	req := new(REQ)
	err := Bind(c, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func BindBody(r fiber.Ctx, obj interface{}) error {
	return BindWith(r, obj, CustomBody)
}

// BindQuery is a shortcut for c.BindWith(obj, binding.Query).
func BindQuery(c fiber.Ctx, obj interface{}) error {
	return BindWith(c, obj, Query)
}

func BindHeader(c fiber.Ctx, obj interface{}) error {
	return BindWith(c, obj, Header)
}

// BindWith binds the passed struct pointer using the specified binding engine.
// BindUri binds the passed struct pointer using binding.Uri.
func BindUri(c fiber.Ctx, obj interface{}) error {
	return Uri.Bind(c, obj)
}

func BindWith(c fiber.Ctx, obj interface{}, b Binding) error {
	return b.Bind(c, obj)
}
