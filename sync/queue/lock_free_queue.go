// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package queue

import (
	"github.com/hopeio/utils/sync"
	"sync/atomic"
	"unsafe"
)

// LockFreeQueue implements lock-free FIFO freelist based queue.
// ref: https://dl.acm.org/citation.cfm?doid=248052.248106
type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	len  uint64
}

// NewLockFreeQueue creates a new lock-free queue.
func NewLockFreeQueue() *LockFreeQueue {
	head := sync.DirectItem{Next: nil, V: nil} // allocate a free item
	return &LockFreeQueue{
		tail: unsafe.Pointer(&head), // both head and tail points
		head: unsafe.Pointer(&head), // to the free item
	}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *LockFreeQueue) Enqueue(v interface{}) {
	i := &sync.DirectItem{Next: nil, V: v} // allocate new item
	var last, lastnext *sync.DirectItem
	for {
		last = sync.LoadItem(&q.tail)
		lastnext = sync.LoadItem(&last.Next)
		if sync.LoadItem(&q.tail) == last { // are tail and next consistent?
			if lastnext == nil { // was tail pointing to the last node?
				if sync.CasItem(&last.Next, lastnext, i) { // try to link item at the end of linked list
					sync.CasItem(&q.tail, last, i) // enqueue is done. try swing tail to the inserted node
					atomic.AddUint64(&q.len, 1)
					return
				}
			} else { // tail was not pointing to the last node
				sync.CasItem(&q.tail, last, lastnext) // try swing tail to the next node
			}
		}
	}
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *LockFreeQueue) Dequeue() interface{} {
	var first, last, firstnext *sync.DirectItem
	for {
		first = sync.LoadItem(&q.head)
		last = sync.LoadItem(&q.tail)
		firstnext = sync.LoadItem(&first.Next)
		if first == sync.LoadItem(&q.head) { // are head, tail and next consistent?
			if first == last { // is queue empty?
				if firstnext == nil { // queue is empty, couldn't dequeue
					return nil
				}
				sync.CasItem(&q.tail, last, firstnext) // tail is falling behind, try to advance it
			} else { // read value before cas, otherwise another dequeue might free the next node
				v := firstnext.V
				if sync.CasItem(&q.head, first, firstnext) { // try to swing head to the next node
					atomic.AddUint64(&q.len, ^uint64(0))
					return v // queue was not empty and dequeue finished.
				}
			}
		}
	}
}

// Length returns the length of the queue.
func (q *LockFreeQueue) Length() uint64 {
	return atomic.LoadUint64(&q.len)
}
