package slices

func Every[T any, S ~[]T](slice S, fn func(T) bool) bool {
	for _, t := range slice {
		if !fn(t) {
			return false
		}
	}
	return true
}

func Some[T any, S ~[]T](slice S, fn func(T) bool) bool {
	for _, t := range slice {
		if fn(t) {
			return true
		}
	}
	return false
}

func Zip[T any, S ~[]T](s1, s2 S) [][2]T {
	var newSlice [][2]T
	for i := range s1 {
		newSlice = append(newSlice, [2]T{s1[i], s2[i]})
	}
	return newSlice
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
