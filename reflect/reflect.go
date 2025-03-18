/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

// TODO:
func containType() {

}

// SetSubField 设置字段值
// 参数均为指针,dst类型为src field的类型
func SetSubField(src any, sub any) bool {
	srcValue := reflect.ValueOf(src).Elem()
	subValue := reflect.ValueOf(sub).Elem()
	return SetSubFieldValue(srcValue, subValue)
}

// SetSubFieldValue 设置字段值
// subValue为srcValue field的类型
func SetSubFieldValue(srcValue reflect.Value, subValue reflect.Value) bool {
	for i := range srcValue.NumField() {
		if srcValue.Field(i).Type() == subValue.Type() {
			srcValue.Field(i).Set(subValue)
			return true
		}
	}
	return false
}

// CopyFieldValueByType 根据类型复制字段
// 参数均为指针,sub类型为src field的类型
func CopyFieldValueByType(src any, sub any) bool {
	srcValue := reflect.ValueOf(src).Elem()
	dstValue := reflect.ValueOf(sub).Elem()
	for i := range srcValue.NumField() {
		if srcValue.Field(i).Type() == dstValue.Type() {
			dstValue.Set(srcValue.Field(i))
			return true
		}
	}
	return false
}

func SetField(structValue reflect.Value, name string, value any) error {

	fieldValue := structValue.FieldByName(name)
	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj ", name)
	}

	if !fieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value ", name)
	}

	fieldType := fieldValue.Type()
	val := reflect.ValueOf(value)

	valTypeKind := val.Type().Kind()
	fieldTypeKind := fieldType.Kind()
	if fieldType != val.Type() && val.CanConvert(fieldType) {
		val = val.Convert(fieldType)
	} else {
		return fmt.Errorf("provided value type %s didn't match obj field type %s", valTypeKind, fieldTypeKind)
	}
	fieldValue.Set(val)
	return nil
}

// Copy 复制结构体,浅表复制
func CopyStruct(src any, dest any) error {
	valueOfS := reflect.ValueOf(src)
	typeOfT := reflect.TypeOf(dest)

	if valueOfS.Kind() == reflect.Ptr {
		valueOfS = valueOfS.Elem()
	}

	if valueOfS.Kind() != reflect.Struct {
		return errors.New("src is not a ptr or struct")
	}

	if typeOfT.Kind() != reflect.Ptr {
		return errors.New("dest is not a ptr")
	}

	typeOfT = typeOfT.Elem()
	valueOfT := reflect.ValueOf(dest).Elem()

	for i := range typeOfT.NumField() {
		// 获取每个成员的结构体字段值
		fieldType := typeOfT.Field(i)
		// 赋值
		valueOfT.Field(i).Set(valueOfS.FieldByName(fieldType.Name))
	}

	return nil
}

func CanCast(t1, t2 reflect.Type, strict bool) bool {
	t1kind, t2kind := t1.Kind(), t2.Kind()
	if strict {
		if (t1 == t2) || (t1kind == reflect.String && t2kind == reflect.String) || (t1kind <= reflect.Complex128 && t2kind <= reflect.Complex128 && t1kind == t2kind) {
			return true
		}
	} else {
		if t1kind <= reflect.Float64 && t2kind <= reflect.Float64 {
			return true
		}
		if (t1kind == reflect.Complex64 || t1kind == reflect.Complex128) && (t2kind == reflect.Complex64 || t2kind == reflect.Complex128) {
			return true
		}
	}

	switch t1kind {
	case reflect.String:
		return t1kind == t2kind
	case reflect.Ptr, reflect.Array, reflect.Chan, reflect.Slice, reflect.Map:
		if t1kind == reflect.Map {
			if !CanCast(t1.Key(), t2.Key(), true) {
				return false
			}
		}
		if t1kind == reflect.Array && t1.Len() != t2.Len() {
			return false
		}
		return CanCast(t1.Elem(), t2.Elem(), true)
	case reflect.Struct:
		if t1.NumField() != t2.NumField() {
			return false
		}
		for i := range t1.NumField() {
			if !CanCast(t1.Field(i).Type, t2.Field(i).Type, true) {
				return false
			}
		}
	case reflect.Interface, reflect.UnsafePointer:
		return t1 == t2
	}
	return true
}

func InitStruct(v reflect.Value) {
	v = InitPtr(v)
	switch v.Kind() {
	case reflect.Struct:
		for i := range v.NumField() {
			field := v.Field(i)
			InitStruct(field)
		}
	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			InitStruct(v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			InitStruct(v.MapIndex(key))
		}
	case reflect.Interface:
		v = v.Elem()
		if v.IsValid() {
			InitStruct(v)
		}
	}
}

func GetSubFieldByType[T any](v any) *T {
	value := reflect.ValueOf(v).Elem()
	for i := range value.NumField() {
		if dao, ok := value.Field(i).Interface().(T); ok {
			return &dao
		}
	}
	return new(T)
}
