package binding

import (
	"github.com/hopeio/utils/encoding"
	"github.com/hopeio/utils/net/http/binding"
	"github.com/valyala/fasthttp"
	"reflect"
)

type MultipartRequest fasthttp.Request

var _ encoding.Setter = (*MultipartRequest)(nil)

// TrySet tries to set a value by the multipart request with the binding a form file
func (r *MultipartRequest) TrySet(value reflect.Value, field *reflect.StructField, key string, opt encoding.SetOptions) (isSet bool, err error) {
	req := (*fasthttp.Request)(r)
	form, err := req.MultipartForm()
	if err != nil {
		return false, err
	}
	if files := form.File[key]; len(files) != 0 {
		return binding.SetByMultipartFormFile(value, field, files)
	}

	return encoding.SetByKVs(value, field, encoding.KVsSource(form.Value), key, opt)
}
