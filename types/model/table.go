package model

type TableMeta struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"comment:名称"`
}

type TableField struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"comment:名称"`
	Type    string `json:"type" gorm:"comment:类型"`
	TableID int    `json:"tableId" gorm:"comment:表id"`
}

type TableValue struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	FieldID int    `json:"fieldId" gorm:"comment:字段id"`
	Value   string `json:"value" gorm:"comment:值"`
}
