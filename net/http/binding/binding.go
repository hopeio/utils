/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package binding

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/net/http/consts"
	"github.com/hopeio/utils/reflect/mtos"
	"github.com/hopeio/utils/validation/validator"
	"io"
	"net/http"
	"reflect"
	"sync"
)

// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v8.18.2
// under the hood.
var Validator = validator.DefaultValidator

var (
	DefaultMemory    int64                   = 32 << 20
	BodyUnmarshaller func([]byte, any) error = json.Unmarshal
	CommonTag                                = "json"
)

func Validate(obj interface{}) error {
	return Validator.ValidateStruct(obj)
}

var defaultTags = []string{"uri", "path", "query", "header", "form"}

type Source interface {
	Uri() mtos.Setter
	Query() mtos.Setter
	Header() mtos.Setter
	Form() mtos.Setter
	BodyBind(obj any) error
}

type Field struct {
	Tag      string
	TagValue string
	Index    int
	Field    *reflect.StructField
}

var cache = sync.Map{}

func Bind(r *http.Request, obj any) error {
	return CommonBind(RequestSource{r}, obj)
}

func CommonBind(s Source, obj any) error {
	value := reflect.ValueOf(obj).Elem()
	typ := value.Type()
	err := s.BodyBind(obj)
	if err != nil {
		return err
	}
	uriSetter, querySetter, headerSetter, formSetter := s.Uri(), s.Query(), s.Header(), s.Form()
	commonSetter := mtos.Setters{Setters: []mtos.Setter{uriSetter, querySetter, headerSetter}}
	if fields, ok := cache.Load(typ); ok {
		for _, field := range fields.([]Field) {
			var setter mtos.Setter
			switch field.Tag {
			case "uri", "path":
				setter = uriSetter
			case "query":
				setter = querySetter
			case "header":
				setter = headerSetter
			case "form":
				setter = formSetter
			case CommonTag:
				setter = commonSetter
			}
			if setter == nil {
				continue
			}
			_, err = setter.TrySet(value.Field(field.Index), field.Field, field.TagValue, mtos.SetOptions{})
			if err != nil {
				return err
			}
		}
		return Validate(obj)
	}
	var fields []Field
	for i := 0; i < value.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}
		var tagValue string
		var tag string
		for _, tag = range defaultTags {
			tagValue = sf.Tag.Get(tag)
			if tagValue != "" && tagValue != "-" {
				break
			}
		}
		if tagValue == "" || tagValue == "-" {
			tagValue = sf.Tag.Get(CommonTag)
			if tagValue != "" && tagValue != "-" {
				fields = append(fields, Field{
					Tag:      CommonTag,
					TagValue: tagValue,
					Index:    i,
					Field:    &sf,
				})
				_, err = commonSetter.TrySet(value.Field(i), &sf, tagValue, mtos.SetOptions{})
				if err != nil {
					return err
				}
			}
			continue
		}

		var setter mtos.Setter
		switch tag {
		case "uri", "path":
			setter = uriSetter
		case "query":
			setter = querySetter
		case "header":
			setter = headerSetter
		case "form":
			setter = formSetter
		}
		fields = append(fields, Field{
			Tag:      tag,
			TagValue: tagValue,
			Index:    i,
			Field:    &sf,
		})
		if setter == nil {
			continue
		}
		_, err = setter.TrySet(value.Field(i), &sf, tagValue, mtos.SetOptions{})
		if err != nil {
			return err
		}

	}
	cache.Store(typ, fields)
	return Validate(obj)
}

type RequestSource struct {
	*http.Request
}

func (s RequestSource) Uri() mtos.Setter {
	return (*UriSource)(s.Request)
}

func (s RequestSource) Query() mtos.Setter {
	return (mtos.KVsSource)(s.URL.Query())
}

func (s RequestSource) Header() mtos.Setter {
	return (HeaderSource)(s.Request.Header)
}

func (s RequestSource) Form() mtos.Setter {
	contentType := s.Request.Header.Get(consts.HeaderContentType)
	if contentType == consts.ContentTypeForm {
		err := s.ParseForm()
		if err != nil {
			return nil
		}
		return (mtos.KVsSource)(s.PostForm)
	}
	if contentType == consts.ContentTypeMultipart {
		err := s.ParseMultipartForm(DefaultMemory)
		if err != nil {
			return nil
		}
		return (*MultipartSource)(s.MultipartForm)
	}
	return nil
}

func (s RequestSource) BodyBind(obj any) error {
	if s.Method == http.MethodGet {
		return nil
	}
	data, err := io.ReadAll(s.Body)
	if err != nil {
		return fmt.Errorf("read body error: %w", err)
	}
	return BodyUnmarshaller(data, obj)
}
