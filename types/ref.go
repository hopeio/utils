/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package types

type Ref[T any] struct {
	value *T
}

func RefOf[T any](v *T) Ref[T] {
	return Ref[T]{v}
}

func (a Ref[T]) Val() (v T, ok bool) {
	if a.value == nil {
		return
	}
	return *a.value, true
}

func (a Ref[T]) Get() T {
	return *a.value
}

func (a Ref[T]) Set(v T) T {
	var old = *a.value
	*a.value = v
	return old
}

func (a Ref[T]) IsNil() bool {
	return a.value == nil
}

func (a Ref[T]) IsNotNil() bool {
	return a.value != nil
}
