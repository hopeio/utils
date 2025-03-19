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
	"strings"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	values := ctx.Request().URI().QueryArgs()
	if err := mtos.MappingByTag(obj, (*ArgsSource)(values), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}

type QuerySource map[string]string

func (q QuerySource) Peek(key string) ([]string, bool) {
	v, ok := q[key]
	if strings.Contains(v, ",") {
		return strings.Split(v, ","), true
	}
	return []string{v}, ok
}
