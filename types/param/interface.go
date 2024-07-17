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
	Column() string
	Type() SortType
}
