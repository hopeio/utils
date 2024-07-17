package binding

import (
	"net/http"
	"net/textproto"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {

	if err := Decode(obj, req.Header); err != nil {
		return err
	}

	return Validate(obj)
}

func MapHeader(ptr interface{}, h map[string][]string) error {
	return MappingByPtr(ptr, HeaderSource(h), "header")
}

type HeaderSource map[string][]string

var _ Setter = HeaderSource(nil)

func (hs HeaderSource) Peek(key string) ([]string, bool) {
	v, ok := hs[textproto.CanonicalMIMEHeaderKey(key)]
	return v, ok
}

func (hs HeaderSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt SetOptions) (isSet bool, err error) {
	return SetByKV(value, field, hs, tagValue, opt)
}
