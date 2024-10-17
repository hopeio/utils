package interfaces

import "github.com/hopeio/utils/types/param"

type PageSort interface {
	Page
	Sort
}

type Page interface {
	PageNo() int
	PageSize() int
}

type Sort interface {
	Column() string
	Type() param.SortType
}
