// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (yamlBinding) Bind(ctx *gin.Context, obj interface{}) error {
	return binding.DecodeYAML(ctx.Request.Body, obj)
}
