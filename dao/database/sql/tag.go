package sql

import (
	"github.com/hopeio/utils/reflect/structtag"
	"reflect"
)

const (
	CondiTagName = "sqlcondi" // e.g: `sqlcondi:"column:id;op:="`
	// e.g: `sqlcondi:"expr:id = ?"`
	// e.g: `sqlcondi:"-"`
)

type ConditionTag struct {
	Column string `meta:"column"`
	Expr   string `meta:"expr"`
	Op     string `meta:"op"`
}

func GetSQLCondition(tag reflect.StructTag) (*ConditionTag, error) {
	return structtag.ParseSettingTagToStruct[ConditionTag](tag.Get(CondiTagName), ';')
}
