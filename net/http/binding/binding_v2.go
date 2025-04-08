package binding

import (
	"github.com/hopeio/utils/reflect/mtos"
	"reflect"
	"sync"
)

var defaultTags = []string{"uri", "path", "query", "header", "form"}

type Peek interface {
	Peek(key string) string
}

type Source interface {
	Uri() mtos.Setter
	Query() mtos.Setter
	Header() mtos.Setter
	BodyBind(obj any) error
}

type Field struct {
	Tag      string
	TagValue string
	Index    int
}

var cache = sync.Map{}

func BindV2(s Source, obj any) error {
	value := reflect.ValueOf(obj)
	typ := value.Type()
	err := s.BodyBind(obj)
	if err != nil {
		return err
	}
	if fields, ok := cache.Load(typ); ok {
		for _, field := range fields.([]Field) {
			var setter mtos.Setter
			switch field.Tag {
			case "uri", "path":
				setter = s.Uri()
			case "query":
				setter = s.Query()
			case "header":
				setter = s.Header()
			}
			_, err = setter.TrySet(value.Field(field.Index), nil, field.TagValue, mtos.SetOptions{})
			if err != nil {
				return err
			}
		}
		return nil
	}
	var fields []Field
	for i := 0; i < value.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath != "" && !sf.Anonymous { // unexported
			continue
		}
		var tagValue string
		var tag string
		for _, tag = range defaultTags {
			tagValue = sf.Tag.Get(tag)
			if tagValue != "" && tagValue != "-" {
				break
			}
		}
		if tagValue == "-" { // just ignoring this field
			continue
		}

		var setter mtos.Setter
		switch tag {
		case "uri", "path":
			setter = s.Uri()
		case "query":
			setter = s.Query()
		case "header":
			setter = s.Header()
		}
		_, err = setter.TrySet(value.Field(i), &sf, tagValue, mtos.SetOptions{})
		if err != nil {
			return err
		}
		fields = append(fields, Field{
			Tag:      tag,
			TagValue: tagValue,
			Index:    i,
		})
	}
	cache.Store(typ, fields)
	return nil
}
