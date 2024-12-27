package model

// 内部保留,{1,int8},{2,int16}...{x,uint8}...{x,uint64},{x,string}
type Struct struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"comment:名称"`
}

type StructField struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	StructID    uint   `json:"structId" gorm:"comment:结构体id"`
	Name        string `json:"name" gorm:"comment:名称"`
	Kind        uint32 `json:"kind" gorm:"comment:类型"`
	Tag         string `json:"tag" gorm:"comment:标签"`
	Offset      uint32 `json:"offset" gorm:"comment:偏移量"`
	Anonymous   bool   `json:"anonymous" gorm:"comment:匿名"`
	SubStructID uint   `json:"subStructId" gorm:"comment:子结构体id"`
}

type Object struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	FieldID uint   `json:"fieldId" gorm:"comment:字段id"`
	Value   string `json:"value" gorm:"comment:值"`
}
