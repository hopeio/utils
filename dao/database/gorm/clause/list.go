package clause

import (
	"gorm.io/gorm/clause"
	"reflect"
)

func EqualByStruct(param any) clause.Expression {
	v := reflect.ValueOf(param)
	v = reflect.Indirect(v)
	var conds []clause.Expression
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldKind := field.Kind()
		switch fieldKind {
		case reflect.Interface, reflect.Pointer:
			if field.IsNil() {
				continue
			}
			field = field.Elem()
		case reflect.Struct, reflect.Map:
			continue
		default:
			conds = append(conds, clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: t.Field(i).Name}, Value: v})
		}

	}
	return nil
}

func NotEqualByStruct() {

}

func RangeByStruct() {

}
