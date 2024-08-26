package iter

import (
	"github.com/hopeio/utils/types"
	"github.com/hopeio/utils/types/funcs"
	"iter"
)

// Dont'use please use types.Pair And Seq
type Stream2[K, V any] interface {
	Seq() iter.Seq[types.Pair[K, V]]
	Seq2() iter.Seq2[K, V]

	Filter(funcs.PredicateKV[K, V]) Stream2[K, V]
	Map(funcs.UnaryKVFunction[K, V, V]) Stream2[K, V]                   //同类型转换,没啥意义
	FlatMap(funcs.UnaryKVFunction[K, V, iter.Seq2[K, V]]) Stream2[K, V] //同Map
	Peek(funcs.ConsumerKV[K, V]) Stream2[K, V]

	Distinct(funcs.UnaryKVFunction[K, V, int]) Stream2[K, V]
	SortedByKeys(funcs.Comparator[K]) Stream2[K, V]
	SortedByValues(funcs.Comparator[V]) Stream2[K, V]
	Limit(int64) Stream2[K, V]
	Skip(int64) Stream2[K, V]

	ForEach(funcs.ConsumerKV[K, V])
	Collect() ([]K, []V)
	AllMatch(funcs.PredicateKV[K, V]) bool
	NoneMatch(funcs.PredicateKV[K, V]) bool
	AnyMatch(funcs.PredicateKV[K, V]) bool

	First() (K, V)
	Count() int64
}

type Seq2[K, V any] iter.Seq2[K, V]

func (s Seq2[K, V]) Seq() iter.Seq[types.Pair[K, V]] {
	return func(yield func(types.Pair[K, V]) bool) {
		for k, v := range s {
			if !yield(types.PairOf(k, v)) {
				return
			}
		}
	}
}
