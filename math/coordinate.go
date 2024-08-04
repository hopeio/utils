package math

import "math"

// 两个原点不重合的坐标系O1,O2,均为x轴向右，y轴向下。O2在O1内部,且经过顺时针旋转c度。其中的点分别用(x1,y1),(x2,y2)表示，已知某个点在两个坐标系中的坐标(x1,y1),(x2,y2),以及另一点在O2内的坐标(x2,y2)，求该点在O1内的坐标(x1,y1)

// Point 结构体用于表示一个点
type Point struct {
	X float64
	Y float64
}

// TransformPoint transforms a point from coordinate system a2 to a1
func TransformPoint(p, p1, p2 Point, angleDeg float64) Point {

	// Convert angle from degrees to radians
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosC := math.Cos(angleRad)
	sinC := math.Sin(angleRad)

	// Calculate dx and dy
	dx := p1.X - (p2.X*cosC - p2.Y*sinC)
	dy := p1.Y - (p2.X*sinC + p2.Y*cosC)

	// Apply rotation and translation
	x1 := p.X*cosC - p.Y*sinC + dx
	y1 := p.X*sinC + p.Y*cosC + dy

	return Point{x1, y1}
}
