/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package geometry

import (
	"testing"
)

func TestAngle(t *testing.T) {
	p1, p2 := Point{X: 1, Y: 1}, Point{X: 2, Y: 2}
	v1, v2 := Vector{X: 1, Y: 1}, Vector{X: 2, Y: 2}
	t.Log(v1.AngleWith(v2))
	t.Log(NewVector(p1, p2).Angle())
	t.Log(v1.AngleWith(v2))
}
