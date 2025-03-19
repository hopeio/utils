/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/reflect/mtos"
	"github.com/hopeio/utils/validation/validator"
	"net/http"
)

var defaultTags = []string{"uri", "path", "query", "header", "form", "json"}

var Tag = "json"

func SetTag(tag string) {
	if tag != "" {
		Tag = tag
	}
	mtos.SetAliasTag(tag)
}

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v8.18.2
// under the hood.
var Validator = validator.DefaultValidator

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	Uri    = uriBinding{}
	Query  = queryBinding{}
	Header = headerBinding{}

	CustomBody    = bodyBinding{name: "json", unmarshaller: json.Unmarshal}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
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
	case consts.ContentTypeForm:
		return FormPost
	case consts.ContentTypeMultipart:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return CustomBody
	}
}

func Validate(obj interface{}) error {
	return Validator.ValidateStruct(obj)
}

func Bind(r *http.Request, obj interface{}) error {
	tag := Tag
	if r.Body != nil && r.ContentLength != 0 {
		b := Body(r.Header.Get("Content-Type"))
		err := b.Bind(r, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = b.Name()
	}

	var args mtos.PeekVsSource
	if r.Pattern != "" {
		args = append(args, (*UriSource)(r))
	}
	if len(r.URL.RawQuery) > 0 {
		args = append(args, mtos.KVsSource(r.URL.Query()))
	}
	if len(r.Header) > 0 {
		args = append(args, HeaderSource(r.Header))
	}
	err := mtos.MappingByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}

func NewReq[REQ any](r *http.Request) (*REQ, error) {
	req := new(REQ)
	err := Bind(r, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func BindBody(r *http.Request, obj interface{}) error {
	return BindWith(r, obj, CustomBody)
}

func BindHeader(r *http.Request, obj interface{}) error {
	return Header.Bind(r, obj)
}

// BindQuery is a shortcut for c.BindWith(obj, binding.Query).
func BindQuery(r *http.Request, obj interface{}) error {
	return BindWith(r, obj, Query)
}

func BindUri(r *http.Request, obj interface{}) error {
	return Uri.Bind(r, obj)
}

// BindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func BindWith(r *http.Request, obj interface{}, b Binding) error {
	return b.Bind(r, obj)
}
