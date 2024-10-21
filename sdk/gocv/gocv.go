package gocv

import (
	imagei "github.com/hopeio/utils/media/image"
	"gocv.io/x/gocv"
	"image"
)

type Circle struct {
	X        int `json:"x"`
	Y        int `json:"y"`
	Diameter int `json:"diameter"`
}

func SearchCircle(path string, rect image.Rectangle) (circles []Circle, err error) {
	gimg := gocv.IMRead(path, gocv.IMReadGrayScale)
	// 定义高斯核的大小和标准差
	ksize := image.Pt(11, 11)
	sigmaX := 0.0
	blurred := gocv.NewMat()
	defer blurred.Close()
	img := gimg.Region(rect)
	defer img.Close()
	gocv.GaussianBlur(img, &blurred, ksize, sigmaX, sigmaX, gocv.BorderDefault)
	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(blurred, &edges, 100, 200)
	circleMap := gocv.NewMat()
	defer circleMap.Close()
	gocv.HoughCirclesWithParams(edges, &circleMap, gocv.HoughGradient, 1, float64(max(rect.Dx(), rect.Dy())), 300,
		10, 50, 300)
	if !circleMap.Empty() {
		for i := 0; i < circleMap.Cols(); i++ {
			v := circleMap.GetVecfAt(0, i)
			x := int(v[0])
			y := int(v[1])
			r := int(v[2])
			// 检查圆是否完整，即圆的边缘不会超出图像边界
			if (x-r) > 0 && (x+r) < gimg.Cols() && (y-r) > 0 && (y+r) < gimg.Rows() {
				circles = append(circles, Circle{X: x + rect.Min.X, Y: y + rect.Min.Y, Diameter: r * 2})
			}
		}
	}
	return
}

// 有一定重合的固定大小的图片拼图
func MergeImagesByOverlap(imgIdxs [][]int, getImage func(int) ([]byte, error), imgWidth, imgHeight int,
	horizontalOverlaps, verticalOverlaps []int, dst string) error {
	var resultWidth, resultHeight int

	for i := range imgIdxs[0] {
		resultWidth += imgWidth
		if i < len(horizontalOverlaps) {
			resultWidth -= horizontalOverlaps[i]
		}
	}
	for i := range imgIdxs {
		resultHeight += imgHeight
		if i < len(verticalOverlaps) {
			resultHeight -= verticalOverlaps[i]
		}
	}

	data, err := getImage(0)
	if err != nil {
		return err
	}
	img0, err := gocv.IMDecode(data, gocv.IMReadAnyColor|gocv.IMReadAnyDepth)
	if err != nil {
		return err
	}
	result := gocv.NewMatWithSize(resultHeight, resultWidth, img0.Type())

	var bounds = image.Rect(0, 0, imgWidth, imgHeight)
	var img gocv.Mat
	// 将 img1 复制到结果图片中
	for i, rowimgs := range imgIdxs {
		for j, imgIdx := range rowimgs {
			if imgIdx != 0 {
				data, err = getImage(imgIdx)
				if err != nil {
					return err
				}
			} else {
				img = img0
			}
			img, err = gocv.IMDecode(data, gocv.IMReadAnyColor|gocv.IMReadAnyDepth)
			if err != nil {
				return err
			}
			rect := result.Region(bounds)
			img.CopyTo(&rect)
			img.Close()
			if j < len(horizontalOverlaps) {
				bounds.Min.X += bounds.Dx() - horizontalOverlaps[j]
				bounds.Max.X = bounds.Dx() + bounds.Min.X
			}
		}
		if i < len(verticalOverlaps) {
			bounds.Min.Y += bounds.Dy() - verticalOverlaps[i]
			bounds.Max.Y = bounds.Dy() + bounds.Min.Y
			bounds.Min.X = 0
			bounds.Max.X = bounds.Dx()
		}
	}
	gocv.IMWrite(dst, result)
	return nil
}

func Sharpness(imgPath string, rect image.Rectangle) (float64, error) {
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale|gocv.IMReadAnyDepth)

	img = img.Region(rect)
	laplacian := gocv.NewMat()
	defer laplacian.Close()
	// 计算拉普拉斯算子的标准差
	gocv.Laplacian(img, &laplacian, gocv.MatTypeCV64F, 1, 1, 0, gocv.BorderDefault)
	// 计算标准差
	mean, stddev := gocv.NewMat(), gocv.NewMat()
	defer mean.Close()
	defer stddev.Close()
	gocv.MeanStdDev(laplacian, &mean, &stddev)
	return stddev.GetDoubleAt(0, 0), nil
}

func AffineMat(p1, p2, p3, q1, q2, q3 image.Point) gocv.Mat {
	src := gocv.NewMatWithSize(3, 1, gocv.MatTypeCV32FC2)
	dst := gocv.NewMatWithSize(3, 1, gocv.MatTypeCV32FC2)
	src.SetFloatAt(0, 0, float32(p1.X))
	src.SetFloatAt(0, 1, float32(p1.Y))
	dst.SetFloatAt(0, 0, float32(q1.X))
	dst.SetFloatAt(0, 1, float32(q1.Y))
	src.SetFloatAt(1, 0, float32(p2.X))
	src.SetFloatAt(1, 1, float32(p2.Y))
	dst.SetFloatAt(1, 0, float32(q2.X))
	dst.SetFloatAt(1, 1, float32(q2.Y))
	src.SetFloatAt(2, 0, float32(p3.X))
	src.SetFloatAt(2, 1, float32(p3.Y))
	dst.SetFloatAt(2, 0, float32(q3.X))
	dst.SetFloatAt(2, 1, float32(q3.Y))
	return gocv.GetAffineTransform2f(gocv.NewPoint2fVectorFromMat(src), gocv.NewPoint2fVectorFromMat(dst))
}

func AffineTransform(p1, p2, p3, q1, q2, q3 image.Point, points []image.Point) []image.Point {
	amat := AffineMat(p1, p2, p3, q1, q2, q3)
	defer amat.Close()
	n := len(points)
	mat := gocv.NewMatWithSize(n, 1, gocv.MatTypeCV32FC2)
	defer mat.Close()
	for i, p := range points {
		mat.SetFloatAt(i, 0, float32(p.X))
		mat.SetFloatAt(i, 1, float32(p.Y))
	}
	oMat := gocv.NewMat()
	defer oMat.Close()
	gocv.Transform(mat, &oMat, amat)
	ret := make([]image.Point, n)
	for i := 0; i < n; i++ {
		ret[i].X, ret[i].Y = int(oMat.GetFloatAt(i, 0)), int(oMat.GetFloatAt(i, 1))
	}
	return ret
}

func CropRotated(img gocv.Mat, centerX, centerY, length, width float64, angle float64) gocv.Mat {
	points := imagei.RectRotateByCenter(int(centerX), int(centerY), int(length), int(width), angle)
	srcPoints := gocv.NewPointVectorFromPoints(points)
	dstPoints := gocv.NewPointVectorFromPoints([]image.Point{
		{X: 0, Y: 0},
		{X: int(length), Y: 0},
		{X: int(length), Y: int(width)},
		{X: 0, Y: int(width)},
	})
	// count perspective transform matrix
	transformMat := gocv.GetPerspectiveTransform(srcPoints, dstPoints)
	srcPoints.Close()
	dstPoints.Close()
	// warp perspective
	dst := gocv.NewMatWithSize(int(length), int(width), img.Type())
	gocv.WarpPerspective(img, &dst, transformMat, image.Point{
		X: int(length),
		Y: int(width),
	})
	transformMat.Close()
	return dst
}
