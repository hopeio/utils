// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
	"io"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(ctx *gin.Context, obj interface{}) error {
	buf, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	return binding.DecodeProtobuf(buf, obj)
}
