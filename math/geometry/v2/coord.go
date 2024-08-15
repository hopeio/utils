package geometry

import (
	constraintsi "github.com/hopeio/utils/types/constraints"
	"math"
)

// 两个原点不重合的坐标系O1,O2,均为x轴向右，y轴向下。O2在O1内部,且经过顺时针旋转c度。其中的点分别用(x1,y1),(x2,y2)表示，已知某个点在两个坐标系中的坐标(x1,y1),(x2,y2),以及另一点在O2内的坐标(x2,y2)，求该点在O1内的坐标(x1,y1)
// ------------------------------
// |         、
// |       /     、
// |      /          、
// |     /               、
// |    /
// |
// |
// Point 结构体用于表示一个点
type Point[T constraintsi.Number] struct {
	X T
	Y T
}

// TransformPoint transforms a point from coordinate system a2 to a1
func TransformPoint[T constraintsi.Number](p1InO2, p2InO1, p2InO2 Point[T], angleDeg float64) Point[T] {

	// Convert angle from degrees to radians
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosC := math.Cos(angleRad)
	sinC := math.Sin(angleRad)

	// Calculate dx and dy
	dx := float64(p2InO1.X) - float64(p2InO2.X)*cosC - float64(p2InO2.Y)*sinC
	dy := float64(p2InO1.Y) - float64(p2InO2.X)*sinC + float64(p2InO2.Y)*cosC

	// Apply rotation and translation
	x1 := float64(p1InO2.X)*cosC - float64(p1InO2.Y)*sinC + dx
	y1 := float64(p1InO2.X)*sinC + float64(p1InO2.Y)*cosC + dy

	return Point[T]{T(x1), T(y1)}
}
