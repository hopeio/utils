package iter

import (
	"github.com/hopeio/utils/types"
	"golang.org/x/exp/constraints"
	"iter"
	"sort"
)

// Filter keep elements which satisfy the Predicate.
// 保留满足断言的元素
func Filter[T any](seq iter.Seq[T], test Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if test(v) && !yield(v) {
				return
			}
		}
	}
}

// Map transform the element use Fuction.
// 使用输入函数对每个元素进行转换
func Map[T, R any](seq iter.Seq[T], f Function[T, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// FlatMap transform each element in Seq[T] to a new Seq[R].
// 将原本序列中的每个元素都转换为一个新的序列，
// 并将所有转换后的序列依次连接起来生成一个新的序列
func FlatMap[T, R any](seq iter.Seq[T], flatten Function[T, iter.Seq[R]]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v := range seq {
			for v2 := range flatten(v) {
				if !yield(v2) {
					return
				}
			}
		}
	}
}

// Peek visit every element in the Seq and leave them on the Seq.
// 访问序列中的每个元素而不消费它
func Peek[T any](seq iter.Seq[T], accept Consumer[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			accept(v)
			if !yield(v) {
				return
			}
		}
	}
}

// Distinct remove duplicate elements.
// 对序列中的元素去重
func Distinct[T any, C comparable](seq iter.Seq[T], f Function[T, C]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var set = make(map[C]struct{})
		for v := range seq {
			k := f(v)
			_, ok := set[k]
			if !ok {
				if !yield(v) {
					return
				}
				set[k] = struct{}{}
			}
		}
	}
}

// Sorted sort elements in the Seq by Comparator.
// 对序列中的元素排序
func Sorted[T any](it iter.Seq[T], cmp Comparator[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		vals := Collect(it)
		sort.SliceStable(vals, func(i, j int) bool {
			return cmp(vals[i], vals[j])
		})
		for _, v := range vals {
			if !yield(v) {
				return
			}
		}
	}
}

// IsSorted
// 对序列中的元素是否排序
func IsSorted[T any](seq iter.Seq[T], cmp Comparator[T]) bool {
	var last T
	check := func(curr T) bool {
		if !cmp(last, curr) {
			return false
		}
		last = curr
		return true
	}

	var has bool
	for v := range seq {
		if !has {
			last = v
			has = true
		} else {
			if !check(v) {
				return false
			}
		}
	}
	return true
}

// Limit limits the number of elements in Seq.
// 限制元素个数
func Limit[T any](seq iter.Seq[T], limit int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			limit--
			if limit < 0 {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// Skip drop some elements of the Seq.
// 跳过指定个数的元素
func Skip[T any](seq iter.Seq[T], skip int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			skip--
			if skip < 0 {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func UntilComparable[T comparable](seq iter.Seq[T], e T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if v == e {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Until[T any](seq iter.Seq[T], match Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range seq {
			if match(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// ForEach consume every elements in the Seq.
// 消费序列中的每个元素
func ForEach[T any](seq iter.Seq[T], accept Consumer[T]) {
	for v := range seq {
		accept(v)
	}
}

// Collect return all elements as a slice.
// 将序列中所有元素收集为切片返回
func Collect[T any](seq iter.Seq[T]) (result []T) {
	for v := range seq {
		result = append(result, v)
	}
	return
}

// AllMatch test if every elements are all match the Predicate.
// 是否每个元素都满足条件
func AllMatch[T any](seq iter.Seq[T], test Predicate[T]) bool {
	for v := range seq {
		if !test(v) {
			return false
		}
	}
	return true
}

// NoneMatch test if none element matches the Predicate.
// 是否没有元素满足条件
func NoneMatch[T any](seq iter.Seq[T], test Predicate[T]) bool {
	for v := range seq {
		if test(v) {
			return false
		}
	}
	return true
}

// AnyMatch test if any element matches the Predicate.
// 是否有任意元素满足条件
func AnyMatch[T any](seq iter.Seq[T], test Predicate[T]) bool {
	for v := range seq {
		if test(v) {
			return true
		}
	}
	return false
}

// Reduce accumulate each element using the binary operation.
// 使用给定的累加函数, 累加序列中的每个元素.
// 序列中可能没有元素因此返回的是 Optional
func Reduce[T any](seq iter.Seq[T], acc BinaryOperator[T]) (T, bool) {
	var result T
	var has bool
	for v := range seq {
		if !has {
			result = v
			has = true
		} else {
			result = acc(result, v)
		}
	}
	if has {
		return result, has
	}
	return result, has
}

// Fold accumulate each element using the BiFunction
// starting from the initial value.
// 从初始值开始, 通过 acc 函数累加每个元素
func Fold[T, R any](seq iter.Seq[T], initVal R, acc BiFunction[R, T, R]) (result R) {
	result = initVal
	for v := range seq {
		result = acc(result, v)
	}
	return result
}

// First find the first element in the Seq.
// 返回序列中的第一个元素(如有).
func First[T any](seq iter.Seq[T]) (T, bool) {
	for v := range seq {
		return v, true
	}
	return *new(T), false
}

// Count return the count of elements in the Seq.
// 返回序列中的元素个数
func Count[T any](seq iter.Seq[T]) (count int) {
	for _ = range seq {
		count++
	}
	return
}

func Enumerate[T any](seq iter.Seq[T]) iter.Seq[types.Pair[int, T]] {
	return func(yield func(types.Pair[int, T]) bool) {
		var count int
		for v := range seq {
			if !yield(types.PairOf(count, v)) {
				return
			}
			count++
		}
	}
}

func Zip[T any](seq1 iter.Seq[T], seq2 iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var firstFinished bool
		if !firstFinished {
			for v1 := range seq1 {
				if !yield(v1) {
					return
				}
			}
			firstFinished = true
		}
		for v2 := range seq2 {
			if !yield(v2) {
				return
			}
		}
	}
}

func Sum[T any](seq iter.Seq[T], add BinaryOperator[T]) T {
	var result T
	var idx int
	for v := range seq {
		if idx == 0 {
			result = v
			continue
		}
		result = add(result, v)
		idx++
	}
	return result
}

func SumComparable[T constraints.Ordered](seq iter.Seq[T]) T {
	var result T
	for v := range seq {
		result += v
	}
	return result
}
