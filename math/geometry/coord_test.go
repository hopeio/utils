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

func TestIsPointInRotatedRectangle(t *testing.T) {

}
