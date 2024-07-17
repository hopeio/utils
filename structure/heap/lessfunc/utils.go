package heap

import (
	"github.com/hopeio/utils/cmp"
)

func HeapInit[T any](heap []T, less cmp.LessFunc[T]) {
	// heapify
	n := len(heap)
	for i := n/2 - 1; i >= 0; i-- {
		Down(heap, i, n, less)
	}
}

func Down[T any](heap []T, i0, n int, less cmp.LessFunc[T]) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && less(heap[j2], heap[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !less(heap[j], heap[i]) {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
	}
	return i > i0
}

func Up[T any](heap []T, j int, less cmp.LessFunc[T]) {

	for {
		i := (j - 1) / 2 // parent
		if i == j || !less(heap[j], heap[i]) {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		j = i
	}
}

func Fix[T any](heap []T, i int, less cmp.LessFunc[T]) {
	if !Down(heap, i, len(heap), less) {
		Up(heap, i, less)
	}
}

// 与Down一致，不同的写法
func AdjustDown[T any](heap []T, i int, less cmp.LessFunc[T]) {
	child := leftChild(i)
	for child < len(heap) {
		if child+1 < len(heap) && less(heap[child+1], heap[child]) {
			child++
		}
		if !less(heap[child], heap[i]) {
			break
		}
		heap[i], heap[child] = heap[child], heap[i]
		i = child
		child = leftChild(i)
	}
}

// 与Up一致，不同的写法
func AdjustUp[T any](heap []T, i int, less cmp.LessFunc[T]) {
	p := parent(i)
	for p >= 0 && less(heap[i], heap[p]) {
		heap[i], heap[p] = heap[p], heap[i]
		i = p
		p = parent(i)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}
func leftChild(i int) int {
	return i*2 + 1
}
