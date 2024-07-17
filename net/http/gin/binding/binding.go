package binding

import (
	"fmt"
	"github.com/hopeio/utils/net/http/binding"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
)

func SetTag(tag string) {
	if tag != "" {
		binding.Tag = tag
	}
}

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*gin.Context, interface{}) error
}

// BindingBody adds BindBody method to Binding. BindBody is similar with GinBind,
// but it reads the body from supplied bytes instead of req.Body.
type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
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

// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
func Default(method string, contentType string) Binding {
	if method == http.MethodGet {
		return Query
	}

	return Body(contentType)
}

func Body(contentType string) Binding {
	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEPOSTForm:
		return FormPost
	case MIMEXML, MIMEXML2:
		return XML
	case MIMEPROTOBUF:
		return ProtoBuf
	case MIMEMSGPACK, MIMEMSGPACK2:
		return MsgPack
	case MIMEYAML:
		return YAML
	case MIMEMultipartPOSTForm:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return JSON
	}
}

func Validate(obj interface{}) error {
	return binding.Validator.ValidateStruct(obj)
}

func Bind(c *gin.Context, obj interface{}) error {
	tag := binding.Tag
	if c.Request.Body != nil && c.Request.ContentLength != 0 {
		b := Body(c.ContentType())
		err := b.Bind(c, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = b.Name()
	}

	var args binding.ArgSource
	if len(c.Params) > 0 {
		args = append(args, uriSource(c.Params))
	}
	if len(c.Request.URL.RawQuery) > 0 {
		args = append(args, binding.FormSource(c.Request.URL.Query()))
	}
	if len(c.Request.Header) > 0 {
		args = append(args, binding.HeaderSource(c.Request.Header))
	}
	err := binding.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}
