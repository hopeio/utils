package gorm

import (
	dbi "github.com/hopeio/utils/dao/database"
	"gorm.io/gorm"
)

type Scope func(*gorm.DB) *gorm.DB

func NewScope(field string, op dbi.Operation, args ...interface{}) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(field+op.SQL(), args...)
	}
}

type ChainScope []func(db *gorm.DB) *gorm.DB

// db.Scope(ById(1),ByName("a")).First(v)
func (c ChainScope) ById(id int) ChainScope {
	if id > 0 {
		return c.ByIdNoCheck(id)
	}
	return c
}

func (c ChainScope) ByIdNoCheck(id any) ChainScope {
	return append(c, NewScope(dbi.ColumnId, dbi.Equal, id))
}

func (c ChainScope) ByName(name string) ChainScope {
	if name != "" {
		return c.ByNameNoCheck(name)
	}
	return c
}

func (c ChainScope) ByNameNoCheck(name any) ChainScope {
	return append(c, func(db *gorm.DB) *gorm.DB {
		return db.Where(dbi.NameEqual, name)
	})
}

func (c ChainScope) Exec(db *gorm.DB) *gorm.DB {
	db = db.Scopes(c...)
	return db
}
