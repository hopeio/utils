/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"

	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
)

func SetTag(tag string) {
	if tag != "" {
		binding.Tag = tag
	}
}

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*gin.Context, interface{}) error
}

// BindingBody adds BindBody method to Binding. BindBody is similar with GinBind,
// but it reads the body from supplied bytes instead of req.Body.
type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
	CustomBody    = bodyBinding{name: "json", unmarshaller: json.Unmarshal}
)

// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
func Default(method string, contentType string) Binding {
	if method == http.MethodGet {
		return Query
	}

	return Body(contentType)
}

func Body(contentType string) Binding {
	switch contentType {
	case MIMEPOSTForm:
		return FormPost
	case MIMEMultipartPOSTForm:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return CustomBody
	}
}

func Validate(obj interface{}) error {
	return binding.Validator.ValidateStruct(obj)
}

func Bind(c *gin.Context, obj interface{}) error {
	tag := binding.Tag
	if c.Request.Body != nil && c.Request.ContentLength != 0 {
		b := Body(c.ContentType())
		err := b.Bind(c, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = b.Name()
	}

	var args mtos.PeekVsSource
	if len(c.Params) > 0 {
		args = append(args, uriSource(c.Params))
	}
	if len(c.Request.URL.RawQuery) > 0 {
		args = append(args, mtos.KVsSource(c.Request.URL.Query()))
	}
	if len(c.Request.Header) > 0 {
		args = append(args, binding.HeaderSource(c.Request.Header))
	}
	err := mtos.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}

func RegisterBodyBinding(name string, unmarshaller func(data []byte, obj any) error) {
	CustomBody.name = name
	CustomBody.unmarshaller = unmarshaller
}

func RegisterBodyBindingByDecoder(name string, newDecoder func(io.Reader) encoding.Decoder) {
	binding.SetTag(name)
	CustomBody.name = name
	CustomBody.decoder = newDecoder
}
