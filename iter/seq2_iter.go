package iter

import (
	"github.com/hopeio/utils/types"
	"github.com/hopeio/utils/types/constraints"
	"github.com/hopeio/utils/types/funcs"
	"iter"
	"sort"
)

// Dont'use please use types.Pair And Seq
func Filter2[K, V any](seq iter.Seq2[K, V], test funcs.PredicateKV[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if test(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

func Map2[K, V, R any](seq iter.Seq2[K, V], f funcs.UnaryKVFunction[K, V, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func Map3[K, V, RK, RV any](seq iter.Seq2[K, V], f funcs.UnaryKVFunction2[K, V, RK, RV]) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
	}
}

func FlatMap2[K, V, R any](seq iter.Seq2[K, V], flatten funcs.UnaryKVFunction[K, V, iter.Seq[R]]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for k, v := range seq {
			for v2 := range flatten(k, v) {
				if !yield(v2) {
					return
				}
			}
		}
	}
}

func FlatMap3[K, V, RK, RV any](seq iter.Seq2[K, V], flatten funcs.UnaryKVFunction[K, V, iter.Seq2[RK, RV]]) iter.Seq2[RK, RV] {
	return func(yield func(RK, RV) bool) {
		for k, v := range seq {
			for k2, v2 := range flatten(k, v) {
				if !yield(k2, v2) {
					return
				}
			}
		}
	}
}

func Peek2[K, V any](seq iter.Seq2[K, V], accept funcs.ConsumerKV[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			accept(k, v)
			if !yield(k, v) {
				return
			}
		}
	}
}

func Distinct2[K, V any, C comparable](seq iter.Seq2[K, V], f funcs.UnaryKVFunction[K, V, C]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var set = make(map[C]struct{})
		for k, v := range seq {
			c := f(k, v)
			_, ok := set[c]
			if !ok {
				if !yield(k, v) {
					return
				}
				set[c] = struct{}{}
			}
		}
	}
}

func Sorted2[K, V any](it iter.Seq2[K, V], cmp funcs.LessKV[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		keys, vals := Collect2(it)
		sort.SliceStable(vals, func(i, j int) bool {
			return cmp(keys[i], vals[i], keys[j], vals[j])
		})
		for i := range vals {
			if !yield(keys[i], vals[i]) {
				return
			}
		}
	}
}

func IsSorted2[K, V any](seq iter.Seq2[K, V], cmp funcs.LessKV[K, V]) bool {
	var lastKey K
	var lastVal V
	check := func(currKey K, currVal V) bool {
		if !cmp(lastKey, lastVal, currKey, currVal) {
			return false
		}
		lastKey = currKey
		lastVal = currVal
		return true
	}

	var has bool
	for k, v := range seq {
		if !has {
			lastKey = k
			lastVal = v
			has = true
		} else {
			if !check(k, v) {
				return false
			}
		}
	}
	return true
}

func Limit2[K, V any, Number constraints.Number](seq iter.Seq2[K, V], limit Number) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			limit--
			if limit < 0 {
				return
			}
			if !yield(k, v) {
				return
			}
		}
	}
}

func Skip2[K, V any, Number constraints.Number](seq iter.Seq2[K, V], skip Number) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			skip--
			if skip < 0 {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

func ForEach2[K, V any](seq iter.Seq2[K, V], accept funcs.ConsumerKV[K, V]) {
	for k, v := range seq {
		accept(k, v)
	}
}

func Collect2[K, V any](seq iter.Seq2[K, V]) (ks []K, vs []V) {
	for k, v := range seq {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return
}

func AllMatch2[K, V any](seq iter.Seq2[K, V], test funcs.PredicateKV[K, V]) bool {
	for k, v := range seq {
		if !test(k, v) {
			return false
		}
	}
	return true
}

func NoneMatch2[K, V any](seq iter.Seq2[K, V], test funcs.PredicateKV[K, V]) bool {
	for k, v := range seq {
		if test(k, v) {
			return false
		}
	}
	return true
}

func AnyMatch2[K, V any](seq iter.Seq2[K, V], test funcs.PredicateKV[K, V]) bool {
	for k, v := range seq {
		if test(k, v) {
			return true
		}
	}
	return false
}

func Reduce2[K, V any](seq iter.Seq2[K, V], acc funcs.BinaryKVOperator[K, V]) (K, V, bool) {
	var resultKey K
	var resultVal V
	var has bool
	for k, v := range seq {
		if !has {
			resultKey = k
			resultVal = v
			has = true
		} else {
			resultKey, resultVal = acc(resultKey, resultVal, k, v)
		}
	}
	if has {
		return resultKey, resultVal, has
	}
	return resultKey, resultVal, has
}

func Fold2[K, V, R any](seq iter.Seq2[K, V], initVal R, acc funcs.BinaryKVFunction[R, K, V, R]) (result R) {
	result = initVal
	for k, v := range seq {
		result = acc(result, k, v)
	}
	return result
}

func Fold3[K, V, RK, RV any](seq iter.Seq2[K, V], initKey RK, initVal RV, acc funcs.BinaryKVFunction2[RK, RV, K, V, RK, RV]) (resultKey RK, resultVal RV) {
	resultKey, resultVal = initKey, initVal
	for k, v := range seq {
		resultKey, resultVal = acc(resultKey, resultVal, k, v)
	}
	return
}

func Count2[K, V any](seq iter.Seq2[K, V]) (count int64) {
	for _, _ = range seq {
		count++
	}
	return
}

func Enumerate2[K, V any](seq iter.Seq2[K, V]) iter.Seq2[int, types.Pair[K, V]] {
	return func(yield func(int, types.Pair[K, V]) bool) {
		var count int
		for k, v := range seq {
			if !yield(count, types.PairOf(k, v)) {
				return
			}
			count++
		}
	}
}

func Zip2[K, V any](seq1 iter.Seq2[K, V], seq2 iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var firstFinished bool
		if !firstFinished {
			for k1, v1 := range seq1 {
				if !yield(k1, v1) {
					return
				}
			}
			firstFinished = true
		}
		for k2, v2 := range seq2 {
			if !yield(k2, v2) {
				return
			}
		}
	}
}
