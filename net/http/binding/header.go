/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/hopeio/utils/reflect/mtos"
	"net/http"
	"net/textproto"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {

	if err := mtos.Decode(obj, req.Header); err != nil {
		return err
	}

	return Validate(obj)
}

func MapHeader(ptr interface{}, h map[string][]string) error {
	return mtos.MappingByTag(ptr, HeaderSource(h), "header")
}

type HeaderSource map[string][]string

var _ mtos.Setter = HeaderSource(nil)

func (hs HeaderSource) Peek(key string) ([]string, bool) {
	v, ok := hs[textproto.CanonicalMIMEHeaderKey(key)]
	return v, ok
}

func (hs HeaderSource) HasValue(key string) bool {
	_, ok := hs[textproto.CanonicalMIMEHeaderKey(key)]
	return ok
}
func (hs HeaderSource) TrySet(value reflect.Value, field *reflect.StructField, key string, opt mtos.SetOptions) (isSet bool, err error) {
	return mtos.SetValueByKVsWithStructField(value, field, hs, key, opt)
}
