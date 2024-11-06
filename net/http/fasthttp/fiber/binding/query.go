package binding

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	fbinding "github.com/hopeio/utils/net/http/fasthttp/binding"
	"github.com/hopeio/utils/reflect/mtos"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	values := ctx.Request().URI().QueryArgs()
	if err := mtos.MapFormByTag(obj, (*fbinding.ArgsSource)(values), binding.Tag); err != nil {
		return err
	}
	return binding.Validate(obj)
}
