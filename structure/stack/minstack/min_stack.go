/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package minstack

import (
	"container/list"
	"github.com/hopeio/utils/cmp"
)

// MinStack ...
type MinStack[T cmp.Comparable[T]] struct {
	store *list.List
}
type node[T cmp.Comparable[T]] struct {
	min   T
	value T
}

// NewMinStack ...
func NewMinStack[T cmp.Comparable[T]]() MinStack[T] {
	return MinStack[T]{store: list.New()}
}

// Push ...
func (ms *MinStack[T]) Push(x T) {
	if ms.store.Front() != nil && ms.store.Front().Value.(*node[T]).min.Compare(x) <= 0 {
		ms.store.PushFront(&node[T]{value: x, min: ms.store.Front().Value.(*node[T]).min})
	} else {
		ms.store.PushFront(&node[T]{value: x, min: x})
	}
}

// Pop ...
func (ms *MinStack[T]) Pop() {
	ms.store.Remove(ms.store.Front())
}

// Top ...
func (ms *MinStack[T]) Top() T {
	return ms.store.Front().Value.(*node[T]).value
}

// GetMin ...
func (ms *MinStack[T]) GetMin() T {
	return ms.store.Front().Value.(*node[T]).min
}
