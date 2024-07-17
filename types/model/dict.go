package model

type Dict struct {
	Id      int    `gorm:"primaryKey"`
	Name    string `gorm:"comment:名称"`
	Group   int    `gorm:"comment:类型" gorm:"uniqueIndex:idx_group_key"`
	Key     string `gorm:"comment:键" gorm:"uniqueIndex:idx_group_key"`
	Value   string `gorm:"comment:值"`
	Comment string `gorm:"comment:注释"`
}
