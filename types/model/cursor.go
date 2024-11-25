/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package model

type Cursor struct {
	Type   string `json:"type" gorm:"primaryKey"`
	Cursor string
	Prev   string
	Next   string
}
