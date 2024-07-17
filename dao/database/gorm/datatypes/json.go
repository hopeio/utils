package datatypes

import (
	dbi "github.com/hopeio/utils/dao/database"
	"github.com/hopeio/utils/dao/database/datatypes"
	"gorm.io/gorm"
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
