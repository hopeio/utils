package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/net/http/binding"
	"reflect"
)

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(c fiber.Ctx, obj interface{}) error {
	if err := encoding.MapFormByTag(obj, (*uriSource)(c.(*fiber.DefaultCtx)), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}

type uriSource fiber.DefaultCtx

// TrySet tries to set a value by request's form source (like map[string][]string)
func (form *uriSource) TrySet(value reflect.Value, field *reflect.StructField, tagValue string, opt encoding.SetOptions) (isSet bool, err error) {
	return encoding.SetByKVs(value, field, form, tagValue, opt)
}

func (form *uriSource) Peek(key string) ([]string, bool) {
	v := (*fiber.DefaultCtx)(form).Params(key)
	return []string{v}, v != ""
}
