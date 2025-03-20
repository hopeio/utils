package v2

import (
	"github.com/hopeio/utils/reflect/mtos"
	"gorm.io/gen/field"
	"reflect"
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

func BindV2(s Source, obj any) error {
	value := reflect.ValueOf(obj)
	typ := value.Type()
	err := s.BodyBind(obj)
	if err != nil {
		return err
	}
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
		_, err := mapping(value.Field(i), &sf, setter)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapping(value reflect.Value, field *reflect.StructField, setter mtos.Setter) (bool, error) {

	var vKind = value.Kind()
	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		if value.IsNil() {
			isNew = true
			vPtr = reflect.New(value.Type().Elem())
		}
		isSet, err := mapping(vPtr.Elem(), field, setter)
		if err != nil {
			return false, err
		}
		if isNew && isSet {
			value.Set(vPtr)
		}
		return isSet, nil
	}

	if vKind == reflect.Struct {
		tValue := value.Type()

		var isSet bool
		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}
			ok, err := mapping(value.Field(i), &sf, setter)
			if err != nil {
				return false, err
			}
			isSet = isSet || ok
		}
		return isSet, nil
	}

	if field != nil && !field.Anonymous {
		ok, err := tryToSetValue(value, field, setter, tagValue)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}

	return false, nil
}
