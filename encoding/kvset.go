package encoding

import (
	"fmt"
	"github.com/hopeio/utils/reflect/converter"
	"reflect"
)

type PeekV interface {
	Peek(key string) (string, bool)
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
func (args Args) TrySet(value reflect.Value, field *reflect.StructField, key string) (isSet bool, err error) {
	return SetByKV(value, field, args, key)
}

func SetByKV(value reflect.Value, field *reflect.StructField, kv PeekV, tagValue string) (isSet bool, err error) {
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

type KVSource map[string]string

func (form KVSource) Peek(key string) (string, bool) {
	v, ok := form[key]
	return v, ok
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form KVSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string) (isSet bool, err error) {
	return SetByKV(value, field, form, tagValue)
}

type KVsSource map[string][]string

var _ Setter = KVsSource(nil)

func (form KVsSource) Peek(key string) ([]string, bool) {
	v, ok := form[key]
	return v, ok
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form KVsSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt SetOptions) (isSet bool, err error) {
	return SetByKVs(value, field, form, tagValue, opt)
}

type PeekVs interface {
	Peek(key string) ([]string, bool)
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

func (args Args2) TrySet(value reflect.Value, field *reflect.StructField, key string, opt SetOptions) (isSet bool, err error) {
	return SetByKVs(value, field, args, key, opt)
}

type PeekVsSource []PeekVs

func (args PeekVsSource) Peek(key string) (v []string, ok bool) {
	for i := range args {
		if v, ok = args[i].Peek(key); ok {
			return
		}
	}
	return
}

func (args PeekVsSource) TrySet(value reflect.Value, field *reflect.StructField, key string, opt SetOptions) (isSet bool, err error) {
	return SetByKVs(value, field, args, key, opt)
}

func SetByKVs(value reflect.Value, field *reflect.StructField, kv PeekVs, tagValue string, opt SetOptions) (isSet bool, err error) {
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
