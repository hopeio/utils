/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/encoding"
	"io"
)

type bodyBinding struct {
	name         string
	unmarshaller func([]byte, any) error
	newDecoder   func(io.Reader) encoding.Decoder
}

func (b bodyBinding) Name() string {
	return b.name
}

func (b bodyBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	return b.unmarshaller(ctx.Body(), obj)
}
