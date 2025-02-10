/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"github.com/hopeio/utils/reflect/mtos"
	"net/http"
	"reflect"
)

// support go 1.22
type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(req *http.Request, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*UriSource)(req), "uri"); err != nil {
		return err
	}
	return Validate(obj)

}

type UriSource http.Request

var _ mtos.Setter = (*UriSource)(nil)

func (req *UriSource) Peek(key string) ([]string, bool) {
	if req.Pattern == "" {
		return nil, false
	}
	v := (*http.Request)(req).PathValue(key)
	return []string{v}, v != ""
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (req *UriSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt mtos.SetOptions) (isSet bool, err error) {
	return mtos.SetValueByKVsWithStructField(value, field, req, tagValue, opt)
}
