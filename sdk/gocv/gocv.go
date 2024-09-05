package gocv

import (
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

func MergeImage(imgs [][]int, getImage func(int) ([]byte, error), bounds image.Rectangle,
	horizontalOverlaps,
	verticalOverlaps []int, dst string) error {
	var resultWidth, resultHeight int
	for i := range imgs[0] {
		resultWidth += bounds.Dx()
		if i < len(horizontalOverlaps) {
			resultWidth -= horizontalOverlaps[i]
		}
	}
	for i := range imgs {
		resultHeight += bounds.Dy()
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

	var rbounds = bounds
	// 将 img1 复制到结果图片中
	for i, rimg := range imgs {
		for j, imgIdx := range rimg {
			if imgIdx != 0 {
				data, err = getImage(imgIdx)
				if err != nil {
					return err
				}
			}
			img, err := gocv.IMDecode(data, gocv.IMReadAnyColor|gocv.IMReadAnyDepth)
			if err != nil {
				return err
			}
			rect := result.Region(rbounds)
			img.CopyTo(&rect)
			if j < len(horizontalOverlaps) {
				rbounds.Min.X += bounds.Dx() - horizontalOverlaps[j]
				rbounds.Max.X = bounds.Dx() + rbounds.Min.X
			}
		}
		if i < len(verticalOverlaps) {
			rbounds.Min.Y += bounds.Dy() - verticalOverlaps[i]
			rbounds.Max.Y = bounds.Dy() + rbounds.Min.Y
			rbounds.Min.X = 0
			rbounds.Max.X = bounds.Dx()
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
