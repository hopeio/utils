package node

import "golang.org/x/exp/constraints"

type Node[T any] struct {
	Value T
}

type ListNode[T any] struct {
	Next  *ListNode[T]
	Value T
}

type LinkNode[T any] struct {
	Prev, Next *LinkNode[T]
	Value      T
}

type ListKNode[K comparable, T any] struct {
	Next  *ListKNode[K, T]
	Key   K
	Value T
}

type LinkKNode[K comparable, T any] struct {
	Prev, Next *LinkKNode[K, T]
	Key        K
	Value      T
}

type ListOrdKNode[K constraints.Ordered, T any] struct {
	Next  *LinkKNode[K, T]
	Key   K
	Value T
}

type LinkOrdKNode[K comparable, T any] struct {
	Prev, Next *LinkOrdKNode[K, T]
	Key        K
	Value      T
}
