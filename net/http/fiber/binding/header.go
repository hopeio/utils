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

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req fiber.Ctx, obj interface{}) error {
	if err := mtos.MappingByTag(obj, (*HeaderSource)(&req.Request().Header), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}
