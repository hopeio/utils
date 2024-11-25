/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	"github.com/valyala/fasthttp"
	"reflect"
)

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

	return mtos.SetByKVs(value, field, mtos.KVsSource(form.Value), key, opt)
}
