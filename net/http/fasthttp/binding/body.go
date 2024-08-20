package binding

import (
	"github.com/hopeio/utils/encoding"
	"github.com/valyala/fasthttp"
	"io"
)

type bodyBinding struct {
	name         string
	unmarshaller func([]byte, any) error
	newDecoder   func(io.Reader) encoding.Decoder
}

func (b bodyBinding) Name() string {
	return b.name
}

func (b bodyBinding) Bind(ctx *fasthttp.RequestCtx, obj interface{}) error {
	return b.unmarshaller(ctx.Request.Body(), obj)
}
