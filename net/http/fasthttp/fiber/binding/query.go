/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	fbinding "github.com/hopeio/utils/net/http/fasthttp/binding"
	"github.com/hopeio/utils/reflect/mtos"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	values := ctx.Request().URI().QueryArgs()
	if err := mtos.MapFormByTag(obj, (*fbinding.ArgsSource)(values), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
