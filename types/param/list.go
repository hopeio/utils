package param

import (
	"github.com/hopeio/utils/types/constraints"
)

type List[T constraints.Ordered] struct {
	PageSort
	*Range[T]
}

func NewList[T constraints.Ordered](pageNo, pageSize int) *List[T] {
	return &List[T]{
		PageSort: PageSort{
			Page: Page{
				PageNo:   pageNo,
				PageSize: pageSize,
			},
		},
	}
}

func (req *List[T]) WithSort(field string, typ SortType) *List[T] {
	req.Sort = &Sort{
		SortField: field,
		SortType:  typ,
	}
	return req
}

func (req *List[T]) WithRange(field string, start, end T, include bool) *List[T] {
	req.Range = &Range[T]{
		RangeField: field,
		RangeBegin: start,
		RangeEnd:   end,
		Include:    include,
	}
	return req
}
