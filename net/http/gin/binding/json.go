// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if ctx == nil || ctx.Request.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.DecodeJSON(ctx.Request.Body, obj)
}
