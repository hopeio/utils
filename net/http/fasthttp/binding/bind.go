package binding

import (
	httpi "github.com/hopeio/utils/net/http"
	"github.com/valyala/fasthttp"
)

func NewReq[REQ any](c *fasthttp.RequestCtx) (*REQ, error) {
	req := new(REQ)
	err := Bind(c, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
func BindJSON(r *fasthttp.RequestCtx, obj interface{}) error {
	return MustBindWith(r, obj, JSON)
}

// BindXML is a shortcut for c.MustBindWith(obj, binding.BindXML).
func BindXML(r *fasthttp.RequestCtx, obj interface{}) error {
	return MustBindWith(r, obj, XML)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func BindQuery(r *fasthttp.RequestCtx, obj interface{}) error {
	return MustBindWith(r, obj, Query)
}

// BindYAML is a shortcut for c.MustBindWith(obj, binding.YAML).
func BindYAML(r *fasthttp.RequestCtx, obj interface{}) error {
	return MustBindWith(r, obj, YAML)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func MustBindWith(r *fasthttp.RequestCtx, obj interface{}, b Binding) error {
	return ShouldBindWith(r, obj, b)
}

// ShouldBind checks the Content-Type to select a binding engine automatically,
// Depending the "Content-Type" header different bindings are used:
//
//	"application/json" --> JSON binding
//	"application/xml"  --> XML binding
//
// otherwise --> returns an error
// It parses the request's body as JSON if Content-Type == "application/json" using JSON or XML as a JSON input.
// It decodes the json payload into the struct specified as a pointer.
// Like c.GinBind() but this method does not set the response status code to 400 and abort if the json is not valid.
func ShouldBind(r *fasthttp.RequestCtx, obj interface{}) error {
	b := Default(r.Method(), r.Request.Header.Peek(httpi.HeaderContentType))
	return ShouldBindWith(r, obj, b)
}

// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
func ShouldBindJSON(r *fasthttp.RequestCtx, obj interface{}) error {
	return ShouldBindWith(r, obj, JSON)
}

// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML).
func ShouldBindXML(r *fasthttp.RequestCtx, obj interface{}) error {
	return ShouldBindWith(r, obj, XML)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func ShouldBindQuery(r *fasthttp.RequestCtx, obj interface{}) error {
	return ShouldBindWith(r, obj, Query)
}

// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML).
func ShouldBindYAML(r *fasthttp.RequestCtx, obj interface{}) error {
	return ShouldBindWith(r, obj, YAML)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(r *fasthttp.RequestCtx, obj interface{}, b Binding) error {
	return b.Bind(r, obj)
}
