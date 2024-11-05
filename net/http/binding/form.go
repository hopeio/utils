// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"errors"
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/reflect/mtos"
	"mime/multipart"
	"net/http"
	"reflect"
)

const defaultMemory = 32 << 20

type formPostBinding struct{}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := mtos.Decode(obj, req.PostForm); err != nil {
		return err
	}
	return Validate(obj)
}

type formMultipartBinding struct{}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := encoding.MapFormByTag(obj, (*MultipartSource)(req), Tag); err != nil {
		return err
	}

	return Validate(obj)
}

type MultipartSource http.Request

var _ encoding.Setter = (*MultipartSource)(nil)

// TrySet tries to set a value by the multipart request with the binding a form file
func (r *MultipartSource) TrySet(value reflect.Value, field *reflect.StructField, key string, opt encoding.SetOptions) (isSet bool, err error) {
	if files := r.MultipartForm.File[key]; len(files) != 0 {
		return SetByMultipartFormFile(value, field, files)
	}

	return encoding.SetByKVs(value, field, encoding.KVsSource(r.MultipartForm.Value), key, opt)
}

func SetByMultipartFormFile(value reflect.Value, field *reflect.StructField, files []*multipart.FileHeader) (isSet bool, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		switch value.Interface().(type) {
		case *multipart.FileHeader:
			value.Set(reflect.ValueOf(files[0]))
			return true, nil
		}
	case reflect.Struct:
		switch value.Interface().(type) {
		case multipart.FileHeader:
			value.Set(reflect.ValueOf(*files[0]))
			return true, nil
		}
	case reflect.Slice:
		slice := reflect.MakeSlice(value.Type(), len(files), len(files))
		isSet, err = setArrayOfMultipartFormFiles(slice, field, files)
		if err != nil || !isSet {
			return isSet, err
		}
		value.Set(slice)
		return true, nil
	case reflect.Array:
		return setArrayOfMultipartFormFiles(value, field, files)
	}
	return false, errors.New("unsupported field type for multipart.FileHeader")
}

func setArrayOfMultipartFormFiles(value reflect.Value, field *reflect.StructField, files []*multipart.FileHeader) (isSet bool, err error) {
	if value.Len() != len(files) {
		return false, errors.New("unsupported len of array for []*multipart.FileHeader")
	}
	for i := range files {
		setted, err := SetByMultipartFormFile(value.Index(i), field, files[i:i+1])
		if err != nil || !setted {
			return setted, err
		}
	}
	return true, nil
}
