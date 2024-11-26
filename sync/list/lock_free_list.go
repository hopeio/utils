// Copyright 2024 hopeio. All rights reserved.
// Licensed under the MIT License that can be found in the LICENSE file.
// @Created by jyb

package list

import (
	"github.com/hopeio/utils/sync"
	"sync/atomic"
	"unsafe"
)

type LockFreeList struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	size uint64
}

func New() *LockFreeList {
	return &LockFreeList{}
}

func (l *LockFreeList) Push(v interface{}) {
	node := &sync.DirectItem{V: v}
	if l.size == 0 {
		atomic.StorePointer(&l.head, unsafe.Pointer(node))
		atomic.StorePointer(&l.tail, unsafe.Pointer(node))
		atomic.AddUint64(&l.size, 1)
		return
	}
	var last, lastnext *sync.DirectItem
	for {
		last = sync.LoadItem(&l.tail)
		lastnext = sync.LoadItem(&last.Next)
		if sync.LoadItem(&l.tail) == last { // are tail and next consistent?
			if lastnext == nil { // was tail pointing to the last node?
				if sync.CasItem(&last.Next, lastnext, node) { // try to link item at the end of linked list
					sync.CasItem(&l.tail, last, node) // enqueue is done. try swing tail to the inserted node
					atomic.AddUint64(&l.size, 1)
					return
				}
			} else { // tail was not pointing to the last node
				sync.CasItem(&l.tail, last, lastnext) // try swing tail to the next node
			}
		}
	}
}
