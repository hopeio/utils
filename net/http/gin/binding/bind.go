package binding

import (
	"github.com/gin-gonic/gin"
	"io"
)

func NewReq[REQ any](c *gin.Context) (*REQ, error) {
	req := new(REQ)
	err := Bind(c, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
func BindJSON(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, JSON)
}

// BindXML is a shortcut for c.MustBindWith(obj, binding.BindXML).
func BindXML(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, XML)
}

// BindQuery is a shortcut for c.MustBindWith(obj, binding.Query).
func BindQuery(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, Query)
}

// BindYAML is a shortcut for c.MustBindWith(obj, binding.YAML).
func BindYAML(c *gin.Context, obj interface{}) error {
	return MustBindWith(c, obj, YAML)
}

// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func MustBindWith(c *gin.Context, obj interface{}, b Binding) error {
	if err := ShouldBindWith(c, obj, b); err != nil {
		return err
	}
	if err := Validate(obj); err != nil {
		return err
	}
	return nil
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
func ShouldBind(c *gin.Context, obj interface{}) error {
	b := Default(c.Request.Method, c.ContentType())
	return b.Bind(c, obj)
}

// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, JSON)
}

// ShouldBindXML is a shortcut for c.ShouldBindWith(obj, binding.XML).
func ShouldBindXML(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, XML)
}

// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func ShouldBindQuery(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, Query)
}

// ShouldBindYAML is a shortcut for c.ShouldBindWith(obj, binding.YAML).
func ShouldBindYAML(c *gin.Context, obj interface{}) error {
	return ShouldBindWith(c, obj, YAML)
}

// ShouldBindUri binds the passed struct pointer using the specified binding engine.
func ShouldBindUri(r *gin.Context, obj interface{}) error {
	return Uri.Bind(r, obj)
}

// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func ShouldBindWith(c *gin.Context, obj interface{}, b Binding) error {
	return b.Bind(c, obj)
}

// ShouldBindBodyWith is similar with ShouldBindWith, but it stores the request
// body into the context, and reuse when it is called again.
//
// NOTE: This method reads the body before binding. So you should use
// ShouldBindWith for better performance if you need to call only once.
func ShouldBindBodyWith(c *gin.Context, obj interface{}, bb BindingBody) (err error) {
	var body []byte
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, err = io.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		c.Set(gin.BodyBytesKey, body)
	}
	return bb.BindBody(body, obj)
}
