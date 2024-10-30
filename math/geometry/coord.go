package geometry

import (
	"fmt"
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

type Point3D struct {
	X float64
	Y float64
	Z float64
}

// 两个原点不重合的坐标系O1,O2。O2在O1内部,且经过顺时针旋转c度。其中的点分别用(x1,y1),(x2,y2)表示，已知某个点在两个坐标系中的坐标(x1,y1),(x2,y2),以及另一点在O2内的坐标(x2,
//y2)，求该点在O1内的坐标(x1,y1).
// 图像如何转换，图像可看做第四象限，输入-y,返回-y

// TranslateRotationTransformByPointAndAngle transforms a point from coordinate system a2 to a1
// 在数学和计算机图形学中，旋转角度的正负通常遵循右手定则。默认情况下，顺时针方向被认为是负的，而逆时针方向被认为是正的。
// O2相对于O1旋转度数
func TranslateRotationTransformByPointAndAngle(pA, qA, qB Point, angleDeg float64) (pB Point) {

	// Convert angle from degrees to radians
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)

	// Apply rotation and translation
	dx := pA.X*cosA - pA.Y*sinA - qA.X*cosA + qA.Y*sinA
	dy := pA.X*sinA + pA.Y*cosA - qA.X*sinA - qA.Y*cosA

	return Point{dx + qB.X, dy + qB.Y}
}

// 已知两点在两平面坐标系中的坐标，求坐标系夹角,A旋转到B的度数，逆时针>0
// AngleBetweenVectors
func AngleBetweenVectors(pA, pB, qA, qB Point) float64 {
	// 计算向量pA->pB和qA->qB
	vectorP := Point{pB.X - pA.X, pB.Y - pA.Y}
	vectorQ := Point{qB.X - qA.X, qB.Y - qA.Y}
	// 计算两个向量之间的夹角
	dx := vectorQ.X - vectorP.X
	dy := vectorQ.Y - vectorP.Y
	angleInRadians := math.Atan2(dy, dx)
	// 将弧度转换为度
	return angleInRadians * (180.0 / math.Pi)
}

func RectangleCorners(rCenter Point, w, h, angleDeg float64) [][]float64 {
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	// 计算矩形四个角的坐标 (A左下-B右下-C右上-D左上)
	Dx := rCenter.X + (w/2)*cosA - (h/2)*sinA
	Dy := rCenter.Y + (w/2)*sinA + (h/2)*cosA
	Ax := rCenter.X - (w/2)*cosA - (h/2)*sinA
	Ay := rCenter.Y - (w/2)*sinA + (h/2)*cosA
	Bx := rCenter.X - (w/2)*cosA + (h/2)*sinA
	By := rCenter.Y - (w/2)*sinA - (h/2)*cosA
	Cx := rCenter.X + (w/2)*cosA + (h/2)*sinA
	Cy := rCenter.Y + (w/2)*sinA - (h/2)*cosA
	return [][]float64{{Ax, Ay}, {Bx, By}, {Cx, Cy}, {Dx, Dy}}
}

// 图片就是第四象限,角度90+θ
func IsPointInRectangle(p Point, rCenter Point, w, h, angleDeg float64) bool {

	// 射线法判断点是否在矩形内
	inside := false
	intersections := 0
	corners := RectangleCorners(rCenter, w, h, angleDeg)

	for i := 0; i < len(corners); i++ {
		x1, y1 := corners[i][0], corners[i][1]
		x2, y2 := corners[(i+1)%len(corners)][0], corners[(i+1)%len(corners)][1]

		if y1 == y2 { // 水平边
			continue
		}
		if p.Y < min(y1, y2) || p.Y > max(y1, y2) { // 在边的外部
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

// 定义一个2x3的仿射变换矩阵
type AffineMatrix [2][3]float64

// 应用仿射变换到点上
func (m AffineMatrix) Apply(p Point) Point {
	return Point{
		X: m[0][0]*p.X + m[0][1]*p.Y + m[0][2],
		Y: m[1][0]*p.X + m[1][1]*p.Y + m[1][2],
	}
}

// 计算同坐标两个向量之间的角度
func AngleBetweenPoints(v1, v2 Point) float64 {
	dotProduct := v1.X*v2.X + v1.Y*v2.Y
	magnitudeV1 := math.Sqrt(v1.X*v1.X + v1.Y*v1.Y)
	magnitudeV2 := math.Sqrt(v2.X*v2.X + v2.Y*v2.Y)
	return math.Acos(dotProduct / (magnitudeV1 * magnitudeV2))
}

// 计算仿射变换矩阵
func calculateAffineTransform(p1, p2, p3, q1, q2, q3 Point) ([][]float64, error) {
	// 构造线性方程组的系数矩阵A和常数向量b
	A := [][]float64{
		{p1.X, p1.Y, 1, 0, 0, 0},
		{0, 0, 0, p1.X, p1.Y, 1},
		{p2.X, p2.Y, 1, 0, 0, 0},
		{0, 0, 0, p2.X, p2.Y, 1},
		{p3.X, p3.Y, 1, 0, 0, 0},
		{0, 0, 0, p3.X, p3.Y, 1},
	}
	b := []float64{q1.X, q1.Y, q2.X, q2.Y, q3.X, q3.Y}

	// 使用高斯-约旦消元法求解线性方程组 Ax = b
	solution, err := gaussJordanElimination(A, b)
	if err != nil {
		return nil, err
	}

	// 构造仿射变换矩阵
	transformMatrix := [][]float64{
		{solution[0], solution[1], solution[2]},
		{solution[3], solution[4], solution[5]},
		{0, 0, 1},
	}

	return transformMatrix, nil
}

// 应用仿射变换
func applyAffineTransform(m [][]float64, p Point) (Point, error) {
	// 将点转换为齐次坐标
	homogeneousP := []float64{p.X, p.Y, 1}

	// 应用变换矩阵
	result := make([]float64, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i] += m[i][j] * homogeneousP[j]
		}
	}

	// 转换回非齐次坐标
	if result[2] == 0 {
		return Point{}, fmt.Errorf("division by zero in transformation")
	}
	return Point{X: result[0] / result[2], Y: result[1] / result[2]}, nil
}

// 高斯-约旦消元法求解线性方程组 Ax = b
func gaussJordanElimination(A [][]float64, b []float64) ([]float64, error) {
	n := len(b)
	m := len(A)
	if m != n || len(A[0]) != n {
		return nil, fmt.Errorf("invalid matrix dimensions")
	}

	// 扩展矩阵 [A | b]
	extendedMatrix := make([][]float64, n)
	for i := range extendedMatrix {
		extendedMatrix[i] = make([]float64, n+1)
		copy(extendedMatrix[i][:n], A[i])
		extendedMatrix[i][n] = b[i]
	}

	// 高斯-约旦消元法
	for i := 0; i < n; i++ {
		// 寻找主元素
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(extendedMatrix[k][i]) > math.Abs(extendedMatrix[maxRow][i]) {
				maxRow = k
			}
		}

		// 交换行
		extendedMatrix[i], extendedMatrix[maxRow] = extendedMatrix[maxRow], extendedMatrix[i]

		// 主元为0则无法继续
		if extendedMatrix[i][i] == 0 {
			return nil, fmt.Errorf("matrix is singular")
		}

		// 消元
		pivot := extendedMatrix[i][i]
		for j := 0; j < n+1; j++ {
			extendedMatrix[i][j] /= pivot
		}
		for k := 0; k < n; k++ {
			if k != i {
				factor := extendedMatrix[k][i]
				for j := 0; j < n+1; j++ {
					extendedMatrix[k][j] -= factor * extendedMatrix[i][j]
				}
			}
		}
	}

	// 提取解
	solution := make([]float64, n)
	for i := 0; i < n; i++ {
		solution[i] = extendedMatrix[i][n]
	}

	return solution, nil
}

// VectorLength 计算两点之间的向量长度
func VectorLength(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// VectorAngle 计算向量与 x 轴之间的角度（以度为单位）
func VectorAngle(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	angleRadians := math.Atan2(dy, dx)
	angleDegrees := angleRadians * (180.0 / math.Pi)
	return angleDegrees
}
