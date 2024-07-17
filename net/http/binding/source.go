package binding

import (
	"fmt"
	"reflect"
)

type Arg interface {
	Peek(key string) ([]string, bool)
}

type ArgSource []Arg

func (args ArgSource) Peek(key string) (v []string, ok bool) {
	for i := range args {
		if v, ok = args[i].Peek(key); ok {
			return
		}
	}
	return
}

func (args ArgSource) TrySet(value reflect.Value, field reflect.StructField, key string, opt SetOptions) (isSet bool, err error) {
	return SetByKV(value, field, args, key, opt)
}

func SetByKV(value reflect.Value, field reflect.StructField, kv Arg, tagValue string, opt SetOptions) (isSet bool, err error) {
	vs, ok := kv.Peek(tagValue)
	if !ok && !opt.isDefaultExists {
		return false, nil
	}

	switch value.Kind() {
	case reflect.Slice:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		return true, setSlice(vs, value, field)
	case reflect.Array:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		if len(vs) != value.Len() {
			return false, fmt.Errorf("%q is not valid value for %s", vs, value.Type().String())
		}
		return true, setArray(vs, value, field)
	default:
		var val string
		if !ok {
			val = opt.defaultValue
		}

		if len(vs) > 0 {
			val = vs[0]
		}
		return true, setWithProperType(val, value, field)
	}
}

type KVSource map[string]string

func (form KVSource) Peek(key string) ([]string, bool) {
	v, ok := form[key]
	return []string{v}, ok
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form KVSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt SetOptions) (isSet bool, err error) {
	return SetByKV(value, field, form, tagValue, opt)
}

type FormSource map[string][]string

var _ Setter = FormSource(nil)

func (form FormSource) Peek(key string) ([]string, bool) {
	v, ok := form[key]
	return v, ok
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form FormSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt SetOptions) (isSet bool, err error) {
	return SetByKV(value, field, form, tagValue, opt)
}
