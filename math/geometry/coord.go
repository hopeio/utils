package geometry

import (
	"math"
	"math/rand"
)

// Point 结构体用于表示一个点
type Point struct {
	X float64
	Y float64
}

func (p *Point) Rotate(angleDeg float64) {
	angleRad := math.Pi * angleDeg / 180.0
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	p.X, p.Y = p.X*cosA-p.Y*sinA, p.X*sinA+p.Y*cosA
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
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)

	// Calculate dx and dy
	dx := qB.X - (qA.X*cosA - qA.Y*sinA)
	dy := qB.Y - (qA.X*sinA + qA.Y*cosA)

	// Apply rotation and translation
	x1 := pA.X*cosA - pA.Y*sinA + dx
	y1 := pA.X*sinA + pA.Y*cosA + dy

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

// 图片就是第四象限,角度90+θ
func IsPointInRectangle(p Point, rCenter Point, W, H, angleDeg float64) bool {
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	// 计算矩形四个角的坐标 (A左下-B右下-C右上-D左上)
	Dx := rCenter.X + (W/2)*cosA - (H/2)*sinA
	Dy := rCenter.Y + (W/2)*sinA + (H/2)*cosA
	Ax := rCenter.X - (W/2)*cosA - (H/2)*sinA
	Ay := rCenter.Y - (W/2)*sinA + (H/2)*cosA
	Bx := rCenter.X - (W/2)*cosA + (H/2)*sinA
	By := rCenter.Y - (W/2)*sinA - (H/2)*cosA
	Cx := rCenter.X + (W/2)*cosA + (H/2)*sinA
	Cy := rCenter.X + (W/2)*sinA - (H/2)*cosA

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
		if p.Y < min(y1, y2) || p.Y >= max(y1, y2) { // 在边的外部
			continue
		}

		x_intersect := x1 + (p.Y-y1)*(x2-x1)/(y2-y1)
		if p.X < x_intersect {
			intersections++
		}
	}

	if intersections%2 == 1 {
		inside = true
	}

	return inside
}

func RandomPoint(min, max Point) Point {
	return Point{
		X: math.Floor(min.X + math.Floor(rand.Float64()*(max.X-min.X))),
		Y: math.Floor(min.Y + math.Floor(rand.Float64()*(max.Y-min.Y))),
	}
}

// RotatePoint 计算点 (x, y) 绕点 (centerX, centerY) 旋转 angle 度后的新坐标
func RotatePoint(p Point, center Point, angleDeg float64) Point {
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	// 计算旋转后的坐标
	newX := center.X + (p.X-center.X)*cosA - (p.Y-center.Y)*sinA
	newY := center.Y + (p.X-center.X)*sinA + (p.Y-center.Y)*cosA

	return Point{newX, newY}
}

func NormalizeAngleDegrees(theta float64) float64 {
	normalized := math.Mod(theta, 360)
	if normalized < 0 {
		normalized += 360
	}
	return normalized
}
