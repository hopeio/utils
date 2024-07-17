package model

type Cursor struct {
	Type   string `json:"type" gorm:"primaryKey"`
	Cursor string
	Prev   string
	Next   string
}
