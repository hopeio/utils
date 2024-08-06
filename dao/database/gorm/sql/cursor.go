package sql

import (
	"github.com/hopeio/utils/dao/database/sql"
	"github.com/hopeio/utils/types/model"
	"gorm.io/gorm"
)

func GetCursor(db *gorm.DB, typ string) (*model.Cursor, error) {
	var cursor model.Cursor
	err := db.Where(`type = ?`, typ).First(&cursor).Error
	if err != nil {
		return nil, err
	}
	return &cursor, nil
}

func EndCallback(db *gorm.DB, typ string) {
	db.Exec(sql.EndCallbackSQL(typ))
}
