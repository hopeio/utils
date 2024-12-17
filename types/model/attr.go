package model

type Attribute struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"comment:名称"`
	Type int    `json:"type" gorm:"comment:类型"`
}

type AttributeValue struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	AttrID int    `json:"attrId" gorm:"comment:属性id" `
	Value  string `json:"value" gorm:"值"`
}
