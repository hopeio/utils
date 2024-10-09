package geometry

import (
	constraintsi "github.com/hopeio/utils/types/constraints"
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

type Point3D[T constraintsi.Number] struct {
	X T
	Y T
	Z T
}
