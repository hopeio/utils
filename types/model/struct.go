package model

type Struct struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"comment:名称"`
}

type StructField struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	StructID int    `json:"structId" gorm:"comment:结构体id"`
	PID      int    `json:"pid" gorm:"comment:父级id"`
	Name     string `json:"name" gorm:"comment:名称"`
	Type     int    `json:"type" gorm:"comment:类型"`
}

type Object struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	FieldID int    `json:"fieldId" gorm:"comment:字段id"`
	Value   string `json:"value" gorm:"comment:值"`
}
