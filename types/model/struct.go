package model

// 内部保留,{1,int8},{2,int16}...{x,uint8}...{x,uint64},{x,string}
type Struct struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"comment:名称"`
}

type StructField struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	StructID    int    `json:"structId" gorm:"comment:结构体id"`
	Name        string `json:"name" gorm:"comment:名称"`
	Kind        uint32 `json:"kind" gorm:"comment:类型"`
	Tag         string `json:"tag" gorm:"comment:标签"`
	Offset      uint32 `json:"offset" gorm:"comment:偏移量"`
	Anonymous   bool   `json:"anonymous" gorm:"comment:匿名"`
	SubStructID int    `json:"subStructId" gorm:"comment:子结构体id"`
}

type Object struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	FieldID int    `json:"fieldId" gorm:"comment:字段id"`
	Value   string `json:"value" gorm:"comment:值"`
}
