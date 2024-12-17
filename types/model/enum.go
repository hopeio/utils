/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

import "github.com/hopeio/utils/dao/database/datatypes"

type Enum struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"comment:名称"`
	Group int    `json:"group" gorm:"uniqueIndex;comment:枚举组"`
}

type PostgresEnum struct {
	ID    int `gorm:"primaryKey"`
	Enums datatypes.StringArray
}

type EnumValue struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	EnumID int    `json:"enumId" gorm:"comment:枚举id"`
	Index  int    `json:"index" gorm:"comment:索引"`
	Value  string `json:"value" gorm:"comment:值"`
}
