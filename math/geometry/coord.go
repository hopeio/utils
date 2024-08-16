package geometry

import (
	"math"
)

// Point 结构体用于表示一个点
type Point struct {
	X float64
	Y float64
}

// 两个原点不重合的坐标系O1,O2。O2在O1内部,且经过顺时针旋转c度。其中的点分别用(x1,y1),(x2,y2)表示，已知某个点在两个坐标系中的坐标(x1,y1),(x2,y2),以及另一点在O2内的坐标(x2,
//y2)，求该点在O1内的坐标(x1,y1).
// 图像如何转换，图像可看做第四象限，输入-y,返回-y

// TransformPointByOnePointAndRotationAngle transforms a point from coordinate system a2 to a1
// 在数学和计算机图形学中，旋转角度的正负通常遵循右手定则。默认情况下，顺时针方向被认为是负的，而逆时针方向被认为是正的。
// O2相对于O1旋转度数
func TransformPointByOnePointAndRotationAngle(pA, qA, qB Point, angleDeg float64) (pB Point) {

	// Convert angle from degrees to radians
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosC := math.Cos(angleRad)
	sinC := math.Sin(angleRad)

	// Calculate dx and dy
	dx := qB.X - (qA.X*cosC - qA.Y*sinC)
	dy := qB.Y - (qA.X*sinC + qA.Y*cosC)

	// Apply rotation and translation
	x1 := pA.X*cosC - pA.Y*sinC + dx
	y1 := pA.X*sinC + pA.Y*cosC + dy

	return Point{x1, y1}
}

// 已知两点在两平面坐标系中的坐标，求坐标系夹角,A旋转到B的度数，逆时针>0
// CalculateRotationAngle
func CalculateRotationAngle(pA, pB, qA, qB Point) float64 {
	// 计算向量pA->pB和qA->qB
	vectorP := Point{pB.X - pA.X, pB.Y - pA.Y}
	vectorQ := Point{qB.X - qA.X, qB.Y - qA.Y}
	// 计算两个向量之间的夹角
	angleInRadians := angleBetweenVectors(vectorP, vectorQ)

	// 将弧度转换为度
	return angleInRadians * (180.0 / math.Pi)
}

// 计算两个向量之间的夹角（以弧度为单位）
func angleBetweenVectors(v1, v2 Point) float64 {
	dx := v2.X - v1.X
	dy := v2.Y - v1.Y
	return math.Atan2(dy, dx)
}
