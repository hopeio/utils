package geometry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransformPoint(t *testing.T) {
	angle := -30.0 // degrees
	point1A := Point{7.46, -6.73}
	point1B := Point{4.0, -2.0}
	point2B := Point{5.0, -3.0}
	point2A := Point{X: 7.826025403784439, Y: -8.09602540378444}
	// {7.826025403784439 8.09602540378444}
	assert.Equal(t, point2A, TransformPointByOnePointAndRotationAngle(point2B, point1B, point1A, angle))
	angle = CalculateRotationAngle(point1A, point1B, point2A, point2B)
	t.Log(angle)
	t.Log(TransformPointByOnePointAndRotationAngle(point2B, point1B, point1A, -angle))
}

func TestIsPointInRectangle(t *testing.T) {
	assert.Equal(t, false, IsPointInRectangle(Point{X: 0.44, Y: 0.62}, Point{X: 0.5, Y: 0.5}, 0.4, 0.2, 30))
	assert.Equal(t, false, IsPointInRectangle(Point{X: 0.44, Y: -0.62}, Point{X: 0.5, Y: -0.5}, 0.4, 0.2, -30))
	assert.Equal(t, true, IsPointInRectangle(Point{X: 0.58, Y: 0.54}, Point{X: 0.5, Y: 0.5}, 0.4, 0.2, 30))
	assert.Equal(t, true, IsPointInRectangle(Point{X: 0.58, Y: -0.54}, Point{X: 0.5, Y: -0.5}, 0.4, 0.2, -30))
	assert.Equal(t, true, IsPointInRectangle(Point{X: 82, Y: 12.5}, Point{X: 596, Y: 1491}, 1129.5, 2957, 0))
	assert.Equal(t, true, IsPointInRectangle(Point{X: 82, Y: -12.5}, Point{X: 596, Y: -1491}, 1129.5, 2957, 0))
}

func FuzzIsPointInRectangle(f *testing.F) {
	f.Fuzz(func(t *testing.T, x, y, cx, cy, w, h, angle float64) {
		angle = NormalizeAngleDegrees(angle)
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		if w < 0 {
			w = -w
		}
		if h < 0 {
			h = -h
		}
		if cx < 0 {
			cx = -cx
		}
		if cy < 0 {
			cy = -cy
		}
		t.Log(x, y, cx, cy, w, h, angle)
		assert.Equal(t, IsPointInRectangle(Point{x, y}, Point{cx, cy}, w, h, angle), IsPointInRectangle(Point{x, -y},
			Point{cx, -cy}, w, h, -angle))
	})
}
