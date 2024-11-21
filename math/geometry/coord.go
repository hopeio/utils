package geometry

import (
	"fmt"
	"math"
	"math/rand/v2"
)

// Point 结构体用于表示一个点
type Point struct {
	X float64
	Y float64
}

func RandomPoint(min, max Point) Point {
	return Point{
		X: math.Floor(min.X + math.Floor(rand.Float64()*(max.X-min.X))),
		Y: math.Floor(min.Y + math.Floor(rand.Float64()*(max.Y-min.Y))),
	}
}

func (p Point) Vector(point Point) Vector {
	return Vector{point.X - p.X, point.Y - p.Y}
}

func (p Point) Rotate(center Point, angleDeg float64) Point {
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	// 计算旋转后的坐标
	newX := center.X + (p.X-center.X)*cosA - (p.Y-center.Y)*sinA
	newY := center.Y + (p.X-center.X)*sinA + (p.Y-center.Y)*cosA

	return Point{newX, newY}
}

// PointsLength 计算两点之间的向量长度
func (p Point) Length(p2 Point) float64 {
	dx := p2.X - p.X
	dy := p2.Y - p.Y
	return math.Sqrt(dx*dx + dy*dy)
}

type Point3D struct {
	X float64
	Y float64
	Z float64
}

type Vector struct {
	X float64
	Y float64
}

func NewVector(p1, p2 Point) Vector {
	return Vector{p2.X - p1.X, p2.Y - p1.Y}
}

func (v Vector) Angle() float64 {
	angleRadians := math.Atan2(v.Y, v.X)
	return angleRadians * (180.0 / math.Pi)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 计算同坐标两个向量之间的角度
func (v Vector) AngleWith(v2 Vector) float64 {
	dotProduct := v.X*v2.X + v.Y*v2.Y
	magnitudeV1 := math.Sqrt(v.X*v.X + v.Y*v.Y)
	magnitudeV2 := math.Sqrt(v2.X*v2.X + v2.Y*v2.Y)
	angleInRadians := math.Acos(dotProduct / (magnitudeV1 * magnitudeV2))
	return angleInRadians * (180.0 / math.Pi)
}

type AngleDegrees float64
type AngleRadian float64

func (a AngleDegrees) Radian() AngleRadian {
	return AngleRadian(a * math.Pi / 180.0)
}

func (a AngleDegrees) Normalize() AngleDegrees {
	if a == 0 {
		return 0
	}

	if a > 0 {
		for a > 360 {
			a -= 360
		}
	} else {
		a += 360
		for a < 0 {
			a += 360
		}
	}
	return a
}

func (a AngleRadian) Degrees() AngleDegrees {
	return AngleDegrees(a / math.Pi * 180.0)
}

func NormalizeAngleDegrees(theta float64) float64 {
	if theta == 0 {
		return 0
	}

	if theta > 0 {
		for theta > 360 {
			theta -= 360
		}
	} else {
		theta += 360
		for theta < 0 {
			theta += 360
		}
	}
	return theta
}

// 定义一个2x3的仿射变换矩阵
type AffineMatrix [2][3]float64

// 应用仿射变换到点上
func (m AffineMatrix) Transform(p Point) Point {
	return Point{
		X: m[0][0]*p.X + m[0][1]*p.Y + m[0][2],
		Y: m[1][0]*p.X + m[1][1]*p.Y + m[1][2],
	}
}

func NewRotationMat(center Point, angleDeg float64) AffineMatrix {
	angleRad := angleDeg * math.Pi / 180.0
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	return AffineMatrix{{cosA, -sinA, center.X - cosA*center.X + sinA*center.Y}, {sinA, cosA, center.Y - sinA*center.X - cosA*center.Y}}
}

// 两个原点不重合的坐标系O1,O2。O2在O1内部,且经过顺时针旋转c度。其中的点分别用(x1,y1),(x2,y2)表示，已知某个点在两个坐标系中的坐标(x1,y1),(x2,y2),以及另一点在O2内的坐标(x2,
//y2)，求该点在O1内的坐标(x1,y1).
// 图像如何转换，图像可看做第四象限，输入-y,返回-y

// NewTranslateRotationMat transforms a point from coordinate system a2 to a1
// 在数学和计算机图形学中，旋转角度的正负通常遵循右手定则。默认情况下，顺时针方向被认为是负的，而逆时针方向被认为是正的。
// O2相对于O1旋转度数
func NewTranslateRotationMat(pA, pB Point, angleDeg float64) AffineMatrix {
	// Convert angle from degrees to radians
	angleRad := angleDeg * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	return AffineMatrix{
		{cosA, -sinA, pB.X - cosA*pA.X + sinA*pA.Y},
		{sinA, cosA, pB.Y - sinA*pA.X - cosA*pA.Y},
	}
}

// 计算仿射变换矩阵
func NewAffineMatrix(p1, p2, p3, q1, q2, q3 Point) (AffineMatrix, error) {
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
	solution, err := GaussJordanElimination(A, b)
	if err != nil {
		return AffineMatrix{}, err
	}

	// 构造仿射变换矩阵
	transformMatrix := AffineMatrix{
		{solution[0], solution[1], solution[2]},
		{solution[3], solution[4], solution[5]},
	}

	return transformMatrix, nil
}

// GaussJordanElimination 高斯-约旦消元法求解线性方程组 Ax = b
func GaussJordanElimination(A [][]float64, b []float64) ([]float64, error) {
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
