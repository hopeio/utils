package types

type Dict[K comparable, V any] struct {
	Key   K
	Value V
}
