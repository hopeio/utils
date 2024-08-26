package iter

import (
	"github.com/hopeio/utils/types"
	"github.com/hopeio/utils/types/funcs"
	"iter"
	"slices"
)

type Stream[T any] interface {
	Seq() iter.Seq[T]

	Filter(funcs.Predicate[T]) Stream[T]
	Map(funcs.UnaryFunction[T, T]) Stream[T]               //同类型转换,没啥意义
	FlatMap(funcs.UnaryFunction[T, iter.Seq[T]]) Stream[T] //同Map
	Peek(funcs.Consumer[T]) Stream[T]
	Sorted(funcs.Comparator[T]) Stream[T]
	Distinct(funcs.UnaryFunction[T, int]) Stream[T]
	Limit(int) Stream[T]
	Until(funcs.Predicate[T]) Stream[T]
	Skip(int) Stream[T]

	ForEach(funcs.Consumer[T])
	Collect() []T
	IsSorted(funcs.Comparator[T]) bool
	All(funcs.Predicate[T]) bool // every
	None(funcs.Predicate[T]) bool
	Any(funcs.Predicate[T]) bool // some
	Reduce(acc funcs.BinaryOperator[T]) (T, bool)
	Fold(initVal T, acc funcs.BinaryOperator[T]) T
	First() (T, bool)
	Count() int
	Sum(funcs.BinaryOperator[T]) T
}

func StreamOf[T any](seq iter.Seq[T]) Stream[T] {
	return Seq[T](seq)
}

func Seq2Seq[K, V any](s iter.Seq2[K, V]) iter.Seq[types.Pair[K, V]] {
	return func(yield func(types.Pair[K, V]) bool) {
		for k, v := range s {
			if !yield(types.PairOf(k, v)) {
				return
			}
		}
	}
}

func Seq2Keys[K, V any](s iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range s {
			if !yield(k) {
				return
			}
		}
	}
}

func Seq2Values[K, V any](s iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

type Seq[T any] iter.Seq[T]

func (it Seq[T]) Seq() iter.Seq[T] {
	return iter.Seq[T](it)
}

func (it Seq[T]) Filter(test funcs.Predicate[T]) Stream[T] {
	return Seq[T](Filter(iter.Seq[T](it), test))
}

func (it Seq[T]) Map(f funcs.UnaryFunction[T, T]) Stream[T] {
	return Seq[T](Map(iter.Seq[T](it), f))
}

func (it Seq[T]) FlatMap(f funcs.UnaryFunction[T, iter.Seq[T]]) Stream[T] {
	return Seq[T](FlatMap(iter.Seq[T](it), f))
}

func (it Seq[T]) Peek(accept funcs.Consumer[T]) Stream[T] {
	return Seq[T](Peek(iter.Seq[T](it), accept))
}

func (it Seq[T]) Distinct(f funcs.UnaryFunction[T, int]) Stream[T] {
	return Seq[T](Distinct(iter.Seq[T](it), f))
}

func (it Seq[T]) Sorted(cmp funcs.Comparator[T]) Stream[T] {
	return Seq[T](Sorted(iter.Seq[T](it), cmp))
}

func (it Seq[T]) IsSorted(cmp funcs.Comparator[T]) bool {
	return IsSorted(iter.Seq[T](it), cmp)
}

func (it Seq[T]) Limit(limit int) Stream[T] {
	return Seq[T](Limit(iter.Seq[T](it), limit))
}

func (it Seq[T]) Until(test funcs.Predicate[T]) Stream[T] {
	return Seq[T](Until(iter.Seq[T](it), test))
}

func (it Seq[T]) Skip(skip int) Stream[T] {
	return Seq[T](Skip(iter.Seq[T](it), skip))
}

func (it Seq[T]) ForEach(accept funcs.Consumer[T]) {
	ForEach(iter.Seq[T](it), accept)
}

func (it Seq[T]) Collect() []T {
	return slices.Collect(iter.Seq[T](it))
}

func (it Seq[T]) All(test funcs.Predicate[T]) bool {
	return AllMatch(iter.Seq[T](it), test)
}

func (it Seq[T]) None(test funcs.Predicate[T]) bool {
	return NoneMatch(iter.Seq[T](it), test)
}

func (it Seq[T]) Any(test funcs.Predicate[T]) bool {
	return AnyMatch(iter.Seq[T](it), test)
}

func (it Seq[T]) Reduce(acc funcs.BinaryOperator[T]) (T, bool) {
	return Reduce(iter.Seq[T](it), acc)
}

func (it Seq[T]) Fold(initVal T, acc funcs.BinaryOperator[T]) T {
	return Fold(iter.Seq[T](it), initVal, funcs.BinaryFunction[T, T, T](acc))
}

func (it Seq[T]) First() (T, bool) {
	return First(iter.Seq[T](it))
}

func (it Seq[T]) Count() int {
	return Count(iter.Seq[T](it))
}

func (it Seq[T]) Sum(add funcs.BinaryOperator[T]) T {
	return Operator(iter.Seq[T](it), add)
}

func (it Seq[T]) Iter() Iterator[T] {
	next, stop := iter.Pull(iter.Seq[T](it))
	return &seqIter[T]{next, stop}
}

func SeqSeq2[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var count int
		for v := range seq {
			if !yield(count, v) {
				return
			}
			count++
		}
	}
}
