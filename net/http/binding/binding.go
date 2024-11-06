package binding

import (
	"encoding/json"
	"fmt"
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/reflect/mtos"
	"io"
	"net/http"
	"reflect"

	"github.com/hopeio/utils/validation/validator"
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

var Tag = "json"

func SetTag(tag string) {
	if tag != "" {
		Tag = tag
	}
	mtos.SetAliasTag(tag)
}

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}

// BindingBody adds BindBody method to Binding. BindBody is similar with GinBind,
// but it reads the body from supplied bytes instead of req.Body.
type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v8.18.2
// under the hood.
var Validator = validator.DefaultValidator

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	Uri    = uriBinding{}
	Query  = queryBinding{}
	Header = headerBinding{}

	CustomBody    = bodyBinding{name: "json", unmarshaller: json.Unmarshal}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
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
	case MIMEPOSTForm:
		return FormPost
	case MIMEMultipartPOSTForm:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return CustomBody
	}
}

func Validate(obj interface{}) error {
	return Validator.ValidateStruct(obj)
}

func Bind(r *http.Request, obj interface{}) error {
	tag := Tag
	if r.Body != nil && r.ContentLength != 0 {
		b := Body(r.Header.Get("Content-Type"))
		err := b.Bind(r, obj)
		if err != nil {
			return fmt.Errorf("body bind error: %w", err)
		}
		tag = b.Name()
	}

	var args mtos.PeekVsSource
	if !reflect.ValueOf(r).Elem().FieldByName("pat").IsNil() {
		args = append(args, (*UriSource)(r))
	}
	if len(r.URL.RawQuery) > 0 {
		args = append(args, mtos.KVsSource(r.URL.Query()))
	}
	if len(r.Header) > 0 {
		args = append(args, HeaderSource(r.Header))
	}
	err := mtos.MapFormByTag(obj, args, tag)
	if err != nil {
		return fmt.Errorf("args bind error: %w", err)
	}
	return nil
}

func RegisterBodyBinding(name string, unmarshaller func(data []byte, obj any) error) {
	CustomBody.name = name
	CustomBody.unmarshaller = unmarshaller
}

func RegisterBodyBindingByDecoder(name string, newDecoder func(io.Reader) encoding.Decoder) {
	CustomBody.name = name
	CustomBody.decoder = newDecoder
}
