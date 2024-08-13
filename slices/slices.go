package slices

type Slice[T any] []T

func (slice Slice[T]) Len() int { return len(slice) }

func (slice Slice[T]) ForEach(fn func(T)) {
	for _, t := range slice {
		fn(t)
	}
}

func (slice Slice[T]) Filter(fn func(T) bool) []T {
	var newSlices []T
	for _, t := range slice {
		if fn(t) {
			newSlices = append(newSlices, t)
		}
	}
	return newSlices
}

func (slice Slice[T]) Every(fn func(T) bool) bool {
	for _, t := range slice {
		if !fn(t) {
			return false
		}
	}
	return true
}

func (slice Slice[T]) Some(fn func(T) bool) bool {
	for _, t := range slice {
		if fn(t) {
			return true
		}
	}
	return false
}

func (slice Slice[T]) Zip(s []T) [][2]T {
	var newSlice [][2]T
	for i := range slice {
		newSlice = append(newSlice, [2]T{slice[i], s[i]})
	}
	return newSlice
}

func (slice Slice[T]) Reduce(fn func(T, T) T) T {
	ret := fn(slice[0], slice[1])
	for i := 2; i < len(slice); i++ {
		ret = fn(ret, slice[i])
	}
	return ret
}

type MapSlice[T, V any] Slice[T]

func (slice MapSlice[T, V]) Map(fn func(T) V) []V {
	ret := make([]V, 0, len(slice))
	for _, s := range slice {
		ret = append(ret, fn(s))
	}
	return ret
}

type ComparableSlice[T comparable] []T

// 去重
func (slices ComparableSlice[T]) Deduplicate() ComparableSlice[T] {
	if len(slices) < SmallArrayLen {
		newslices := make(ComparableSlice[T], 0, 2)
		for i := 0; i < len(slices); i++ {
			if !In(slices[i], newslices) {
				newslices = append(newslices, slices[i])
			}
		}
		return newslices
	}
	set := make(map[T]struct{})
	for i := 0; i < len(slices); i++ {
		set[slices[i]] = struct{}{}
	}
	newslices := make(ComparableSlice[T], 0, len(slices))
	for k := range set {
		newslices = append(newslices, k)
	}
	return newslices
}

// 去重
func Deduplicate[S ~[]T, T comparable](slice S) S {
	if len(slice) < SmallArrayLen {
		newslice := make(S, 0, 2)
		for i := 0; i < len(slice); i++ {
			if !In(slice[i], newslice) {
				newslice = append(newslice, slice[i])
			}
		}
		return newslice
	}
	set := make(map[T]struct{})
	for i := 0; i < len(slice); i++ {
		set[slice[i]] = struct{}{}
	}
	newslice := make(S, 0, len(slice))
	for k := range set {
		newslice = append(newslice, k)
	}
	return newslice
}
