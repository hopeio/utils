/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gocv

import (
	"image"
	"testing"
)

func TestAffineMatrix(t *testing.T) {
	p1, p2, p3, q1, q2, q3 := image.Point{2000, 7000}, image.Point{48000, 80000}, image.Point{2000, 85000}, image.Point{3558, 17895}, image.Point{11016, 5997}, image.Point{3538, 5182}
	t.Log(AffineTransform(p1, p2, p3, q1, q2, q3, []image.Point{{X: 48000, Y: 13000}}))
}
