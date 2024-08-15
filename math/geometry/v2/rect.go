package geometry

import constraintsi "github.com/hopeio/utils/types/constraints"

type Rectangle[T constraintsi.Number] struct {
	Min, Max Point[T]
}

func Rect[T constraintsi.Number](x0, y0, x1, y1 T) Rectangle[T] {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle[T]{Point[T]{x0, y0}, Point[T]{x1, y1}}
}
