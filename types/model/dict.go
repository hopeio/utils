/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

type Dict struct {
	Type  int    `json:"type" gorm:"primaryKey;comment:类型"`
	Key   string `json:"key" gorm:"primaryKey;comment:键"`
	Value string `json:"value" gorm:"comment:值"`
	Seq   int    `json:"seq" gorm:"comment:排序"`
}
