package constraints

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Callback[T any] interface {
	~func() | ~func() error | ~func(T) | ~func(T) error
}

type Range interface {
	constraints.Ordered | time.Time | ~*time.Time | ~string
}

type Key interface {
	constraints.Integer | ~string | ~[8]byte | ~[16]byte | ~[32]byte | constraints.Float //| ~[]byte
}

type ID = Key

type Basic interface {
	Number | ~bool
}

type Ordered interface {
	constraints.Ordered | time.Time
}
