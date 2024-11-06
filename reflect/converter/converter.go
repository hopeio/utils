// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"encoding"
	"errors"
	constraintsi "github.com/hopeio/utils/types/constraints"
	"github.com/spf13/cast"
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

type StringConverter func(string) any
type StringConverterE func(string) (any, error)

func (c StringConverterE) IgnoreError() StringConverter {
	if c == nil {
		return nil
	}
	return func(value string) any {
		r, _ := c(value)
		return r
	}
}

var (
	invalidValue = reflect.Value{}
)

// Default converters for basic types.
/*var stringConverterMaps = map[reflect.Kind]StringConverterE{
	reflect.Bool:    stringConvertBool,
	reflect.Float32: stringConvertFloat32,
	reflect.Float64: stringConvertFloat64,
	reflect.Int:     stringConvertInt,
	reflect.Int8:    stringConvertInt8,
	reflect.Int16:   stringConvertInt16,
	reflect.Int32:   stringConvertInt32,
	reflect.Int64:   stringConvertInt64,
	reflect.String:  stringConvertString,
	reflect.Uint:    stringConvertUint,
	reflect.Uint8:   stringConvertUint8,
	reflect.Uint16:  stringConvertUint16,
	reflect.Uint32:  stringConvertUint32,
	reflect.Uint64:  stringConvertUint64,
}*/

// Deprecated: unsupported string slices array map
var stringConverterArrays = []StringConverterE{
	reflect.Invalid: nil,
	reflect.Bool:    stringConvertBool,
	reflect.Int:     stringConvertInt,
	reflect.Int8:    stringConvertInt8,
	reflect.Int16:   stringConvertInt16,
	reflect.Int32:   stringConvertInt32,
	reflect.Int64:   stringConvertInt64,
	reflect.Uint:    stringConvertUint,
	reflect.Uint8:   stringConvertUint8,
	reflect.Uint16:  stringConvertUint16,
	reflect.Uint32:  stringConvertUint32,
	reflect.Uint64:  stringConvertUint64,
	reflect.Uintptr: nil,
	reflect.Float32: stringConvertFloat32,
	reflect.Float64: stringConvertFloat64,
}

func GetStringConverter(kind reflect.Kind) StringConverter {
	if kind == reflect.String {
		return func(value string) any {
			return value
		}
	}
	return GetStringConverterE(kind).IgnoreError()
}

func GetStringConverterE(kind reflect.Kind) StringConverterE {
	if kind == reflect.String {
		return stringConvertString
	}
	return stringConverterArrays[kind]
}

const (
	array        = 100
	ArrayBool    = reflect.Bool + array
	ArrayInt     = reflect.Int + array
	ArrayInt8    = reflect.Int8 + array
	ArrayInt16   = reflect.Int16 + array
	ArrayInt32   = reflect.Int32 + array
	ArrayInt64   = reflect.Int64 + array
	ArrayUint    = reflect.Uint + array
	ArrayUint8   = reflect.Uint8 + array
	ArrayUint16  = reflect.Uint16 + array
	ArrayUint32  = reflect.Uint32 + array
	ArrayUint64  = reflect.Uint64 + array
	ArrayFloat32 = reflect.Float32 + array
	ArrayFloat64 = reflect.Float64 + array
	ArrayString  = reflect.String + array
)

const (
	slice        = 200
	SliceBool    = reflect.Bool + slice
	SliceInt     = reflect.Int + slice
	SliceInt8    = reflect.Int8 + slice
	SliceInt16   = reflect.Int16 + slice
	SliceInt32   = reflect.Int32 + slice
	SliceInt64   = reflect.Int64 + slice
	SliceUint    = reflect.Uint + slice
	SliceUint8   = reflect.Uint8 + slice
	SliceUint16  = reflect.Uint16 + slice
	SliceUint32  = reflect.Uint32 + slice
	SliceUint64  = reflect.Uint64 + slice
	SliceFloat32 = reflect.Float32 + slice
	SliceFloat64 = reflect.Float64 + slice
	SliceString  = reflect.String + slice
)

func stringConvertBool(value string) (any, error) {
	if value == "on" {
		return true, nil
	}
	return strconv.ParseBool(value)
}

func StringConvertBool(value string) (bool, error) {
	if value == "on" {
		return true, nil
	}
	return strconv.ParseBool(value)
}

func stringConvertFloat32(value string) (any, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func StringConvertFloat32(value string) (float32, error) {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func stringConvertFloat64(value string) (any, error) {
	return strconv.ParseFloat(value, 64)
}

func StringConvertFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func stringConvertInt(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func stringConvertInt8(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

func stringConvertInt16(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func stringConvertInt32(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func stringConvertInt64(value string) (any, error) {
	return strconv.ParseInt(value, 10, 64)
}

func stringConvertString(value string) (any, error) {
	return value, nil
}

func StringConvertIntFor[T constraints.Signed](value string) (T, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

// TODO
func stringConvertArray(value string) (any, error) {
	return value, nil
}

func stringConvertUint(value string) (any, error) {
	u, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(u), nil
}

func stringConvertUint8(value string) (any, error) {
	u, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(u), nil
}

func stringConvertUint16(value string) (any, error) {
	u, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(u), nil
}

func stringConvertUint32(value string) (any, error) {
	u, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(u), nil
}

func stringConvertUint64(value string) (any, error) {
	return strconv.ParseUint(value, 10, 64)
}

func StringConvertUintFor[T constraints.Unsigned](value string) (T, error) {
	i, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func CastInt64(v any) int64 {
	return cast.ToInt64(v)
}

func StringConvert(kind reflect.Kind, value string) (any, error) {
	if kind == reflect.String {
		return value, nil
	}
	converter := stringConverterArrays[kind]
	if converter != nil {
		return converter(value)
	}
	return nil, errors.New("unsupported kind")
}

func StringConvertFor[T any](value string) (T, error) {
	kind := reflect.TypeFor[T]().Kind()
	if kind == reflect.String {
		return any(value).(T), nil
	}
	converter := stringConverterArrays[kind]
	if converter != nil {
		if v, err := converter(value); err != nil {
			return *new(T), err
		} else {
			return v.(T), nil
		}
	} else {
		var v T
		a, ap := any(v), any(&v)
		vv, ok := a.(encoding.TextUnmarshaler)
		if !ok {
			vv, ok = ap.(encoding.TextUnmarshaler)
		}
		if ok {
			err := vv.UnmarshalText([]byte(value))
			if err != nil {
				return v, err
			}
		}
	}
	return *new(T), errors.New("unsupported kind")
}

func StringConvertNumberFor[T constraintsi.Number](value string) (T, error) {
	var v T
	a, ap := any(v), any(&v)
	switch vv := a.(type) {
	case encoding.TextUnmarshaler:
		err := vv.UnmarshalText([]byte(value))
		if err != nil {
			return v, err
		}
		return v, nil
	case int, int8, int16, int32, int64:
		i, err := strconv.ParseInt(value, 10, int(unsafe.Sizeof(v))*8)
		if err != nil {
			return v, err
		}
		return T(i), nil
	case uint, uint8, uint16, uint32, uint64:
		i, err := strconv.ParseUint(value, 10, int(unsafe.Sizeof(v))*8)
		if err != nil {
			return v, err
		}
		return T(i), nil
	case float64, float32:
		f, err := strconv.ParseFloat(value, int(unsafe.Sizeof(v))*8)
		if err != nil {
			return v, err
		}
		return T(f), nil
	}
	switch vv := ap.(type) {
	case encoding.TextUnmarshaler:
		err := vv.UnmarshalText([]byte(value))
		if err != nil {
			return v, err
		}
		return v, nil
	}
	return *new(T), errors.New("unsupported type")
}

func String(value reflect.Value) string {
	v := value.Interface()
	if t, ok := v.(encoding.TextMarshaler); ok {
		s, _ := t.MarshalText()
		return string(s)
	}

	kind := value.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Pointer, reflect.UnsafePointer:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.String:
		return value.String()
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float64, reflect.Float32:
		return strconv.FormatFloat(value.Float(), 'g', -1, 64)
	case reflect.Array, reflect.Slice:
		var strs []string
		for i := 0; i < value.Len(); i++ {
			strs = append(strs, String(value.Index(i)))
		}
		return strings.Join(strs, ",")
	}
	return ""
}

func StringFor[T any](t T) string {
	v := reflect.ValueOf(t)
	return String(v)
}
