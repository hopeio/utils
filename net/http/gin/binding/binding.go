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

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
	CustomBody    = &bodyBinding{name: "json", unmarshaller: json.Unmarshal}
)

func Validate(obj interface{}) error {
	return binding.Validator.ValidateStruct(obj)
}

func Bind(c *gin.Context, obj interface{}) error {
	tag := binding.Tag
	var args mtos.CanSetters
	if c.Request.Body != nil && c.Request.ContentLength != 0 {
		switch c.ContentType() {
		case MIMEPOSTForm:
			err := c.Request.ParseForm()
			if err != nil {
				return err
			}
			args = append(args, mtos.KVsSource(c.Request.PostForm))
			tag = "form"
		case MIMEMultipartPOSTForm:
			err := c.Request.ParseMultipartForm(defaultMemory)
			if err != nil {
				return err
			}
			args = append(args, (*binding.MultipartSource)(c.Request.MultipartForm))
			tag = "form"
		default: // case MIMEPOSTForm:
			err := CustomBody.Bind(c, obj)
			if err != nil {
				return fmt.Errorf("body bind error: %w", err)
			}
			tag = CustomBody.Name()
		}
	}

	if len(c.Params) > 0 {
		args = append(args, uriSource(c.Params))
	}
	if len(c.Request.URL.RawQuery) > 0 {
		args = append(args, mtos.KVsSource(c.Request.URL.Query()))
	}
	if len(c.Request.Header) > 0 {
		args = append(args, binding.HeaderSource(c.Request.Header))
	}
	if len(args) > 0 {
		err := mtos.MappingByTag(obj, args, tag)
		if err != nil {
			return fmt.Errorf("args bind error: %w", err)
		}
	}
	return nil
}

func NewReq[REQ any](c *gin.Context) (*REQ, error) {
	req := new(REQ)
	err := Bind(c, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func BindBody(c *gin.Context, obj interface{}) error {
	return BindWith(c, obj, CustomBody)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func BindQuery(c *gin.Context, obj interface{}) error {
	return BindWith(c, obj, Query)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func BindWith(c *gin.Context, obj interface{}, b Binding) error {
	if err := b.Bind(c, obj); err != nil {
		return err
	}
	if err := Validate(obj); err != nil {
		return err
	}
	return nil
}

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func BindUri(r *gin.Context, obj interface{}) error {
	return Uri.Bind(r, obj)
}
