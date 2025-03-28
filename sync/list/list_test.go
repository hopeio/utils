package list

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

type listInterface[T any] interface {
	Push(v T)
	Pop() (T, bool)
}

func BenchmarkList(b *testing.B) {
	length := 1 << 12
	inputs := make([]int, length)
	for range length {
		inputs = append(inputs, rand.Int())
	}
	q, mq := NewLockFreeList[int](), NewMutexList[int]()
	b.ResetTimer()

	for _, q := range [...]listInterface[int]{q, mq} {
		b.Run(fmt.Sprintf("%T", q), func(b *testing.B) {
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 0 {
						q.Push(v)
					} else {
						q.Pop()
					}
				}
			})
		})
	}
}
