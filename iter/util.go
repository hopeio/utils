package iter

import (
	"github.com/hopeio/utils/cmp"
	"github.com/hopeio/utils/types"
	constraintsi "github.com/hopeio/utils/types/constraints"
	"github.com/hopeio/utils/types/funcs"
	"github.com/hopeio/utils/types/interfaces"
	"golang.org/x/exp/constraints"
	"iter"
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

func HashMapAll[K comparable, V any](m map[K]V) Seq[types.Pair[K, V]] {
	return func(yield func(types.Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(types.PairOf(k, v)) {
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

func StringAll2(input string) iter.Seq2[int, rune] {
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

func ChannelAll2[T any](c chan T) iter.Seq2[int, T] {
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

func RangeAll[T constraintsi.Number](begin, end, step T) Seq[T] {
	return func(yield func(T) bool) {
		for v := begin; v <= end; v += step {
			if !yield(v) {
				return
			}
		}
	}
}

func RangeAll2[T constraintsi.Number](begin, end, step T) iter.Seq2[int, T] {
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

// Returns the sum of all the elements in the iterator.
func Sum[T constraintsi.Number](it iter.Seq[T]) T {
	return Fold(it, 0, func(a, b T) T {
		return a + b
	})
}

func OperatorBy[T any](it iter.Seq[T], f funcs.BinaryOperator[T]) T {
	result, _ := Reduce(it, func(a, b T) T {
		return f(a, b)
	})
	return result
}

// Returns the product of all the elements in the iterator.
func Product[T constraintsi.Number](it iter.Seq[T]) T {
	return Fold(it, 1, func(a, b T) T {
		return a * b
	})
}

// Returns the average of all the elements in the iterator.
func Average[T constraintsi.Number](it iter.Seq[T]) T {
	return Fold(Enumerate(it), *new(T), func(result T, item types.Pair[int, T]) T {
		return result + (T(item.Second)-result)/T(item.First+1)
	})
}

// Ruturns true if the count of Iterator is 0.
func IsEmpty[T any](it iter.Seq[T]) bool {
	for _ = range it {
		return false
	}
	return true
}

// Ruturns true if the count of Iterator is 0.
func IsNotEmpty[T any](it iter.Seq[T]) bool {
	for _ = range it {
		return true
	}
	return false
}

// Returns true if the target is included in the iterator.
func Contains[T comparable](it iter.Seq[T], target T) bool {
	for v := range it {
		if v == target {
			return true
		}
	}
	return false
}

// Return the maximum value of all elements of the iterator.
func Max[T constraints.Ordered](it iter.Seq[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if a > b {
			return a
		} else {
			return b
		}
	})
}

// Return the maximum value of all elements of the iterator.
func MaxBy[T any](it iter.Seq[T], greater cmp.LessFunc[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if greater(a, b) {
			return a
		} else {
			return b
		}
	})
}

// Return the minimum value of all elements of the iterator.
func Min[T constraints.Ordered](it iter.Seq[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if a < b {
			return a
		} else {
			return b
		}
	})
}

// Return the right element.
func Last[T any](it iter.Seq[T]) (T, bool) {
	var result T
	var ok bool
	for v := range it {
		if !ok {
			ok = true
		}
		result = v
	}
	return result, ok
}

// Return the element at index.
func At[T any](it iter.Seq[T], index int) (T, bool) {
	var zero T
	var ok bool
	var i int
	for v := range it {
		if i == index {
			return v, true
		}
	}
	return zero, ok
}

// Splitting an iterator whose elements are pair into two lists.
func Unzip[A any, B any](it iter.Seq[types.Pair[A, B]]) ([]A, []B) {
	var arrA = make([]A, 0)
	var arrB = make([]B, 0)
	for p := range it {
		arrA = append(arrA, p.First)
		arrB = append(arrB, p.Second)
	}
	return arrA, arrB
}

// to built-in map.
func ToMap[K comparable, V any](it iter.Seq[types.Pair[K, V]]) map[K]V {
	var r = make(map[K]V)
	for p := range it {
		r[p.First] = p.Second
	}
	return r
}

// Collecting via Collector.
func Collect[T any, S any, R any](it iter.Seq[T], collector interfaces.Collector[S, T, R]) R {
	var s = collector.Builder()
	for v := range it {
		collector.Append(s, v)
	}
	return collector.Finish(s)
}
