package interfaces

import (
	"github.com/hopeio/utils/types/constraints"
	"time"
)

type IdGenerator[T constraints.ID] interface {
	Id() T
}

type DurationGenerator interface {
	Duration() time.Duration
}
