package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	fbinding "github.com/hopeio/utils/net/http/fasthttp/binding"
)

type headerBinding struct{}

func (headerBinding) Name() string {
	return "header"
}

func (headerBinding) Bind(req fiber.Ctx, obj interface{}) error {

	if err := binding.MappingByPtr(obj, (*fbinding.HeaderSource)(&req.Request().Header), binding.Tag); err != nil {
		return err
	}

	return binding.Validate(obj)
}
