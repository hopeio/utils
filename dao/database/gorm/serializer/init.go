package serializer

import "gorm.io/gorm/schema"

func init() {
	schema.RegisterSerializer("json", JSONSerializer{})

	schema.RegisterSerializer("string_array", StringArraySerializer{})

}
