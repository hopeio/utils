package clause

import (
	sql2 "github.com/hopeio/utils/dao/database/sql"
	"gorm.io/gorm/clause"
	"reflect"
)

func EqualByStruct(param any) clause.Expression {
	v := reflect.ValueOf(param)
	v = reflect.Indirect(v)
	var conds []clause.Expression
	t := v.Type()
	for i := range v.NumField() {
		field := v.Field(i)
		fieldKind := field.Kind()
		if fieldKind == reflect.Interface || fieldKind == reflect.Ptr || fieldKind == reflect.Struct {
			if t.Field(i).Anonymous {
				conds = append(conds, EqualByStruct(field.Interface()))
			}
		} else {
			conds = append(conds, clause.Eq{Column: clause.Column{Table: clause.CurrentTable, Name: t.Field(i).Name}, Value: v.Field(i).Interface()})
		}

	}
	return clause.AndConditions{Exprs: conds}
}

func ConditionByStruct(param any) (clause.Expression, error) {
	v := reflect.ValueOf(param)
	v = reflect.Indirect(v)
	var conds []clause.Expression
	t := v.Type()
	for i := range v.NumField() {
		field := v.Field(i)
		fieldKind := field.Kind()
		if fieldKind == reflect.Interface || fieldKind == reflect.Ptr || fieldKind == reflect.Struct {
			if t.Field(i).Anonymous {
				subCondition, err := ConditionByStruct(field.Interface())
				if err != nil {
					return nil, err
				}
				conds = append(conds, subCondition)
			}
		} else {
			condition, err := sql2.GetSQLCondition(t.Field(i).Tag)
			if err != nil {
				return nil, err
			}
			if condition.Expr != "" {
				conds = append(conds, clause.Expr{SQL: condition.Expr, Vars: []any{v.Field(i).Interface()}})
			} else {
				var column any = condition.Column
				if column == "" {
					column = clause.Column{Table: clause.CurrentTable, Name: t.Field(i).Name}
				}
				if condition.Op == "" {
					condition.Op = "Equal"
				}
				op := sql2.ParseConditionOperation(condition.Op)
				conds = append(conds, NewCondition(column, op, v.Field(i).Interface()))
			}
		}
	}
	return clause.AndConditions{Exprs: conds}, nil
}
