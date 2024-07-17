package binding

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (j jsonBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	body := ctx.Request().Body()
	if body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.DecodeJsonData(body, obj)
}
