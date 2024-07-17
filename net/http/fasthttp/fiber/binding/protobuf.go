package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (b protobufBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	return binding.DecodeProtobuf(ctx.Body(), obj)
}
