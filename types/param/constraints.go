package param

import (
	"golang.org/x/exp/constraints"
	"time"
)

type Rangeable interface {
	constraints.Ordered | time.Time | ~*time.Time | ~string
}

type Ordered interface {
	constraints.Ordered | time.Time
}
