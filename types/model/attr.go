package model

type Attribute struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"comment:名称"`
	Type  int    `json:"type" gorm:"comment:类型"`
	Value string `json:"value" gorm:"值"`
}
