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
	assert.Equal(t, point2A, TranslateRotationTransformByPointAndAngle(point2B, point1B, point1A, angle))
	angle = VectorsAngle(NewVector(point1A, point2A), NewVector(point1B, point2B))
	t.Log(angle)
	t.Log(TranslateRotationTransformByPointAndAngle(point2B, point1B, point1A, -angle))
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

func TestAffineMatrix(t *testing.T) {
	p1, p2, p3, q1, q2, q3 := Point{2000, 7000}, Point{48000, 80000}, Point{2000, 85000}, Point{3558, 17895}, Point{11016, 5997}, Point{3538, 5182}
	transformMatrix, err := calculateAffineTransform(p1, p2, p3, q1, q2, q3)
	if err != nil {
		t.Error(err)
	}

	// 对某个点应用变换
	p := Point{X: 48000, Y: 13000}
	q, err := applyAffineTransform(transformMatrix, p)
	if err != nil {
		t.Error(err)
	}
	t.Log(q)
}

func TestRotate(t *testing.T) {
	w, h := float64(277), float64(199)
	t.Log(RectangleCorners(Point{}, w, h, 45))
	t.Log(RotationTransformByAngle(Point{-w / 2, h / 2}, 45))
	t.Log(RotationTransformByAngle(Point{-w / 2, -h / 2}, 45))
}

func TestAngle(t *testing.T) {
	p1, p2 := Point{X: 1, Y: 1}, Point{X: 2, Y: 2}
	v1, v2 := Vector{X: 1, Y: 1}, Vector{X: 2, Y: 2}
	t.Log(VectorsAngle(v1, v2))
	t.Log(NewVector(p1, p2).Angle())
	t.Log(VectorsAngle(v1, v2))
}
