package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (x xmlBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	return binding.DecodeXmlData(ctx.Body(), obj)
}
