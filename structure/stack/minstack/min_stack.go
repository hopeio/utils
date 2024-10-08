package minstack

import "container/list"

// MinStack ...
type MinStack struct {
	store *list.List
}
type node struct {
	min   int
	value int
}

// NewMinStack ...
func NewMinStack() MinStack {
	return MinStack{store: list.New()}
}

// Push ...
func (ms *MinStack) Push(x int) {
	if ms.store.Front() != nil && ms.store.Front().Value.(*node).min > x {
		ms.store.PushFront(&node{value: x, min: x})
	} else if ms.store.Front() != nil && ms.store.Front().Value.(*node).min <= x {
		ms.store.PushFront(&node{value: x, min: ms.store.Front().Value.(*node).min})
	} else {
		ms.store.PushFront(&node{value: x, min: x})
	}
}

// Pop ...
func (ms *MinStack[T]) Pop() {
	ms.store.Remove(ms.store.Front())
}

// Top ...
func (ms *MinStack[T]) Top() T {
	return ms.store.Front().Value.(*node).value
}

// GetMin ...
func (ms *MinStack[T]) GetMin() T {
	return ms.store.Front().Value.(*node).min
}
