/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

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
	mat := NewTranslateRotationMat(point1B, point1A, angle)
	assert.Equal(t, point2A, mat.Transform(point2B))
	angle = NewVector(point1A, point2A).AngleWith(NewVector(point1B, point2B))
	t.Log(angle)
	mat = NewTranslateRotationMat(point1B, point1A, -angle)
	t.Log(mat.Transform(point2B))
}

func TestIsPointInRectangle(t *testing.T) {
	rect1 := NewRect(Point{X: 0.5, Y: 0.5}, 0.4, 0.2, 30)
	rect2 := NewRect(Point{X: 0.5, Y: -0.5}, 0.4, 0.2, -30)
	assert.Equal(t, false, rect1.ContainsPoint(Point{X: 0.44, Y: 0.62}))
	assert.Equal(t, false, rect2.ContainsPoint(Point{X: 0.44, Y: -0.62}))
	assert.Equal(t, true, rect1.ContainsPoint(Point{X: 0.58, Y: 0.54}))
	assert.Equal(t, true, rect2.ContainsPoint(Point{X: 0.58, Y: -0.54}))
	assert.Equal(t, true, NewRect(Point{X: 596, Y: 1491}, 1129.5, 2957, 0).ContainsPoint(Point{X: 82, Y: 12.5}))
	assert.Equal(t, true, NewRect(Point{X: 596, Y: -1491}, 1129.5, 2957, 0).ContainsPoint(Point{X: 82, Y: -12.5}))
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
		assert.Equal(t, NewRect(Point{cx, cy}, w, h, angle).ContainsPoint(Point{x, y}), NewRect(
			Point{cx, -cy}, w, h, -angle).ContainsPoint(Point{x, -y}))
	})
}

func TestAffineMatrix(t *testing.T) {
	p1, p2, p3, q1, q2, q3 := Point{2000, 7000}, Point{48000, 80000}, Point{2000, 85000}, Point{3558, 17895}, Point{11016, 5997}, Point{3538, 5182}
	transformMatrix, err := NewAffineMatrix(p1, p2, p3, q1, q2, q3)
	if err != nil {
		t.Error(err)
	}

	// 对某个点应用变换
	p := Point{X: 48000, Y: 13000}
	q := transformMatrix.Transform(p)
	t.Log(q)
}

func TestRotate(t *testing.T) {
	w, h := float64(277), float64(199)
	t.Log(NewRect(Point{}, w, h, 45).Corners())
	t.Log(Point{-w / 2, h / 2}.Rotate(Point{}, 45))
	t.Log(Point{-w / 2, -h / 2}.Rotate(Point{}, 45))
}

func TestAngle(t *testing.T) {
	p1, p2 := Point{X: 1, Y: 1}, Point{X: 2, Y: 2}
	v1, v2 := Vector{X: 1, Y: 1}, Vector{X: 2, Y: 2}
	t.Log(v1.AngleWith(v2))
	t.Log(NewVector(p1, p2).Angle())
	t.Log(v1.AngleWith(v2))
}
