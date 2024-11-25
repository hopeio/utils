/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

import (
	constraintsi "github.com/hopeio/utils/types/constraints"
)

type String string

func (s String) Key() string {
	return string(s)
}

type Int int

func (s Int) Key() int {
	return int(s)
}

type Basic struct {
}

type ID[T constraintsi.ID] struct {
	Id T `json:"id"`
}

func (s ID[KEY]) Key() KEY {
	return s.Id
}
