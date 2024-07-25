package set

import "github.com/hopeio/utils/maps"

type Set[K comparable] map[K]struct{}

func New[K comparable]() Set[K] {
	return make(Set[K])
}

func (s Set[K]) Add(key K) {
	s[key] = struct{}{}
}

func (s Set[K]) Contains(key K) bool {
	_, ok := s[key]
	return ok
}

func (s Set[K]) Remove(key K) {
	delete(s, key)
}

func (s Set[K]) ToSlice() []K {
	return maps.Keys(s)
}
