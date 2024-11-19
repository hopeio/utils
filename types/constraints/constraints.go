package constraints

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Callback[T any] interface {
	~func() | ~func() error | ~func(T) | ~func(T) error
}

type Rangeable constraints.Ordered

type Key interface {
	constraints.Integer | ~string | ~[8]byte | ~[16]byte | ~[32]byte | constraints.Float //| ~[]byte
}

type ID = Key

type Basic interface {
	Number | ~bool
}

type Ordered constraints.Ordered
