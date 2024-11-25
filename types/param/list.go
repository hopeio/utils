/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package param

type List[T Ordered] struct {
	PageSort
	*Range[T]
}

func NewList[T Ordered](pageNo, pageSize int) *List[T] {
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
