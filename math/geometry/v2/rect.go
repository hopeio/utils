/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package geometry

import constraintsi "github.com/hopeio/utils/types/constraints"

type Rectangle[T constraintsi.Number] struct {
	Center Point[T]
	Width  T
	Height T
	Angle  float64
}

func NewRect[T constraintsi.Number](center Point[T], width, height T, angle float64) *Rectangle[T] {
	return &Rectangle[T]{center, width, height, angle}
}
