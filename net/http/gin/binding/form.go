// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
)

const defaultMemory = 32 << 20

type formPostBinding struct{}
type formMultipartBinding struct{}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseForm(); err != nil {
		return err
	}

	args := mtos.PeekVsSource{mtos.KVsSource(ctx.Request.PostForm)}
	if err := mtos.MappingByTag(obj, args, binding.Tag); err != nil {
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
	if err := mtos.MappingByTag(obj, (*binding.MultipartSource)(ctx.Request), binding.Tag); err != nil {
		return err
	}

	return Validate(obj)
}
