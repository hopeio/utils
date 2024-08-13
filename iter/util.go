package iter

import (
	"github.com/hopeio/utils/types"
	"github.com/hopeio/utils/types/constraints"
)

func SliceAll[T any](input []T) Seq[types.Pair[int, T]] {
	return func(yield func(types.Pair[int, T]) bool) {
		for i, v := range input {
			if !yield(types.PairOf(i, v)) {
				return
			}
		}
	}
}

func SliceAll2[T any](input []T) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range input {
			if !yield(i, v) {
				return
			}
		}
	}
}

func SliceValues[T any](input []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}
}

func SliceBackwardValues[T any](input []T) Seq[T] {
	return func(yield func(T) bool) {
		n := len(input) - 1
		for i := n; n > 0; n-- {
			if !yield(input[i]) {
				return
			}
		}
	}
}

func SliceBackward[T any](input []T) Seq[types.Pair[int, T]] {
	return func(yield func(types.Pair[int, T]) bool) {
		n := len(input) - 1
		for i := n; n > 0; n-- {
			if !yield(types.PairOf(i, input[i])) {
				return
			}
		}
	}
}

func SliceBackward2[T any](input []T) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		n := len(input) - 1
		for i := n; n > 0; n-- {
			if !yield(i, input[i]) {
				return
			}
		}
	}
}

func HashMapAll[K comparable, V any](m map[K]V) Seq[types.Pair[K, V]] {
	return func(yield func(types.Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(types.PairOf(k, v)) {
				return
			}
		}
	}
}

func HashMapAll2[K comparable, V any](m map[K]V) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func HashMapValues[K comparable, V any](m map[K]V) Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

func HashMaKeys[K comparable, V any](m map[K]V) Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				return
			}
		}
	}
}

func StringAll(input string) Seq[types.Pair[int, rune]] {
	return func(yield func(types.Pair[int, rune]) bool) {
		for i, v := range input {
			if !yield(types.PairOf(i, v)) {
				return
			}
		}
	}
}

func StringAll2(input string) Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, v := range input {
			if !yield(i, v) {
				return
			}
		}
	}
}

func StringRunes(input string) Seq[rune] {
	return func(yield func(rune) bool) {
		for _, v := range input {
			if !yield(v) {
				return
			}
		}
	}
}

func ChannelAll[T any](c chan T) Seq[T] {
	return func(yield func(T) bool) {
		for v := range c {
			if !yield(v) {
				return
			}
		}
	}
}

func ChannelAll2[T any](c chan T) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var count int
		for v := range c {
			if !yield(count, v) {
				return
			}
			count++
		}
	}
}

func RangeAll[T constraints.Number](begin, end, step T) Seq[T] {
	return func(yield func(T) bool) {
		for v := begin; v <= end; v += step {
			if !yield(v) {
				return
			}
		}
	}
}

func RangeAll2[T constraints.Number](begin, end, step T) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var count int
		for v := begin; v <= end; v += step {
			if !yield(count, v) {
				return
			}
			count++
		}
	}
}
