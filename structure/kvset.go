package structure

import (
	"github.com/hopeio/utils/reflect/converter"
	"reflect"
)

type PeekV interface {
	Peek(key string) (string, bool)
}
type PeekVs interface {
	Peek(key string) ([]string, bool)
}

type Setter interface {
	TrySet(value reflect.Value, field reflect.StructField, key string) (isSet bool, err error)
}

type Args []PeekV

func (args Args) Peek(key string) (v string, ok bool) {
	for i := range args {
		if v, ok = args[i].Peek(key); ok {
			return
		}
	}
	return
}
func (args Args) TrySet(value reflect.Value, field reflect.StructField, key string) (isSet bool, err error) {
	return SetByKV(value, field, args, key)
}

func SetByKV(value reflect.Value, field reflect.StructField, kv PeekV, tagValue string) (isSet bool, err error) {
	vs, ok := kv.Peek(tagValue)
	if !ok {
		return false, nil
	}
	err = converter.SetValueByString(value, vs)
	if err != nil {
		return false, err
	}
	return true, nil
}

type Args2 []PeekVs

func (args Args2) Peek(key string) (v []string, ok bool) {
	for i := range args {
		if v, ok = args[i].Peek(key); ok {
			return
		}
	}
	return
}

func (args Args2) TrySet(value reflect.Value, field reflect.StructField, key string) (isSet bool, err error) {
	return SetByKV2(value, field, args, key)
}

func SetByKV2(value reflect.Value, field reflect.StructField, kv PeekVs, tagValue string) (isSet bool, err error) {
	vs, ok := kv.Peek(tagValue)
	if !ok {
		return false, nil
	}
	// TODO
	/*	err = converter.SetValueByStrings(value, vs)
		if err != nil {
			return false, err
		}*/
	err = converter.SetValueByString(value, vs[0])
	if err != nil {
		return false, err
	}
	return true, nil
}

type KVSource map[string]string

func (form KVSource) Peek(key string) ([]string, bool) {
	v, ok := form[key]
	return []string{v}, ok
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form KVSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string) (isSet bool, err error) {
	return SetByKV2(value, field, form, tagValue)
}

type KVsSource map[string][]string

var _ Setter = KVsSource(nil)

func (form KVsSource) Peek(key string) ([]string, bool) {
	v, ok := form[key]
	return v, ok
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form KVsSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string) (isSet bool, err error) {
	return SetByKV2(value, field, form, tagValue)
}
