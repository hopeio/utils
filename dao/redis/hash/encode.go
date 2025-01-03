/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package hash

import (
	"encoding"
	"reflect"
)

func Marshal(v interface{}) []interface{} {
	uValue := reflect.ValueOf(v).Elem()
	uType := uValue.Type()
	var redisArgs = make([]interface{}, 0, uValue.NumField())
	for i := range uValue.NumField() {
		redisArgs = append(redisArgs, uType.Field(i).Name, uValue.Field(i).Interface())
	}
	return redisArgs
}

type encodeState struct {
	strings []interface{}
}

func (e *encodeState) encode(key string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.String:
		e.strings = append(e.strings, key, v.Interface())
	case reflect.Interface, reflect.Ptr:
		e.structEncoder(key, v.Elem())
	case reflect.Struct:
		e.structEncoder(key, v)
		/*	case reflect.Map:
				return newMapEncoder(t)
			case reflect.Slice:
				return newSliceEncoder(t)
			case reflect.Array:
				return newArrayEncoder(t)
			default:
				return unsupportedTypeEncoder*/
	}
}

var textMarshallerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()

func (e *encodeState) structEncoder(key string, v reflect.Value) {
	t := v.Type()
	if t.Implements(textMarshallerType) {
		m := v.Interface().(encoding.TextMarshaler)
		bytes, _ := m.MarshalText()
		e.strings = append(e.strings, key, string(bytes))
		return
	}
	if key != "" {
		key += "."
	}

	for i := range v.NumField() {
		field := t.Field(i).Name
		if 'A' <= field[0] && field[0] <= 'Z' {
			e.encode(key+field, v.Field(i))
		}
	}
}

func (e *encodeState) mapEncoder(v reflect.Value) {

}
