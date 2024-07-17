package model

import (
	"time"
)

type ModelTime struct {
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:null"`
}

type ModelTimeStamp struct {
	CreatedAt int64 `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt int64 `json:"updated_at" gorm:"default:current_timestamp"`
	DeletedAt int64 `json:"deleted_at" gorm:"index"`
}
