// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package stack

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

func TestStackPopEmpty(t *testing.T) {
	s := NewLockFreeStack[int]()
	if _, ok := s.Pop(); ok {
		t.Fatal("pop empty stack returns non-nil")
	}
}

func ExampleLockFreeStack() {
	s := NewLockFreeStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	// 3 true
	// 2 true
	// 1 true
}

type stackInterface[T any] interface {
	Push(T)
	Pop() (T, bool)
}

func BenchmarkStack(b *testing.B) {
	length := 1 << 12
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs = append(inputs, rand.Int())
	}
	s, ms := NewLockFreeStack[int](), NewMutexStack[int]()
	b.ResetTimer()
	for _, s := range [...]stackInterface[int]{s, ms} {
		b.Run(fmt.Sprintf("%T", s), func(b *testing.B) {
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 0 {
						s.Push(v)
					} else {
						s.Pop()
					}
				}
			})
		})
	}
}
