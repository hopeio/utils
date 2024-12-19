/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

type Dict struct {
	Group uint32 `json:"group" gorm:"primaryKey;comment:组"`
	Key   string `json:"key" gorm:"primaryKey;comment:键"`
	Value string `json:"value" gorm:"comment:值"`
	Type  uint32 `json:"type" gorm:"comment:类型"`
	Seq   uint32 `json:"seq" gorm:"comment:排序"`
}
