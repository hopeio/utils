package geometry

type Rectangle struct {
	Min, Max Point
}

func Rect(x0, y0, x1, y1 float64) Rectangle {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle{Point{x0, y0}, Point{x1, y1}}
}

type Triangle struct {
	A, B, C Point
}
