package minstack

import (
	"github.com/hopeio/utils/structure/list"
	"golang.org/x/exp/constraints"
)

// MinStack ...
type MinStack[T constraints.Ordered] struct {
	store *list.List[node[T]]
}
type node[T constraints.Ordered] struct {
	min   T
	value T
}

// NewMinStack ...
func NewMinStack[T constraints.Ordered]() MinStack[T] {
	return MinStack[T]{store: list.New[node[T]]()}
}

// Push ...
func (ms *MinStack[T]) Push(x T) {
	if ms.store.Head() != nil && ms.store.Head().Value.min <= x {
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
