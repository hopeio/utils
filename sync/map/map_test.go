package smap

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	length := 1 << 12
	inputs := make([]int, length)
	for range length {
		inputs = append(inputs, rand.Int())
	}
	syncMap := sync.Map{}
	lockMap := New[int, any]()
	b.ResetTimer()

	b.Run(fmt.Sprintf("%T", syncMap), func(b *testing.B) {
		var c int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				i := int(atomic.AddInt64(&c, 1)-1) % length
				v := inputs[i]
				if v >= 0 {
					syncMap.Store(v, 1)
				} else {
					syncMap.Load(v)
				}
			}
		})
	})
	b.Run(fmt.Sprintf("%T", lockMap), func(b *testing.B) {
		var c int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				i := int(atomic.AddInt64(&c, 1)-1) % length
				v := inputs[i]
				if v >= 0 {
					lockMap.Set(v, 1)
				} else {
					lockMap.Get(v)
				}
			}
		})
	})
}
