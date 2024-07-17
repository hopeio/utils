// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (xmlBinding) Bind(ctx *gin.Context, obj interface{}) error {
	return binding.DecodeXML(ctx.Request.Body, obj)
}
