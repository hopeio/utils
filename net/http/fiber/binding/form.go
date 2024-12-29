/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
)

type formPostBinding struct{}
type formMultipartBinding struct{}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*MultipartRequest)(ctx.Request()), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*ArgsSource)(ctx.Request().PostArgs()), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
