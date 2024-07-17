package gorm

import (
	"gorm.io/gorm"
)

type BaseDao[T any] struct {
	*gorm.DB
}

func NewRepository[T any](db *gorm.DB) *BaseDao[T] {
	return &BaseDao[T]{db}
}

func (r *BaseDao[T]) Create(t *T) error {
	return r.DB.Create(t).Error
}

func (r *BaseDao[T]) Retrieve(id any) (*T, error) {
	var t T
	err := r.DB.First(&t, id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *BaseDao[T]) Update(t *T) error {
	return r.DB.Updates(&t).Error
}

func (r *BaseDao[T]) Delete(id any) error {
	var t T
	return r.DB.Delete(&t, id).Error
}
