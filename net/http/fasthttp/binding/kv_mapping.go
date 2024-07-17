package binding

import (
	"github.com/hopeio/utils/net/http/binding"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/valyala/fasthttp"
	"reflect"
)

type ArgsSource fasthttp.Args

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *ArgsSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSet bool, err error) {
	return binding.SetByKV(value, field, form, tagValue, opt)
}

func (form *ArgsSource) Peek(key string) ([]string, bool) {
	v := stringsi.BytesToString((*fasthttp.Args)(form).Peek(key))
	return []string{v}, v != ""
}

type CtxSource fasthttp.RequestCtx

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *CtxSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSet bool, err error) {
	return binding.SetByKV(value, field, form, tagValue, opt)
}

func (form *CtxSource) Peek(key string) ([]string, bool) {
	v := (*fasthttp.RequestCtx)(form).UserValue(key).(string)
	return []string{v}, v != ""
}

type HeaderSource fasthttp.RequestHeader

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *HeaderSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSet bool, err error) {
	return binding.SetByKV(value, field, form, tagValue, opt)
}

func (form *HeaderSource) Peek(key string) ([]string, bool) {
	v := stringsi.BytesToString((*fasthttp.RequestHeader)(form).Peek(key))
	return []string{v}, v != ""
}
