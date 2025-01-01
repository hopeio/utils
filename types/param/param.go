/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package param

type IPageSort interface {
	IPage
	ISort
}

type IPage interface {
	PageNo() int
	PageSize() int
}

type ISort interface {
	SortType() SortType
}

type SortType int

const (
	_ SortType = iota
	SortTypeAsc
	SortTypeDesc
)

type PageSortEmbed struct {
	PageEmbed
	*SortEmbed
}

type PageEmbed struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type SortEmbed struct {
	SortField string   `json:"sortField,omitempty"`
	SortType  SortType `json:"sortType,omitempty"`
}

type PageSort struct {
	Page Page  `json:"page"`
	Sort *Sort `json:"sort,omitempty"`
}

type Page struct {
	No   int `json:"no"`
	Size int `json:"size"`
}

type Sort struct {
	Field string   `json:"field,omitempty"`
	Type  SortType `json:"type,omitempty"`
}

type Range[T Rangeable] struct {
	Field string    `json:"field,omitempty"`
	Begin T         `json:"begin,omitempty"`
	End   T         `json:"end,omitempty"`
	Type  RangeType `json:"type,omitempty"`
}

type Id struct {
	Id uint `json:"id"`
}

type RangeType int8

func (r RangeType) HasBegin() bool {
	return r&RangeTypeHasBegin != 0
}

func (r RangeType) HasEnd() bool {
	return r&RangeTypeHasEnd != 0
}

func (r RangeType) ContainsBegin() bool {
	return r&RangeTypeContainsBegin != 0
}

func (r RangeType) ContainsEnd() bool {
	return r&RangeTypeContainsEnd != 0
}

const (
	RangeTypeContainsEnd RangeType = 1 << iota
	RangeTypeContainsBegin
	RangeTypeHasEnd
	RangeTypeHasBegin
)
