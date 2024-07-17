package binding

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	stringsi "github.com/hopeio/utils/strings"
	"net/http"
)

type Binding interface {
	Name() string

	Bind(fiber.Ctx, interface{}) error
}

type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

var (
	JSON          = jsonBinding{}
	XML           = xmlBinding{}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	Uri           = uriBinding{}
	ProtoBuf      = protobufBinding{}
	MsgPack       = msgpackBinding{}
	YAML          = yamlBinding{}
	Header        = headerBinding{}
)

func Default(method string, contentType []byte) Binding {
	if method == http.MethodGet {
		return Query
	}

	return Body(contentType)
}

func Body(contentType []byte) Binding {
	switch stringsi.BytesToString(contentType) {
	case binding.MIMEJSON:
		return JSON
	case binding.MIMEPOSTForm:
		return FormPost
	case binding.MIMEXML, binding.MIMEXML2:
		return XML
	case binding.MIMEPROTOBUF:
		return ProtoBuf
	case binding.MIMEMSGPACK, binding.MIMEMSGPACK2:
		return MsgPack
	case binding.MIMEYAML:
		return YAML
	case binding.MIMEMultipartPOSTForm:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return JSON
	}
}

func Bind(c fiber.Ctx, obj interface{}) error {
	tag := binding.Tag
	if data := c.Body(); len(data) > 0 {
		b := Body(c.Request().Header.ContentType())
		err := b.Bind(c, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = b.Name()
	}

	var args binding.ArgSource

	args = append(args, (*uriSource)(c.(*fiber.DefaultCtx)))

	if query := c.Queries(); len(query) > 0 {
		args = append(args, binding.KVSource(query))
	}
	if headers := c.GetReqHeaders(); len(headers) > 0 {
		args = append(args, binding.HeaderSource(headers))
	}
	err := binding.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}
