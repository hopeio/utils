package gorm

import "gorm.io/gorm"

func GetById[T any](db *gorm.DB, id any) (*T, error) {
	t := new(T)
	err := db.First(t, id).Error
	return t, err
}
