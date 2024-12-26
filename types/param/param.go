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
	Type() SortType
}

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

type Range[T Rangeable] struct {
	RangeField string    `json:"rangeField,omitempty"`
	RangeBegin T         `json:"rangeBegin,omitempty"`
	RangeEnd   T         `json:"rangeEnd,omitempty"`
	RangeType  RangeType `json:"include,omitempty"`
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
	RangeTypeHasBegin RangeType = 1 << iota
	RangeTypeHasEnd
	RangeTypeContainsBegin
	RangeTypeContainsEnd
)
