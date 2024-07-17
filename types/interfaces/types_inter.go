package interfaces

import "github.com/hopeio/utils/types/constraints"

type Key[T constraints.Key] interface {
	Key() T
}
