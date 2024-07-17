package binding

import (
	http2 "github.com/hopeio/utils/net/http"
	"net/http"
)

func NewReq[REQ any](r *http.Request) (*REQ, error) {
	req := new(REQ)
	err := Bind(r, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
func BindJSON(r *http.Request, obj interface{}) error {
	return MustBindWith(r, obj, JSON)
}

// BindXML is a shortcut for c.MustBindWith(obj, binding.BindXML).
func BindXML(r *http.Request, obj interface{}) error {
	return MustBindWith(r, obj, XML)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func BindQuery(r *http.Request, obj interface{}) error {
	return MustBindWith(r, obj, Query)
}

// BindYAML is a shortcut for c.MustBindWith(obj, binding.YAML).
func BindYAML(r *http.Request, obj interface{}) error {
	return MustBindWith(r, obj, YAML)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// BindUri binds the passed struct pointer using binding.Uri.
// It will abort the request with HTTP 400 if any error occurs.
func BindUri(r *http.Request, obj interface{}) error {
	return ShouldBindUri(r, obj)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func MustBindWith(r *http.Request, obj interface{}, b Binding) error {
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
func ShouldBind(r *http.Request, obj interface{}) error {
	b := Default(r.Method, r.Header.Get(http2.HeaderContentType))
	return ShouldBindWith(r, obj, b)
}

// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
func ShouldBindJSON(r *http.Request, obj interface{}) error {
	return ShouldBindWith(r, obj, JSON)
}

// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML).
func ShouldBindXML(r *http.Request, obj interface{}) error {
	return ShouldBindWith(r, obj, XML)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func ShouldBindQuery(r *http.Request, obj interface{}) error {
	return ShouldBindWith(r, obj, Query)
}

// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML).
func ShouldBindYAML(r *http.Request, obj interface{}) error {
	return ShouldBindWith(r, obj, YAML)
}

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func ShouldBindUri(r *http.Request, obj interface{}) error {
	return Uri.Bind(r, obj)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(r *http.Request, obj interface{}, b Binding) error {
	return b.Bind(r, obj)
}
