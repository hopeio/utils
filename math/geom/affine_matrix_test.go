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
