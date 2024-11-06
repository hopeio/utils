package binding

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/hopeio/utils/reflect/mtos"
	stringsi "github.com/hopeio/utils/strings"
	"net/http"
	"strings"
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
	CustomBody    = bodyBinding{name: "json", unmarshaller: json.Unmarshal}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	Uri           = uriBinding{}
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
	case binding.MIMEPOSTForm:
		return FormPost
	case binding.MIMEMultipartPOSTForm:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return CustomBody
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

	var args mtos.PeekVsSource

	args = append(args, (*uriSource)(c.(*fiber.DefaultCtx)))

	if query := c.Queries(); len(query) > 0 {
		args = append(args, QuerySource(query))
	}
	if headers := c.GetReqHeaders(); len(headers) > 0 {
		args = append(args, binding.HeaderSource(headers))
	}
	err := mtos.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}

func RegisterBodyBinding(name string, unmarshaller func(data []byte, obj any) error) {
	binding.SetTag(name)
	CustomBody.name = name
	CustomBody.unmarshaller = unmarshaller
}

func SetTag(tag string) {
	binding.SetTag(tag)
}

type QuerySource map[string]string

func (q QuerySource) Peek(key string) ([]string, bool) {
	v, ok := q[key]
	if strings.Contains(v, ",") {
		return strings.Split(v, ","), true
	}
	return []string{v}, ok
}
