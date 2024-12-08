package apidoc

import (
	reflecti "github.com/hopeio/utils/reflect"
	basehigh "github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"reflect"
	"strings"
)

func DefinitionsApi(schema *basehigh.Schema, v any) {
	schema.Type = []string{"object"}
	schema.Properties = orderedmap.New[string, *basehigh.SchemaProxy]()

	body := reflect.TypeOf(v).Elem()
	var typ, subFieldName string
	for i := range body.NumField() {
		json := strings.Split(body.Field(i).Tag.Get("json"), ",")[0]
		if json == "" || json == "-" {
			continue
		}
		fieldType := body.Field(i).Type
		switch fieldType.Kind() {
		case reflect.Struct:
			typ = "object"
			v = reflect.ValueOf(v).Elem().Field(i).Addr().Interface()
			subFieldName = fieldType.Name()
		case reflect.Ptr:
			typ = "object"
			v = reflect.New(fieldType.Elem()).Interface()
			subFieldName = fieldType.Elem().Name()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			typ = "integer"
		case reflect.Array, reflect.Slice:
			typ = "array"
			v = reflect.New(reflecti.DerefType(fieldType)).Interface()
			subFieldName = reflecti.DerefType(fieldType).Name()
		case reflect.Float32, reflect.Float64:
			typ = "number"
		case reflect.String:
			typ = "string"
		case reflect.Bool:
			typ = "boolean"

		}
		var subSchemaProxy *basehigh.SchemaProxy
		if typ == "object" {
			subSchemaProxy = basehigh.CreateSchemaProxyRef(subFieldName)
			subSchema := subSchemaProxy.Schema()
			subSchema.Type = []string{typ}
			DefinitionsApi(subSchema, v)
		}
		if typ == "array" {
			subSchemaProxy = basehigh.CreateSchemaProxyRef(subFieldName)
			subSchema := subSchemaProxy.Schema()
			arrSubSchemaProxy := basehigh.CreateSchemaProxyRef(subFieldName)
			subSchema.Items = &basehigh.DynamicValue[*basehigh.SchemaProxy, bool]{A: arrSubSchemaProxy}
			DefinitionsApi(arrSubSchemaProxy.Schema(), v)
		}
		schema.Properties.Set(json, subSchemaProxy)
	}
}
