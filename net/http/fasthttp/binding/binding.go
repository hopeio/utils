/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/valyala/fasthttp"
	"net/http"
)

type Binding interface {
	Name() string
	Bind(*fasthttp.RequestCtx, interface{}) error
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
	Header        = headerBinding{}
)

func Default(method, contentType []byte) Binding {
	if stringsi.BytesToString(method) == http.MethodGet {
		return Query
	}

	return Body(contentType)
}

func Body(contentType []byte) Binding {
	switch stringsi.BytesToString(contentType) {

	case binding.MIMEPOSTForm:
		return FormPost
	case binding.MIMEMultipartPOSTForm:
		return FormMultipart
	default:
		return CustomBody
	}
}

func Bind(c *fasthttp.RequestCtx, obj interface{}) error {
	tag := binding.Tag
	if data := c.Request.Body(); len(data) > 0 {
		b := Body(c.Request.Header.ContentType())
		err := b.Bind(c, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = binding.Tag
	}

	var args mtos.PeekVsSource

	if query := c.QueryArgs(); query != nil {
		args = append(args, (*ArgsSource)(query))
	}
	args = append(args, (*HeaderSource)(&c.Request.Header))
	err := mtos.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}

func RegisterBodyBinding(name string, unmarshaller func(data []byte, obj any) error) {
	binding.SetTag(name)
	CustomBody.name = name
	CustomBody.unmarshaller = unmarshaller
}

func SetTag(tag string) {
	binding.SetTag(tag)
}
