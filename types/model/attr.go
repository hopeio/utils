package model

type Attribute struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"comment:名称"`
	Type  uint32 `json:"type" gorm:"comment:类型"`
	Value string `json:"value" gorm:"值"`
}
