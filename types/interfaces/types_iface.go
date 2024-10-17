package interfaces

import "github.com/hopeio/utils/types/constraints"

type Key[T constraints.Key] interface {
	Key() T
}

type Collector[S any, T any, R any] interface {
	Builder() S
	Append(builder S, element T)
	Finish(builder S) R
}
