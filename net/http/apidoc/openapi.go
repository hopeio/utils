package apidoc

import (
	"github.com/getkin/kin-openapi/openapi3"
	reflecti "github.com/hopeio/utils/reflect"
	"reflect"
	"strings"
)

func DefinitionsApi(schema *openapi3.Schema, v interface{}) {
	typs := openapi3.Types{"object"}
	schema.Type = &typs
	schema.Properties = make(map[string]*openapi3.SchemaRef)

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
		typs := openapi3.Types{typ}
		subSchema := openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: &typs,
			},
		}
		if typ == "object" {
			subSchema.Ref = "#/definitions/" + subFieldName
			DefinitionsApi(subSchema.Value, v)
		}
		if typ == "array" {
			subSchema.Value.Items = new(openapi3.SchemaRef)
			subSchema.Value.Items.Value = &openapi3.Schema{}
			subSchema.Value.Items.Ref = "#/definitions/" + subFieldName
			DefinitionsApi(subSchema.Value, v)
		}
		schema.Properties[json] = &subSchema
	}
}

func genSchema(v interface{}) *openapi3.Schema {
	return nil
}
