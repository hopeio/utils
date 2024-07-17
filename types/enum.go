package types

import "golang.org/x/exp/constraints"

type Enum[T constraints.Unsigned | ~string] struct {
	Value T
}
