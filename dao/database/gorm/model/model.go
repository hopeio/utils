package model

import (
	"gorm.io/gorm"
	"time"
)

type ModelTime struct {
	CreatedAt time.Time      `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"default:null"`
}
