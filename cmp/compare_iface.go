package cmp

import (
	constraintsi "github.com/hopeio/utils/types/constraints"
	"golang.org/x/exp/constraints"
)

// 包含了CompareLess和IsEqual,尽量统一使用Comparable
type Comparable[T any] interface {
	Compare(T) int
}

type CompareLess[T any] interface {
	Less(T) bool
}

type IsEqual[T any] interface {
	Equal(T) bool
}

type EqualKey[T comparable] interface {
	EqualKey() T
}

// 下面不推荐使用
// 合理使用,如int, 正序 return v,倒序return -v,并适当考虑边界值问题
type CompareKey[T constraints.Ordered] interface {
	CompareKey() T
}

type ComparableKey[T Comparable[T]] interface {
	CompareKey() T
}

// 可以直接用-号
type CompareNumKey[T constraintsi.Number] interface {
	CompareKey() T
}
