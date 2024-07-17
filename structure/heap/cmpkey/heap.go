package heap

import (
	"github.com/hopeio/utils/cmp"
	"golang.org/x/exp/constraints"
)

type Heap[T cmp.CompareKey[V], V constraints.Ordered] []T

func New[T cmp.CompareKey[V], V constraints.Ordered](l int) Heap[T, V] {
	return make([]T, 0, l)
}

func NewFromArray[T cmp.CompareKey[V], V constraints.Ordered](arr []T) Heap[T, V] {
	heap := Heap[T, V](arr)
	for i := 1; i < len(arr); i++ {
		Up(heap, i)
	}
	return heap
}

func (h Heap[T, V]) Init() {
	// heapify
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		Down(h, i, n)
	}
}

// 当达到堆预设大小时会增加堆的大小
func (heap *Heap[T, V]) Push(x T) {
	*heap = append(*heap, x)
	h := *heap
	Up(h, len(h)-1)
}

// 不会改变预设堆的大小
func (heap *Heap[T, V]) Put(val T) {
	h := *heap
	if len(h) < cap(h) {
		h = append(h, val)
		for i := 1; i < len(h); i++ {
			Up(h, i)
		}
		return
	}
	if val.CompareKey() > h[0].CompareKey() {
		return
	}
	h[0] = val
	Down(h, 0, len(h))
	*heap = h
}

func (heap *Heap[T, V]) Pop() (T, bool) {
	h := *heap
	if len(h) == 0 {
		return *new(T), false
	}
	n := len(h) - 1
	item := h[0]
	h[0], h[n] = h[n], h[0]
	Down(h, 0, n)
	*heap = h[:n]
	return item, true
}

func (heap *Heap[T, V]) Remove(i int) (T, bool) {
	h := *heap
	if len(h) == 0 {
		return *new(T), false
	}
	n := len(h) - 1
	item := h[i]
	if n != i {
		h[i], h[n] = h[n], h[i]
		if !Down(h, i, n) {
			Up(h, i)
		}
	}
	*heap = h[:n]
	return item, true
}

func (heap Heap[T, V]) First() (T, bool) {
	if len(heap) == 0 {
		return *new(T), false
	}
	return heap[0], true
}

func (heap Heap[T, V]) Last() (T, bool) {
	if len(heap) == 0 {
		return *new(T), false
	}
	return heap[len(heap)-1], false
}
