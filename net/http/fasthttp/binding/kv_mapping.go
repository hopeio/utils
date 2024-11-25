/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/hopeio/utils/reflect/mtos"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/valyala/fasthttp"
	"reflect"
)

type ArgsSource fasthttp.Args

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *ArgsSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt mtos.SetOptions) (isSet bool, err error) {
	return mtos.SetByKVs(value, field, form, tagValue, opt)
}

func (form *ArgsSource) Peek(key string) ([]string, bool) {
	v := stringsi.BytesToString((*fasthttp.Args)(form).Peek(key))
	return []string{v}, v != ""
}

type CtxSource fasthttp.RequestCtx

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *CtxSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt mtos.SetOptions) (isSet bool, err error) {
	return mtos.SetByKVs(value, field, form, tagValue, opt)
}

func (form *CtxSource) Peek(key string) ([]string, bool) {
	v := (*fasthttp.RequestCtx)(form).UserValue(key).(string)
	return []string{v}, v != ""
}

type HeaderSource fasthttp.RequestHeader

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *HeaderSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt mtos.SetOptions) (isSet bool, err error) {
	return mtos.SetByKVs(value, field, form, tagValue, opt)
}

func (form *HeaderSource) Peek(key string) ([]string, bool) {
	v := stringsi.BytesToString((*fasthttp.RequestHeader)(form).Peek(key))
	return []string{v}, v != ""
}
