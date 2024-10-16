// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package stack

import (
	"github.com/hopeio/utils/sync"
	"sync/atomic"
	"unsafe"
)

// LockFreeStack implements lock-free freelist based stack.
type LockFreeStack struct {
	top unsafe.Pointer
	len uint64
}

// NewLockFreeStack creates a new lock-free queue.
func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

// Pop pops value from the top of the stack.
func (s *LockFreeStack) Pop() interface{} {
	var top, next unsafe.Pointer
	var item *sync.DirectItem
	for {
		top = atomic.LoadPointer(&s.top)
		if top == nil {
			return nil
		}
		item = (*sync.DirectItem)(top)
		next = atomic.LoadPointer(&item.Next)
		if atomic.CompareAndSwapPointer(&s.top, top, next) {
			atomic.AddUint64(&s.len, ^uint64(0))
			return item.V
		}
	}
}

// Push pushes a value on top of the stack.
func (s *LockFreeStack) Push(v interface{}) {
	item := sync.DirectItem{V: v}
	var top unsafe.Pointer
	for {
		top = atomic.LoadPointer(&s.top)
		item.Next = top
		if atomic.CompareAndSwapPointer(&s.top, top, unsafe.Pointer(&item)) {
			atomic.AddUint64(&s.len, 1)
			return
		}
	}
}
