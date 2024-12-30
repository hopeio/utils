package model

import (
	"time"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty" gorm:"index"`
}

type ModelExt struct {
	OperatorId uint `json:"operatorId"`
}
