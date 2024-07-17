//go:build !nomsgpack
// +build !nomsgpack

package binding

import (
	"bytes"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (m msgpackBinding) Bind(ctx fiber.Ctx, obj interface{}) error {
	return m.BindBody(ctx.Body(), obj)
}

func (msgpackBinding) BindBody(body []byte, obj interface{}) error {
	return binding.DecodeMsgPack(bytes.NewReader(body), obj)
}
