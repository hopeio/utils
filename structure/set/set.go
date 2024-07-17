package set

type Set[Key comparable] map[Key]struct{}

func New[Key comparable]() Set[Key] {
	return make(Set[Key])
}

func (s Set[Key]) Add(key Key) {
	s[key] = struct{}{}
}

func (s Set[Key]) ToSlice() []Key {
	arr := make([]Key, 0, len(s))
	for k := range s {
		arr = append(arr, k)
	}
	return arr
}
