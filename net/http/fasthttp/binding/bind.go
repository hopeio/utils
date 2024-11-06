package binding

import (
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

func BindBody(r *fasthttp.RequestCtx, obj interface{}) error {
	return MustBindWith(r, obj, CustomBody)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func BindQuery(r *fasthttp.RequestCtx, obj interface{}) error {
	return MustBindWith(r, obj, Query)
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
	b := Default(r.Method(), r.Request.Header.Peek("Content-Type"))
	return ShouldBindWith(r, obj, b)
}

func ShouldBindBody(r *fasthttp.RequestCtx, obj interface{}) error {
	return ShouldBindWith(r, obj, CustomBody)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func ShouldBindQuery(r *fasthttp.RequestCtx, obj interface{}) error {
	return ShouldBindWith(r, obj, Query)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(r *fasthttp.RequestCtx, obj interface{}, b Binding) error {
	return b.Bind(r, obj)
}
