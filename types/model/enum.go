/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

import "github.com/hopeio/utils/dao/database/datatypes"

type Enum struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"index;comment:名称"`
	Group uint32 `json:"group" gorm:"comment:枚举组"`
	PID   int    `json:"pId" gorm:"comment:父id"`
}

type PostgresEnum struct {
	ID    int `gorm:"primaryKey"`
	Enums datatypes.StringArray
}

type EnumValue struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	EnumID int    `json:"enumId" gorm:"index;comment:枚举id"`
	Index  uint32 `json:"index" gorm:"comment:索引"`
	Value  string `json:"value" gorm:"comment:值"`
	Type   uint32 `json:"type" gorm:"comment:类型"`
}
