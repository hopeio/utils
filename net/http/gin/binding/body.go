/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/encoding"
	"io"
)

type bodyBinding struct {
	name         string
	unmarshaller func([]byte, any) error
	decoder      func(io.Reader) encoding.Decoder
}

func (b bodyBinding) Name() string {
	return b.name
}

func (b bodyBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if ctx == nil || ctx.Request.Body == nil {
		return fmt.Errorf("invalid request")
	}
	if b.decoder != nil {
		return b.decoder(ctx.Request.Body).Decode(obj)
	}
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	return b.unmarshaller(data, obj)
}

func (b *bodyBinding) RegisterUnmarshaller(name string, unmarshaller func(data []byte, obj any) error) {
	CustomBody.name = name
	CustomBody.unmarshaller = unmarshaller
}

func (b *bodyBinding) RegisterDecoder(name string, decoder func(io.Reader) encoding.Decoder) {
	CustomBody.name = name
	CustomBody.decoder = decoder
}
