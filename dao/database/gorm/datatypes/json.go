package datatypes

import (
	"context"
	"database/sql/driver"
	dbi "github.com/hopeio/utils/dao/database"
	"github.com/hopeio/utils/dao/database/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type JsonT[T any] datatypes.JsonT[T]

func (*JsonT[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case dbi.Sqlite, dbi.Mysql:
		return "json"
	case dbi.Postgres:
		return "jsonb"
	}
	return ""
}

func (j JsonT[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	v, _ := (datatypes.JsonT[T])(j).Value()
	return clause.Expr{
		SQL:  "?",
		Vars: []any{v},
	}
}

func (j JsonT[T]) Value() (driver.Value, error) {
	// Scan a value into struct from database driver
	return (datatypes.JsonT[T])(j).Value()
}

func (j *JsonT[T]) Scan(v any) error {
	// Scan a value into struct from database driver
	return (*datatypes.JsonT[T])(j).Scan(v)
}
