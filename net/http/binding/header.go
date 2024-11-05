package binding

import (
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/reflect/mtos"
	"net/http"
	"net/textproto"
	"reflect"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req *http.Request, obj interface{}) error {

	if err := mtos.Decode(obj, req.Header); err != nil {
		return err
	}

	return Validate(obj)
}

func MapHeader(ptr interface{}, h map[string][]string) error {
	return encoding.MapFormByTag(ptr, HeaderSource(h), "header")
}

type HeaderSource map[string][]string

var _ encoding.Setter = HeaderSource(nil)

func (hs HeaderSource) Peek(key string) ([]string, bool) {
	v, ok := hs[textproto.CanonicalMIMEHeaderKey(key)]
	return v, ok
}

func (hs HeaderSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt encoding.SetOptions) (isSet bool, err error) {
	return encoding.SetByKVs(value, field, hs, tagValue, opt)
}
