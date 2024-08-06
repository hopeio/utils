package coordinate

import "testing"

func TestTransformPoint(t *testing.T) {
	angle := 30.0 // degrees
	pointA1 := Point{7.46, 6.73}
	pointA2 := Point{4.0, 2.0}
	anotherPointA2 := Point{5.0, 3.0}
	t.Log(TransformPoint(anotherPointA2, pointA1, pointA2, angle))
}
