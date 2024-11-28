/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package geometry

import (
	"math"
)

type AngleDegrees float64
type AngleRadian float64

func (a AngleDegrees) Radian() AngleRadian {
	return AngleRadian(a * math.Pi / 180.0)
}

func (a AngleDegrees) Normalize() AngleDegrees {
	if a == 0 {
		return 0
	}

	if a > 0 {
		for a > 360 {
			a -= 360
		}
	} else {
		a += 360
		for a < 0 {
			a += 360
		}
	}
	return a
}

func (a AngleRadian) Degrees() AngleDegrees {
	return AngleDegrees(a / math.Pi * 180.0)
}

func NormalizeAngleDegrees(theta float64) float64 {
	if theta == 0 {
		return 0
	}

	if theta > 0 {
		for theta > 360 {
			theta -= 360
		}
	} else {
		theta += 360
		for theta < 0 {
			theta += 360
		}
	}
	return theta
}
