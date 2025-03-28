// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package queue

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

func TestQueueDequeueEmpty(t *testing.T) {
	q := NewLockFreeQueue[any]()
	if _, ok := q.Dequeue(); ok {
		t.Fatalf("dequeue empty queue returns non-nil")
	}
}

func TestQueue_Length(t *testing.T) {
	q := NewLockFreeQueue[any]()
	if q.Len() != 0 {
		t.Fatalf("empty queue has non-zero length")
	}

	q.Enqueue(1)
	if q.Len() != 1 {
		t.Fatalf("count of enqueue wrong, want %d, got %d.", 1, q.Len())
	}

	q.Dequeue()
	if q.Len() != 0 {
		t.Fatalf("count of dequeue wrong, want %d, got %d", 0, q.Len())
	}
}

func ExampleLockFreeQueue() {
	q := NewLockFreeQueue[string]()

	q.Enqueue("1st item")
	q.Enqueue("2nd item")
	q.Enqueue("3rd item")

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	// Output:
	// 1st item true
	// 2nd item true
	// 3rd item true
}

type queueInterface[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)
}

// 数组锁比链表无锁快
func BenchmarkQueue(b *testing.B) {
	length := 1 << 12
	inputs := make([]int, length)
	for range length {
		inputs = append(inputs, rand.Int())
	}
	q, mq := NewLockFreeQueue[int](), NewMutexQueue[int]()
	b.ResetTimer()

	for _, q := range [...]queueInterface[int]{q, mq} {
		b.Run(fmt.Sprintf("%T", q), func(b *testing.B) {
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 0 {
						q.Enqueue(v)
					} else {
						q.Dequeue()
					}
				}
			})
		})
	}
}
