/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package sync

import (
	"sync/atomic"
	"unsafe"
)

type Node[T any] struct {
	Next atomic.Pointer[Node[T]]
	V    T
}

func LoadNode[T any](p *unsafe.Pointer) *Node[T] {
	return (*Node[T])(atomic.LoadPointer(p))
}
func CasNode[T any](p *unsafe.Pointer, old, new *Node[T]) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
