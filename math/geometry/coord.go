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

func IsPointInRotatedRectangle(Px, Py, Cx, Cy, W, H, theta float64) bool {
	// 将角度转换为弧度
	theta = theta * math.Pi / 180
	Cost := math.Cos(theta)
	Sint := math.Sin(theta)
	// 计算矩形四个角的坐标
	Dx := Cx + (W/2)*Cost - (H/2)*Sint
	Dy := Cy + (W/2)*Sint + (H/2)*Cost
	Ax := Cx - (W/2)*Cost - (H/2)*Sint
	Ay := Cy - (W/2)*Sint + (H/2)*Cost
	Bx := Cx - (W/2)*Cost + (H/2)*Sint
	By := Cy - (W/2)*Sint - (H/2)*Cost
	Cx = Cx + (W/2)*Cost + (H/2)*Sint
	Cy = Cy + (W/2)*Sint - (H/2)*Cost

	// 射线法判断点是否在矩形内
	inside := false
	intersections := 0
	corners := [][]float64{{Ax, Ay}, {Bx, By}, {Cx, Cy}, {Dx, Dy}}

	for i := 0; i < len(corners); i++ {
		x1, y1 := corners[i][0], corners[i][1]
		x2, y2 := corners[(i+1)%len(corners)][0], corners[(i+1)%len(corners)][1]

		if y1 == y2 { // 水平边
			continue
		}
		if Py < min(y1, y2) || Py >= max(y1, y2) { // 在边的外部
			continue
		}

		x_intersect := x1 + (Py-y1)*(x2-x1)/(y2-y1)
		if Px < x_intersect {
			intersections++
		}
	}

	if intersections%2 == 1 {
		inside = true
	}

	return inside
}
