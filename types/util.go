/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

import (
	constraintsi "github.com/hopeio/gox/types/constraints"
	"golang.org/x/exp/constraints"
)

func CastSigned[T, V constraints.Signed](v V) T {
	return T(v)
}

func CastFloat[T, V constraints.Float](v V) T {
	return T(v)
}

func CastUnsigned[T, V constraints.Unsigned](v V) T {
	return T(v)
}

func CastInteger[T, V constraints.Integer](v V) T {
	return T(v)
}

func CastNumber[T, V constraintsi.Number](v V) T {
	return T(v)
}

func Match[T any](yes bool, a, b T) T {
	if yes {
		return a
	}
	return b
}

func Zero[T any]() T {
	return *new(T)
}

// can compile,but will panic
func none[T any]() T {
	return *(*T)(nil)
}

func Zero2[T any]() T {
	var zero T
	return zero
}

// 两种转换,any(i).(T), T(any(i))
