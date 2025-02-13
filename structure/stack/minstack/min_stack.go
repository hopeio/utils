/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package minstack

import (
	"github.com/hopeio/utils/cmp"
	"github.com/hopeio/utils/structure/list"
)

// MinStack ...
type MinStack[T any] struct {
	store *list.List[node[T]]
	less  cmp.LessFunc[T]
}

type node[T any] struct {
	value T
	min   T
}

// NewMinStack ...
func NewMinStack[T any](less cmp.LessFunc[T]) MinStack[T] {
	return MinStack[T]{store: list.New[node[T]](), less: less}
}

// Push ...
func (ms *MinStack[T]) Push(x T) {
	if ms.store.Head() != nil && ms.less(ms.store.Head().Value.min, x) {
		ms.store.PushFront(node[T]{value: x, min: ms.store.Head().Value.min})
	} else {
		ms.store.PushFront(node[T]{value: x, min: x})
	}
}

// Pop ...
func (ms *MinStack[T]) Pop() {
	ms.store.Pop()
}

// Top ...
func (ms *MinStack[T]) Top() T {
	return ms.store.Head().Value.value
}

// GetMin ...
func (ms *MinStack[T]) GetMin() T {
	return ms.store.Head().Value.min
}
