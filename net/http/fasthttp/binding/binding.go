package binding

import (
	"fmt"
	"github.com/hopeio/utils/net/http/binding"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/valyala/fasthttp"
	"net/http"
)

type Binding interface {
	Name() string
	Bind(*fasthttp.RequestCtx, interface{}) error
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
	ProtoBuf      = protobufBinding{}
	MsgPack       = msgpackBinding{}
	YAML          = yamlBinding{}
	Header        = headerBinding{}
)

func Default(method, contentType []byte) Binding {
	if stringsi.BytesToString(method) == http.MethodGet {
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
	default:
		return JSON
	}
}

func Bind(c *fasthttp.RequestCtx, obj interface{}) error {
	tag := binding.Tag
	if data := c.Request.Body(); len(data) > 0 {
		b := Body(c.Request.Header.ContentType())
		err := b.Bind(c, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = binding.Tag
	}

	var args binding.ArgSource

	if query := c.QueryArgs(); query != nil {
		args = append(args, (*ArgsSource)(query))
	}
	args = append(args, (*HeaderSource)(&c.Request.Header))
	err := binding.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}
