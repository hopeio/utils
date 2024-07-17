package binding

import (
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/net/http/binding"
	"reflect"
)

// support go 1.22
type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if err := binding.MappingByPtr(obj, (uriSource)(ctx.Params), binding.Tag); err != nil {
		return err
	}
	return Validate(obj)
}

type uriSource gin.Params

var _ binding.Setter = uriSource(nil)

func (param uriSource) Peek(key string) ([]string, bool) {
	for i := range param {
		if param[i].Key == key {
			return []string{param[i].Value}, true
		}
	}
	return nil, false
}

// TrySet tries to set a value by request's form source (like map[string][]string)
func (param uriSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt binding.SetOptions) (isSet bool, err error) {
	return binding.SetByKV(value, field, param, tagValue, opt)
}
