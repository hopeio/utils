/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

type Dict struct {
	Type    int    `gorm:"comment:类型" gorm:"primaryKey"`
	Key     string `gorm:"comment:键" gorm:"primaryKey"`
	Name    string `gorm:"comment:名称"`
	Value   string `gorm:"comment:值"`
	Comment string `gorm:"comment:注释"`
}
