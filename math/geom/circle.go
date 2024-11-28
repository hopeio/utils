package geom

import "golang.org/x/exp/constraints"

type Circle struct {
	X        float64
	Y        float64
	Diameter float64
}

func (c *Circle) Bounds() *Bounds {
	return NewBounds(c.X, c.Y, c.Diameter, c.Diameter)
}

type CircleInt[T constraints.Integer] struct {
	X        T
	Y        T
	Diameter T
}
