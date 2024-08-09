package model

type Dict struct {
	Type    int    `gorm:"comment:类型" gorm:"primaryKey"`
	Key     string `gorm:"comment:键" gorm:"primaryKey"`
	Name    string `gorm:"comment:名称"`
	Value   string `gorm:"comment:值"`
	Comment string `gorm:"comment:注释"`
}
