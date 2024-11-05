package binding

import (
	"github.com/hopeio/utils/encoding"
	"net/http"
	"reflect"
)

// support go 1.22
type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(req *http.Request, obj interface{}) error {
	if err := encoding.MapFormByTag(obj, (*UriSource)(req), "uri"); err != nil {
		return err
	}
	return Validate(obj)

}

type UriSource http.Request

var _ encoding.Setter = (*UriSource)(nil)

func (req *UriSource) Peek(key string) ([]string, bool) {
	v := (*http.Request)(req).PathValue(key)
	return []string{v}, v != ""
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (req *UriSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt encoding.SetOptions) (isSet bool, err error) {
	return encoding.SetByKVs(value, field, req, tagValue, opt)
}
