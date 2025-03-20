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
	"reflect"
)

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(c fiber.Ctx, obj interface{}) error {
	if err := mtos.MappingByTag(obj, (*uriSource)(c.(*fiber.DefaultCtx)), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}

type uriSource fiber.DefaultCtx

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *uriSource) TrySet(value reflect.Value, field *reflect.StructField, key string, opt mtos.SetOptions) (isSet bool, err error) {
	return mtos.SetValueByKVsWithStructField(value, field, form, key, opt)
}

func (form *uriSource) Peek(key string) ([]string, bool) {
	v := (*fiber.DefaultCtx)(form).Params(key)
	return []string{v}, v != ""
}

func (form *uriSource) HasValue(key string) bool {
	v := (*fiber.DefaultCtx)(form).Params(key)
	return v != ""
}
