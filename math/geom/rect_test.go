package geom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPointInRectangle(t *testing.T) {
	rect1 := NewRect(Pt(0.5, 0.5), 0.4, 0.2, 30)
	rect2 := NewRect(Pt(0.5, -0.5), 0.4, 0.2, -30)
	assert.Equal(t, false, rect1.ContainsPoint(Point{X: 0.44, Y: 0.62}))
	assert.Equal(t, false, rect2.ContainsPoint(Point{X: 0.44, Y: -0.62}))
	assert.Equal(t, true, rect1.ContainsPoint(Point{X: 0.58, Y: 0.54}))
	assert.Equal(t, true, rect2.ContainsPoint(Point{X: 0.58, Y: -0.54}))
	assert.Equal(t, true, NewRect(Pt(596, 1491), 1129.5, 2957, 0).ContainsPoint(Point{X: 82, Y: 12.5}))
	assert.Equal(t, true, NewRect(Pt(596, -1491), 1129.5, 2957, 0).ContainsPoint(Point{X: 82, Y: -12.5}))
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
		assert.Equal(t, NewRect(Pt(cx, cy), w, h, angle).ContainsPoint(Point{x, y}), NewRect(
			Pt(cx, -cy), w, h, -angle).ContainsPoint(Point{x, -y}))
	})
}

func TestRotate(t *testing.T) {
	w, h := float64(277), float64(199)
	t.Log(NewRect(Point{}, w, h, 45).Corners())
	t.Log(Point{-w / 2, h / 2}.Rotate(Point{}, 45))
	t.Log(Point{-w / 2, -h / 2}.Rotate(Point{}, 45))
}
