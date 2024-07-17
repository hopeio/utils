package pull

import (
	"github.com/hopeio/utils/cmp"
	"github.com/hopeio/utils/types"
	"golang.org/x/exp/constraints"
)

// Ruturns true if the count of Iterator is 0.
func IsEmpty[T any](it Iterator[T]) bool {
	_, ok := it.Next()
	return ok == false
}

// Ruturns true if the count of Iterator is 0.
func IsNotEmpty[T any](it Iterator[T]) bool {
	_, ok := it.Next()
	return ok == true
}

// Converts a Iterator to a Slice.
func ToSlice[T any](it Iterator[T]) []T {
	var arr = make([]T, 0)
	ForEach(it, func(t T) {
		arr = append(arr, t)
	})
	return arr
}

// Returns true if the target is included in the iterator.
func Contains[T comparable](it Iterator[T], target T) bool {
	for {
		if v, ok := it.Next(); ok {
			if v == target {
				return true
			}
		} else {
			break
		}
	}
	return false
}

// Returns the sum of all the elements in the iterator.
func Sum[T constraints.Integer | constraints.Float](it Iterator[T]) T {
	return Fold(it, 0, func(a, b T) T {
		return a + b
	})
}

// Returns the product of all the elements in the iterator.
func Product[T constraints.Integer | constraints.Float](it Iterator[T]) T {
	return Fold(it, 1, func(a, b T) T {
		return a * b
	})
}

// Returns the average of all the elements in the iterator.
func Average[T constraints.Integer | constraints.Float](it Iterator[T]) float64 {
	return Fold(Enumerate(it), 0.0, func(result float64, item *types.Pair[int, T]) float64 {
		return result + (float64(item.Second)-result)/float64(item.First+1)
	})
}

// Return the total number of iterators.
func Count[T any](it Iterator[T]) uint64 {
	return Fold(it, 0, func(v uint64, _ T) uint64 {
		return v + 1
	})
}

// Return the maximum value of all elements of the iterator.
func Max[T constraints.Ordered](it Iterator[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if a > b {
			return a
		} else {
			return b
		}
	})
}

// Return the maximum value of all elements of the iterator.
func MaxBy[T any](it Iterator[T], greater cmp.LessFunc[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if greater(a, b) {
			return a
		} else {
			return b
		}
	})
}

// Return the minimum value of all elements of the iterator.
func Min[T constraints.Ordered](it Iterator[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if a < b {
			return a
		} else {
			return b
		}
	})
}

// Return the minimum value of all elements of the iterator.
func MinBy[T any](it Iterator[T], less cmp.LessFunc[T]) (T, bool) {
	return Reduce(it, func(a T, b T) T {
		if less(a, b) {
			return a
		} else {
			return b
		}
	})
}

// The action is executed for each element of the iterator, and the argument to the action is the element.
func ForEach[T any](it Iterator[T], action func(T)) {
	for {
		if v, ok := it.Next(); ok {
			action(v)
		} else {
			break
		}
	}
}

// Returns true if all elements in the iterator match the condition.
func AllMatch[T any](it Iterator[T], predicate func(T) bool) bool {
	for {
		if v, ok := it.Next(); ok {
			if !predicate(v) {
				return false
			}
		} else {
			break
		}
	}
	return true
}

// Returns true if none elements in the iterator match the condition.
func NoneMatch[T any](it Iterator[T], predicate func(T) bool) bool {
	for {
		if v, ok := it.Next(); ok {
			if predicate(v) {
				return false
			}
		} else {
			break
		}
	}
	return true
}

// Returns true if any elements in the iterator match the condition.
func AnyMatch[T any](it Iterator[T], predicate func(T) bool) bool {
	for {
		if v, ok := it.Next(); ok {
			if predicate(v) {
				return true
			}
		} else {
			break
		}
	}
	return false
}

// Return the first element.
func First[T any](it Iterator[T]) (T, bool) {
	return it.Next()
}

// Return the last element.
func Last[T any](it Iterator[T]) (T, bool) {
	var result T
	var ok bool
	for {
		if v, ok := it.Next(); ok {
			result = v
		} else {
			break
		}
	}
	return result, ok
}

// Return the element at index.
func At[T any](it Iterator[T], index int) (T, bool) {
	var result, ok = it.Next()
	var i = 0
	for i < index && ok {
		result, ok = it.Next()
		i++
	}
	return result, ok
}

// Return the value of the final composite, operates on the iterator from front to back.
func Reduce[T any](it Iterator[T], operation func(T, T) T) (T, bool) {
	if v, ok := it.Next(); ok {
		return Fold(it, v, operation), true
	}
	return *new(T), false
}

// Return the value of the final composite, operates on the iterator from back to front.
func Fold[T any, R any](it Iterator[T], initial R, operation func(R, T) R) R {
	var result = initial
	for {
		if v, ok := it.Next(); ok {
			result = operation(result, v)
		} else {
			break
		}
	}
	return result
}

// Splitting an iterator whose elements are pair into two lists.
func Unzip[A any, B any](it Iterator[*types.Pair[A, B]]) ([]A, []B) {
	var arrA = make([]A, 0)
	var arrB = make([]B, 0)
	for {
		if v, ok := it.Next(); ok {
			arrA = append(arrA, v.First)
			arrB = append(arrB, v.Second)
		} else {
			break
		}
	}
	return arrA, arrB
}

// to built-in map.
func ToMap[K comparable, V any](it Iterator[*types.Pair[K, V]]) map[K]V {
	var r = make(map[K]V)
	for {
		if v, ok := it.Next(); ok {
			r[v.First] = v.Second
		} else {
			break
		}
	}
	return r
}

type Collector[S any, T any, R any] interface {
	Builder() S
	Append(builder S, element T)
	Finish(builder S) R
}

// Collecting via Collector.
func Collect[T any, S any, R any](it Iterator[T], collector Collector[S, T, R]) R {
	var s = collector.Builder()
	for {
		if v, ok := it.Next(); ok {
			collector.Append(s, v)
		} else {
			break
		}
	}
	return collector.Finish(s)
}
