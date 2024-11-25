/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package param

import (
	"time"
)

type SortType int

const (
	_ SortType = iota
	SortTypeAsc
	SortTypeDesc
)

type PageSort struct {
	Page
	*Sort
}

type Page struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type Sort struct {
	SortField string   `json:"sortField,omitempty"`
	SortType  SortType `json:"sortType,omitempty"`
}

func (receiver *Sort) Column() string {
	return receiver.SortField
}

func (receiver *Sort) Type() SortType {
	return receiver.SortType
}

type DateRange[T ~string | time.Time] Range[T]

type Range[T Rangeable] struct {
	RangeField string `json:"rangeField,omitempty"`
	RangeBegin T      `json:"rangeBegin,omitempty"`
	RangeEnd   T      `json:"rangeEnd,omitempty"`
	Include    bool   `json:"include,omitempty"`
}
