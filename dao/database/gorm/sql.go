package gorm

import (
	dbi "github.com/hopeio/utils/dao/database"
	"gorm.io/gorm"
)

func DeleteById(db *gorm.DB, tableName string, id uint64) error {
	sql := dbi.DeleteSQL(tableName, "id")
	return db.Exec(sql, id).Error
}

func Delete(db *gorm.DB, tableName string, column string, value any) error {
	sql := dbi.DeleteSQL(tableName, column)
	return db.Exec(sql, value).Error
}

func ExistsByIdWithDeletedAt(db *gorm.DB, tableName string, id uint64) (bool, error) {
	return ExistsBySQL(db, dbi.ExistsSQL(tableName, "id", true), id)
}

func ExistsById(db *gorm.DB, tableName string, id uint64) (bool, error) {
	return ExistsBySQL(db, dbi.ExistsSQL(tableName, "id", false), id)
}

func ExistsByColumn(db *gorm.DB, tableName, column string, value interface{}) (bool, error) {
	return ExistsBySQL(db, dbi.ExistsSQL(tableName, column, false), value)
}

func ExistsBySQL(db *gorm.DB, sql string, value ...any) (bool, error) {
	var exists bool
	err := db.Raw(sql, value...).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

// 根据查询语句查询数据是否存在
func ExistsByQuerySQL(db *gorm.DB, qsql string, value ...any) (bool, error) {
	var exists bool
	err := db.Raw(dbi.ExistsSQLByQuerySQL(qsql), value...).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}

func Exists(db *gorm.DB, tableName, column string, value interface{}, withDeletedAt bool) (bool, error) {
	return ExistsBySQL(db, dbi.ExistsSQL(tableName, column, withDeletedAt), value)
}

func ExistsByFilterExprs(db *gorm.DB, tableName string, filters dbi.FilterExprs) (bool, error) {
	var exists bool
	err := db.Raw(dbi.ExistsSQLByFilterExprs(tableName, filters)).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
