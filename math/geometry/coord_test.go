package geometry

import (
	"math"
	"testing"
)

func TestTransformPoint(t *testing.T) {
	angle := -30.0 // degrees
	point1A := Point{7.46, -6.73}
	point1B := Point{4.0, -2.0}
	point2B := Point{5.0, -3.0}
	point2A := Point{X: 7.83, Y: -8.10}
	// {7.826025403784439 8.09602540378444}
	t.Log(TransformPointByOnePointAndRotationAngle(point2B, point1B, point1A, angle))
	angle = CalculateRotationAngle(point1A, point2A, point1B, point2B)
	t.Log(angle * 180 / math.Pi)
	t.Log(TransformPointByOnePointAndRotationAngle(point2B, point1B, point1A, angle))
}
