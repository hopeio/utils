package slices

import (
	"fmt"
	"github.com/hopeio/utils/cmp"
	reflecti "github.com/hopeio/utils/reflect"
	"golang.org/x/exp/constraints"

	"reflect"
	"strings"
	"unsafe"
)

func Contains[S ~[]T, T comparable](s S, v T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return true
		}
	}
	return false
}

func ContainsByKey[S ~[]cmp.EqualKey[K], K comparable](s S, v K) bool {
	for i := 0; i < len(s); i++ {
		if s[i].EqualKey() == v {
			return true
		}
	}
	return false
}

func In[S ~[]T, T comparable](v T, s S) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}

func InByKey[S ~[]E, E cmp.EqualKey[K], K comparable](key K, s S) bool {
	for _, x := range s {
		if x.EqualKey() == key {
			return true
		}
	}
	return false
}

func Index[S ~[]T, T comparable](v T, s S) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func Reverse[S ~[]T, T any](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func Max[S ~[]T, T constraints.Ordered](s S) T {
	if len(s) == 0 {
		return *new(T)
	}
	max := s[0]
	if len(s) == 1 {
		return max
	}
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}

	return max
}

func Min[S ~[]T, T constraints.Ordered](s S) T {
	if len(s) == 0 {
		return *new(T)
	}
	min := s[0]
	if len(s) == 1 {
		return min
	}
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}

	return min
}

// 将切片转换为map
func ToMap[S ~[]T, T any, K comparable, V any](s S, getKV func(T) (K, V)) map[K]V {
	m := make(map[K]V)
	for _, s := range s {
		k, v := getKV(s)
		m[k] = v
	}
	return m
}

// 将切片按照某个key分类
func Classify[S ~[]T, T any, K comparable, V any](s S, getKV func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for _, s := range s {
		k, v := getKV(s)
		if ms, ok := m[k]; ok {
			m[k] = append(ms, v)
		} else {
			m[k] = []V{v}
		}

	}
	return m
}

func Swap[S ~[]T, T any](s S, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func ForEach[S ~[]T, T any](s S, handle func(idx int, v T)) {
	for i, t := range s {
		handle(i, t)
	}
}

func ForEachValue[S ~[]T, T any](s S, handle func(v T)) {
	for _, v := range s {
		handle(v)
	}
}

// 遍历切片,参数为下标，利用闭包实现遍历
func ForEachIndex[S ~[]T, T any](s S, handle func(i int)) {
	for i := range s {
		handle(i)
	}
}

func JoinByIndex[S ~[]T, T any](s S, toString func(i int) string, sep string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return toString(0)
	}
	n := len(sep) * (len(s) - 1)
	for i := 0; i < len(s); i++ {
		n += len(toString(i))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(toString(0))
	for i := range s[1:] {
		b.WriteString(sep)
		b.WriteString(toString(i))
	}
	return b.String()
}

func JoinByValue[S ~[]T, T any](s S, toString func(v T) string, sep string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return toString(s[0])
	}
	n := len(sep) * (len(s) - 1)
	for i := 0; i < len(s); i++ {
		n += len(toString(s[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(toString(s[0]))
	for _, s := range s[1:] {
		b.WriteString(sep)
		b.WriteString(toString(s))
	}
	return b.String()
}

func Join[S ~[]T, T fmt.Stringer](s S, sep string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s[0].String()
	}
	n := len(sep) * (len(s) - 1)
	for i := 0; i < len(s); i++ {
		n += len(s[i].String())
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(s[0].String())
	for _, s := range s[1:] {
		b.WriteString(sep)
		b.WriteString(s.String())
	}
	return b.String()
}

func ReverseForEach[S ~[]T, T any](s S, handle func(idx int, v T)) {
	l := len(s)
	for i := l - 1; i > 0; i-- {
		handle(i, s[i])
	}
}

func Map[T1, T2 any, T1S ~[]T1](s T1S, fn func(T1) T2) []T2 {
	ret := make([]T2, 0, len(s))
	for _, s := range s {
		ret = append(ret, fn(s))
	}
	return ret
}

func Map2[T1, T2 any, T1S ~[]T1, T2S ~[]T2](s T1S, fn func(T1) T2) T2S {
	ret := make([]T2, 0, len(s))
	for _, s := range s {
		ret = append(ret, fn(s))
	}
	return ret
}

func Filter[S ~[]T, T any](fn func(T) bool, src S) S {
	var dst []T
	for _, v := range src {
		if fn(v) {
			dst = append(dst, v)
		}
	}
	return dst
}

func Reduce[S ~[]T, T any](slices S, fn func(T, T) T) T {
	ret := fn(slices[0], slices[1])
	for i := 2; i < len(slices); i++ {
		ret = fn(ret, slices[i])
	}
	return ret
}

func Cast[T1, T2 any, T1S ~[]T1](s T1S) []T2 {
	t1, t2 := new(T1), new(T2)
	t1type, t2type := reflect.TypeOf(t1).Elem(), reflect.TypeOf(t2).Elem()
	t1kind, t2kind := t1type.Kind(), t2type.Kind()

	if t1type.ConvertibleTo(t2type) && reflecti.CanCast(t1type, t2type, false) {
		if t1kind == t2kind {
			return unsafe.Slice((*T2)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
		}
		if t1kind != reflect.Interface && t2kind != reflect.Interface {
			return Map(s, func(v T1) T2 { return *(*T2)(unsafe.Pointer(&v)) })
		}
	}

	if _, ok := any(t1).(T2); ok {
		return Map(s, func(v T1) T2 { return any(v).(T2) })
	}
	if _, ok := any(t2).(T1); ok {
		return Map(s, func(v T1) T2 { return any(v).(T2) })
	}
	panic("unsupported type")
}

func Cast2[T1, T2 any, T1S ~[]T1, T2S ~[]T2](s T1S) T2S {
	return (T2S)(Cast[T1, T2](s))
}

func GuardSlice(buf *[]byte, n int) {
	c := cap(*buf)
	l := len(*buf)
	if c-l < n {
		c = c>>1 + n + l
		if c < 32 {
			c = 32
		}
		tmp := make([]byte, l, c)
		copy(tmp, *buf)
		*buf = tmp
	}
}

//go:nosplit
func Ptr2SlicePtr(s unsafe.Pointer, l int, c int) unsafe.Pointer {
	slice := &reflecti.Slice{
		Ptr: s,
		Len: l,
		Cap: c,
	}
	return unsafe.Pointer(slice)
}

func FilterPlace[S ~[]T, T any](slices S, fn func(T) bool) S {
	n := len(slices) - 1
	for i := 0; i <= n; {
		if fn(slices[i]) {
			if i < n {
				Swap(slices, i, n)
			}
			n--
			continue
		}
		i++
	}
	return slices[:n+1]
}

func Remove[S ~[]T, T any](slices S, i int) S {
	return append(slices[:i], slices[i+1:]...)
}

type Collector[T any] struct {
}

func (c Collector[T]) Builder() *[]T {
	s := make([]T, 0)
	return &s
}
func (c Collector[T]) Append(builder *[]T, element T) {
	*builder = append(*builder, element)
}
func (c Collector[T]) Finish(builder *[]T) []T {
	return *builder
}

func TwoDimensionalSlice[S ~[][]T, T any](s S, rowStart, rowEnd, colStart, colEnd int) S {
	ret := make([][]T, rowEnd-rowStart)
	for i := range ret {
		ret[i] = s[rowStart+i][colStart:colEnd]
	}
	return ret
}

func ThreeDimensionalSlice[S ~[][][]T, T any](s S, rowStart, rowEnd, colStart, colEnd, sliceStart, sliceEnd int) S {
	ret := make([][][]T, rowEnd-rowStart)
	for i := range ret {
		ret[i] = make([][]T, colStart-colEnd)
		for j := range ret[i] {
			ret[i][j] = s[rowStart+i][colStart+j][sliceStart:sliceEnd]
		}
	}
	return ret
}
