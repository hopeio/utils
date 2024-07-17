package arraylist

type ArrayList[T any] []T

func New[T any](capacity int) ArrayList[T] {
	return make([]T, 0, capacity)
}

func (a *ArrayList[T]) Push(v T) {
	*a = append(*a, v)
}

func (a *ArrayList[T]) Pop() (T, bool) {
	l := *a
	if len(l) == 0 {
		return *new(T), false
	}
	v := l[0]
	*a = l[1:]
	return v, true
}

func (l ArrayList[T]) First() (T, bool) {
	if len(l) == 0 {
		return *new(T), false
	}
	return l[0], true
}

func (l ArrayList[T]) Last() (T, bool) {
	if len(l) == 0 {
		return *new(T), false
	}
	return l[len(l)-1], true
}

func (l ArrayList[T]) Len() int {
	return len(l)
}
