package model

type TableMeta struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"comment:名称"`
}

type TableColumn struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"comment:名称"`
	Type    string `json:"type" gorm:"comment:类型"`
	TableID uint   `json:"tableId" gorm:"comment:表id"`
}

type TableValue struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	ColumnID uint   `json:"columnId" gorm:"comment:字段id"`
	Value    string `json:"value" gorm:"comment:值"`
}
