package pull

import "github.com/hopeio/utils/types"

// Supplier 产生一个元素
type Supplier[T any] func() T

// Function 将一个类型转为另一个类型
type Function[T, R any] func(T) R

// Predicate 断言是否满足指定条件
type Predicate[T any] func(T) bool

// UnaryOperator 对输入进行一元运算返回相同类型的结果
type UnaryOperator[T any] func(T) T

// BiFunction 将两个类型转为第三个类型
type BiFunction[T, R, U any] func(T, R) U

// BinaryOperator 输入两个相同类型的参数，对其做二元运算，返回相同类型的结果
type BinaryOperator[T any] func(T, T) T

// Comparator 比较两个元素.
// 第一个元素大于第二个元素时，返回正数;
// 第一个元素小于第二个元素时，返回负数;
// 否则返回 0.
type Comparator[T any] func(T, T) int

// Consumer 消费一个元素
type Consumer[T any] func(T)

type Stream[T any] interface {
	Filter(Predicate[T]) Stream[T]
	Map(Function[T, T]) Stream[T]               //同类型转换,没啥意义
	FlatMap(Function[T, Iterator[T]]) Stream[T] //同Map
	Peek(Consumer[T]) Stream[T]
	Fold(initVal T, acc BinaryOperator[T])
	Zip(Iterator[T], Iterator[T]) Stream[T]

	Distinct(Function[T, int]) Stream[T]
	Sorted(Comparator[T]) Stream[T]
	Limit(int64) Stream[T]
	Skip(int64) Stream[T]

	ForEach(Consumer[T])
	Collect() []T
	All(Predicate[T]) bool
	None(Predicate[T]) bool
	Any(Predicate[T]) bool
	Reduce(acc BinaryOperator[T]) (T, bool)
	ReduceFrom(initVal T, acc BinaryOperator[T]) T
	First() (T, bool)
	Count() int64
}

// Add subscripts to the incoming iterators.
func Enumerate[T any](it Iterator[T]) Iterator[*types.Pair[int, T]] {
	return &enumerateStream[T]{-1, it}
}

type enumerateStream[T any] struct {
	index    int
	iterator Iterator[T]
}

func (a *enumerateStream[T]) Next() (*types.Pair[int, T], bool) {
	if v, ok := a.iterator.Next(); ok {
		a.index++
		return &types.Pair[int, T]{First: a.index, Second: v}, ok
	}
	return nil, false
}

// Use transform to map an iterator to another iterator.
func Map[T any, R any](it Iterator[T], transform func(T) R) Iterator[R] {
	return &mapStream[T, R]{transform, it}
}

type mapStream[T any, R any] struct {
	transform func(T) R
	iterator  Iterator[T]
}

func (a *mapStream[T, R]) Next() (R, bool) {
	if v, ok := a.iterator.Next(); ok {
		return a.transform(v), ok
	}
	return *new(R), false
}

// Use predicate to filter an iterator to another iterator。
func Filter[T any](it Iterator[T], predicate func(T) bool) Iterator[T] {
	return &filterStream[T]{predicate, it}
}

type filterStream[T any] struct {
	predicate func(T) bool
	iterator  Iterator[T]
}

func (a *filterStream[T]) Next() (T, bool) {
	for {
		if v, ok := a.iterator.Next(); ok {
			if a.predicate(v) {
				return v, ok
			}
		} else {
			break
		}
	}
	return *new(T), false
}

// Convert an iterator to another iterator that limits the maximum number of iterations.
func Limit[T any](it Iterator[T], count int) Iterator[T] {
	return &limitStream[T]{count, it}
}

type limitStream[T any] struct {
	limit    int
	iterator Iterator[T]
}

func (a *limitStream[T]) Next() (T, bool) {
	if a.limit != 0 {
		a.limit -= 1
		return a.iterator.Next()
	}
	return *new(T), false
}

// Converts an iterator to another iterator that skips a specified number of times.
func Skip[T any](it Iterator[T], count int) Iterator[T] {
	return &skipStream[T]{count, it}
}

type skipStream[T any] struct {
	skip     int
	iterator Iterator[T]
}

func (a *skipStream[T]) Next() (T, bool) {
	for a.skip > 0 {
		if v, ok := a.iterator.Next(); !ok {
			return v, ok
		}
		a.skip -= 1
	}
	return a.iterator.Next()
}

// Converts an iterator to another iterator that skips a specified number of times each time.
func Step[T any](it Iterator[T], count int) Iterator[T] {
	return &stepStream[T]{count - 1, true, it}
}

type stepStream[T any] struct {
	step      int
	firstTake bool
	iterator  Iterator[T]
}

func (a *stepStream[T]) Next() (T, bool) {
	if a.firstTake {
		a.firstTake = false
		return a.iterator.Next()
	} else {
		return At(a.iterator, a.step)
	}
}

// By connecting two iterators in series,
// the new iterator will iterate over the first iterator before continuing with the second iterator.
func Concat[T any](left Iterator[T], right Iterator[T]) Iterator[T] {
	return &concatStream[T]{false, left, right}
}

type concatStream[T any] struct {
	firstNotFinished bool
	first            Iterator[T]
	last             Iterator[T]
}

func (a *concatStream[T]) Next() (T, bool) {
	if a.firstNotFinished {
		if v, ok := a.first.Next(); ok {
			return v, true
		}
		a.firstNotFinished = false
		return a.Next()
	}
	return a.last.Next()
}

// Converting a nested iterator to a flat iterator.
func Flatten[T Iterable[U], U any](it Iterator[T]) Iterator[U] {
	return &flattenStream[T, U]{iterator: it, subIter: nil, subIterOk: false}
}

type flattenStream[T Iterable[U], U any] struct {
	iterator  Iterator[T]
	subIter   Iterator[U]
	subIterOk bool
}

func (a *flattenStream[T, U]) Next() (U, bool) {
	if a.subIterOk {
		if item, ok := a.subIter.Next(); ok {
			return item, false
		} else {
			a.subIter, a.subIterOk = nil, false
			return a.Next()
		}
	} else if nextIter, ok := a.iterator.Next(); ok {
		a.subIter, a.subIterOk = nextIter.Iter(), true
		return a.Next()
	} else {
		return *new(U), false
	}
}

// Compress two iterators into one iterator. The length is the length of the shortest iterator.
func Zip[T any, U any](left Iterator[T], right Iterator[U]) Iterator[*types.Pair[T, U]] {
	return &zipStream[T, U]{left, right}
}

type zipStream[T any, U any] struct {
	first Iterator[T]
	last  Iterator[U]
}

func (a *zipStream[T, U]) Next() (types.Pair[T, U], bool) {
	if v1, ok1 := a.first.Next(); ok1 {
		if v2, ok2 := a.last.Next(); ok2 {
			return types.PairOf(v1, v2), true
		}
	}
	return types.Pair[T, U]{}, false
}

type IterStream[T any] struct {
	iter Iterator[T]
}

func (it *IterStream[T]) Map(transform func(T) T) *IterStream[T] {
	it.iter = Map(it.iter, transform)
	return it
}

func (it *IterStream[T]) Filter(f func(T) bool) *IterStream[T] {
	it.iter = Filter(it.iter, f)
	return it
}

func (it *IterStream[T]) Count() uint64 {
	return Count(it.iter)
}

func (it *IterStream[T]) ForEach(f func(T)) {
	ForEach(it.iter, f)
}

func (it *IterStream[T]) All(f func(T) bool) bool {
	return AllMatch(it.iter, f)
}

func (it *IterStream[T]) Any(f func(T) bool) bool {
	return AnyMatch(it.iter, f)
}

func (it *IterStream[T]) None(f func(T) bool) bool {
	return NoneMatch(it.iter, f)
}

func (it *IterStream[T]) Skip(n int) *IterStream[T] {
	it.iter = Skip(it.iter, n)
	return it
}

func (it *IterStream[T]) Limit(n int) *IterStream[T] {
	it.iter = Limit(it.iter, n)
	return it
}

func (it *IterStream[T]) Reduce(operation func(T, T) T) (T, bool) {
	return Reduce(it.iter, operation)
}

func (it *IterStream[T]) Fold(initVal T, operation func(T, T) T) T {
	return Fold(it.iter, initVal, operation)
}

func SliceStreamOf[S ~[]T, T any](source S) *IterStream[T] {
	return &IterStream[T]{iter: Slice(source)}
}

func StreamOf[T any](iter Iterator[T]) *IterStream[T] {
	return &IterStream[T]{iter: iter}
}
