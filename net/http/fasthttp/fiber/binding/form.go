package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	binding2 "github.com/hopeio/utils/net/http/fasthttp/binding"
	"github.com/hopeio/utils/reflect/mtos"
)

type formPostBinding struct{}
type formMultipartBinding struct{}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*binding2.MultipartRequest)(ctx.Request()), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	if err := mtos.MapFormByTag(obj, (*binding2.ArgsSource)(ctx.Request().PostArgs()), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
