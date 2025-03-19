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
	"github.com/valyala/fasthttp"
	"reflect"
)

type formPostBinding struct{}
type formMultipartBinding struct{}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	if err := mtos.MappingByTag(obj, (*MultipartRequest)(ctx.Request()), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	if err := mtos.MappingByTag(obj, (*ArgsSource)(ctx.Request().PostArgs()), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}

type MultipartRequest fasthttp.Request

var _ mtos.Setter = (*MultipartRequest)(nil)

// TrySet tries to set a value by the multipart request with the binding a form file
func (r *MultipartRequest) TrySet(value reflect.Value, field *reflect.StructField, key string, opt mtos.SetOptions) (isSet bool, err error) {
	req := (*fasthttp.Request)(r)
	form, err := req.MultipartForm()
	if err != nil {
		return false, err
	}
	if files := form.File[key]; len(files) != 0 {
		return binding.SetByMultipartFormFile(value, field, files)
	}

	return mtos.SetValueByKVsWithStructField(value, field, mtos.KVsSource(form.Value), key, opt)
}
