package binding

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/net/http/binding"

	"io"
)

type bodyBinding struct {
	name         string
	unmarshaller func([]byte, any) error
	decoder      func(io.Reader) encoding.Decoder
}

func (b bodyBinding) Name() string {
	return b.name
}

func (b bodyBinding) Bind(ctx *gin.Context, obj interface{}) error {
	if ctx == nil || ctx.Request.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return binding.CustomBody.Bind(ctx.Request, obj)
}
