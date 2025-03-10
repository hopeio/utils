// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(ctx *gin.Context, obj interface{}) error {
	values := ctx.Request.URL.Query()
	if err := mtos.MapFormByTag(obj, mtos.KVsSource(values), binding.Tag); err != nil {
		return err
	}
	return Validate(obj)
}
