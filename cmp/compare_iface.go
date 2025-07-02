/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package cmp

import (
	"golang.org/x/exp/constraints"
)

// 包含了CompareLess和IsEqual,尽量统一使用Comparable
type Comparable[T any] interface {
	Compare(T) int
}

// Deprecated: use cmp.Comparable
type CompareLess[T any] interface {
	Less(T) bool
}

type ComparableIdx[T any] interface {
	Compare(i, j int) int
}

// Deprecated: use cmp.Comparable
type CompareIdxLess[T any] interface {
	Less(i, j int) bool
}

type IsEqual[T any] interface {
	Equal(T) bool
}

type EqualKey[T comparable] interface {
	EqualKey() T
}

// 下面不推荐使用
// 合理使用,如int, 正序 return v,倒序return -v,并适当考虑边界值问题
// Deprecated: use cmp.Comparable
type CompareKey[T constraints.Ordered] interface {
	CompareKey() T
}
