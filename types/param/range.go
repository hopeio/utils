package param

import (
	"github.com/hopeio/utils/types/constraints"
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

type Range[T constraints.Range] struct {
	RangeField string `json:"rangeField,omitempty"`
	RangeBegin T      `json:"rangeBegin,omitempty"`
	RangeEnd   T      `json:"rangeEnd,omitempty"`
	Include    bool   `json:"include,omitempty"`
}
