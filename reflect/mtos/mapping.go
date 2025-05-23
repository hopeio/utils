// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mtos

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var errUnknownType = errors.New("unknown type")

// Setter tries to set value on a walking by fields of a struct
type Setter interface {
	TrySet(value reflect.Value, field *reflect.StructField, key string, opt SetOptions) (isSet bool, err error)
}

type Setters struct {
	Setters []Setter
}

func (receiver Setters) TrySet(value reflect.Value, field *reflect.StructField, key string, opt SetOptions) (isSet bool, err error) {
	defaultValue := opt.defaultValue
	opt.defaultValue = ""
	for _, arg := range receiver.Setters {
		if arg != nil {
			isSet, err = arg.TrySet(value, field, key, opt)
			if isSet {
				return
			}
		}
	}
	if defaultValue != "" {
		return true, SetValueByStringWithStructField(value, field, defaultValue)
	}
	return
}

func MappingByTag(ptr interface{}, setter Setter, tag string) error {
	_, err := mapping(reflect.ValueOf(ptr), nil, setter, tag)
	return err
}

func mapping(value reflect.Value, field *reflect.StructField, setter Setter, tag string) (bool, error) {
	var tagValue string
	if field != nil {
		tagValue = field.Tag.Get(tag)
	}
	if tagValue == "-" { // just ignoring this field
		return false, nil
	}

	var vKind = value.Kind()

	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		if value.IsNil() {
			isNew = true
			vPtr = reflect.New(value.Type().Elem())
		}
		isSet, err := mapping(vPtr.Elem(), field, setter, tag)
		if err != nil {
			return false, err
		}
		if isNew && isSet {
			value.Set(vPtr)
		}
		return isSet, nil
	}

	if vKind == reflect.Struct {
		tValue := value.Type()

		var isSet bool
		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}
			ok, err := mapping(value.Field(i), &sf, setter, tag)
			if err != nil {
				return false, err
			}
			isSet = isSet || ok
		}
		return isSet, nil
	}

	if field != nil && !field.Anonymous {
		ok, err := tryToSetValue(value, field, setter, tagValue)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}

type SetOptions struct {
	defaultValue string
}

func tryToSetValue(value reflect.Value, field *reflect.StructField, setter Setter, tagValue string) (bool, error) {

	var setOpt SetOptions
	tagValue, opts := head(tagValue, ",")

	if tagValue == "" { // default value is FieldName
		tagValue = field.Name
	}
	if tagValue == "" { // when field is "emptyField" variable
		return false, nil
	}

	var opt string
	for len(opts) > 0 {
		opt, opts = head(opts, ",")

		if k, v := head(opt, "="); k == "default" {
			setOpt.defaultValue = v
			break
		}
	}

	return setter.TrySet(value, field, tagValue, setOpt)
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField *reflect.StructField, value reflect.Value) error {
	timeFormat := time.RFC3339
	l := time.Local
	if structField != nil {
		timeFormat = structField.Tag.Get("time_format")
		switch tf := strings.ToLower(timeFormat); tf {
		case "unix", "unixnano":
			tv, err := strconv.ParseInt(val, 10, 0)
			if err != nil {
				return err
			}

			d := time.Duration(1)
			if tf == "unixnano" {
				d = time.Second
			}

			t := time.Unix(tv/int64(d), tv%int64(d))
			value.Set(reflect.ValueOf(t))
			return nil

		}

		if val == "" {
			value.Set(reflect.ValueOf(time.Time{}))
			return nil
		}

		if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
			l = time.UTC
		}

		if locTag := structField.Tag.Get("time_location"); locTag != "" {
			loc, err := time.LoadLocation(locTag)
			if err != nil {
				return err
			}
			l = loc
		}
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return nil
}

func setArray(vals []string, value reflect.Value, field *reflect.StructField) error {
	for i, s := range vals {
		err := SetValueByStringWithStructField(value.Index(i), field, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func setSlice(vals []string, value reflect.Value, field *reflect.StructField) error {
	slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))
	err := setArray(vals, slice, field)
	if err != nil {
		return err
	}
	value.Set(slice)
	return nil
}

func setTimeDuration(val string, value reflect.Value) error {
	d, err := time.ParseDuration(val)
	if err != nil {
		return err
	}
	value.Set(reflect.ValueOf(d))
	return nil
}

func head(str, sep string) (head string, tail string) {
	idx := strings.Index(str, sep)
	if idx < 0 {
		return str, ""
	}
	return str[:idx], str[idx+len(sep):]
}

type CanSetter interface {
	Setter
	HasValue(key string) bool
}

type CanSetters []CanSetter

func (args CanSetters) TrySet(value reflect.Value, field *reflect.StructField, key string, opt SetOptions) (isSet bool, err error) {
	for _, arg := range args {
		if arg.HasValue(key) {
			return arg.TrySet(value, field, key, opt)
		}
	}
	return false, nil
}
