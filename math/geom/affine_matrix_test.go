package geom

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

	transformMatrix = NewRotationMat(Pt(1, 1), 90)
	t.Log(transformMatrix.Transform(Pt(0, 0)))
	t.Log(transformMatrix.Transform(Pt(2, 2)))
	t.Log(transformMatrix.Transform(Pt(0, 2)))
	t.Log(transformMatrix.Transform(Pt(2, 0)))
}
