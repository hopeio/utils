// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultMemory = 32 << 20

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	args := mtos.PeekVsSource{mtos.KVsSource(ctx.Request.Form)}
	if err := mtos.MapFormByTag(obj, args, binding.Tag); err != nil {
		return err
	}
	return Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseForm(); err != nil {
		return err
	}

	args := mtos.PeekVsSource{mtos.KVsSource(ctx.Request.Form)}
	if err := mtos.MapFormByTag(obj, args, binding.Tag); err != nil {
		return err
	}
	return Validate(obj)
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := mtos.MapFormByTag(obj, (*binding.MultipartSource)(ctx.Request), binding.Tag); err != nil {
		return err
	}

	return Validate(obj)
}
