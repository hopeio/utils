package serializer

import (
	"context"
	"github.com/hopeio/utils/dao/database/datatypes"
	reflecti "github.com/hopeio/utils/reflect"
	"gorm.io/gorm/schema"
	"reflect"
	"unsafe"
)

// StringArraySerializer array序列化器
type StringArraySerializer struct {
}

// 实现 Scan 方法
func (StringArraySerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value,
	dbValue any) (err error) {
	if dbValue != nil {
		var arr datatypes.StringArray
		err = arr.Scan(dbValue)
		if err != nil {
			return err
		}
		fieldValue := reflect.ValueOf(arr)
		field.ReflectValueOf(ctx, dst).Set(fieldValue)
	}
	return
}

// 实现 Value 方法
func (StringArraySerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue any) (any, error) {
	arr := (*datatypes.StringArray)(unsafe.Pointer((*reflecti.Eface)(unsafe.Pointer(&fieldValue)).Value))
	return (*arr).Value()
}
