/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

import "github.com/hopeio/utils/dao/database/datatypes"

type Enum struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"comment:名称"`
	Group   int    `gorm:"comment:枚举组" gorm:"uniqueIndex:idx_group_index"`
	Index   int    `gorm:"comment:索引" gorm:"uniqueIndex:idx_group_index"`
	Value   string `gorm:"comment:值" `
	Comment string `gorm:"comment:注释"`
}

type PostgresEnum struct {
	ID    int `gorm:"primaryKey"`
	Enums datatypes.StringArray
}

type EnumItem struct {
	Group   int    `gorm:"comment:枚举类型"`
	Index   int    `gorm:"comment:索引"`
	Comment string `gorm:"comment:注释"`
}
