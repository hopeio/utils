package image

import (
	"image"
	"image/color"
	"math"
)

func Union(rect image.Rectangle, p image.Point) image.Rectangle {
	if p.X < rect.Min.X {
		rect.Min.X = p.X
	}
	if p.X > rect.Max.X {
		rect.Max.X = p.X
	}
	if p.Y < rect.Min.Y {
		rect.Min.Y = p.Y
	}
	if p.Y > rect.Max.Y {
		rect.Max.Y = p.Y
	}
	return rect
}

// 计算两张图的重合像素,第一张的图的后半部分和第二张的前半部分
func CalculateOverlap(img1, img2 image.Image, col bool, minOverlap, maxOverlap int) int {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	dx1, dy1 := bounds1.Dx(), bounds1.Dy()
	dx2, dy2 := bounds2.Dx(), bounds2.Dy()
	dx, dy := min(dx1, dx2), min(dy1, dy2)
	minX1, minY1 := max(bounds1.Min.X, bounds1.Max.X-dx), bounds1.Min.Y
	minX2, minY2 := bounds2.Min.X, bounds2.Min.Y
	maxX1, maxY1 := minX1+dx, minY1+dy
	maxY2 := minY2 + dy
	maxOverlap = min(maxOverlap, dy)
	if col {
		minX1, minY1 = max(bounds1.Min.Y, bounds1.Max.Y-dy), bounds1.Min.X
		minX2, minY2 = bounds2.Min.Y, bounds2.Min.X
		maxX1, maxY1 = minY1+dy, minX1+dx
		maxY2 = minX2 + dx
		dy = dx
	}
	data1 := make([]uint8, maxOverlap*dy)
	// 遍历原始图像的每个像素并转换为灰度值
	var i int
	var c color.Color
	for x := maxX1 - maxOverlap; x < maxX1; x++ {
		for y := minY1; y < maxY1; y++ {
			if col {
				c = img1.At(y, x)
			} else {
				c = img1.At(x, y)
			}
			r, g, b, _ := c.RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
			data1[i] = gray
			i++
		}
	}
	data2 := make([]uint8, maxOverlap*dy)
	var j int
	for x := minX2; x < maxOverlap; x++ {
		for y := minY2; y < maxY2; y++ {
			if col {
				c = img2.At(y, x)
			} else {
				c = img2.At(x, y)
			}
			r, g, b, _ := c.RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
			data2[j] = gray
			j++
		}
	}
	n := len(data1)
	minMean := math.MaxInt
	y := bounds2.Dy()
	var overlap int
	for o := minOverlap; o <= maxOverlap; o++ {
		var sum int
		m := o * y
		subimg1 := data1[n-m:]
		subimg2 := data2[:m]
		for i := range m {
			v := int(subimg1[i]) - int(subimg2[i])
			sum += v * v
		}
		mse := sum / m
		if mse < minMean {
			minMean = mse
			overlap = o
		}
	}

	return overlap
}

func CalculateOverlapReuseMemory(img1, img2 image.Image, col bool, minOverlap, maxOverlap int, gary1, gary2 []uint8) int {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	dx1, dy1 := bounds1.Dx(), bounds1.Dy()
	dx2, dy2 := bounds2.Dx(), bounds2.Dy()
	dx, dy := min(dx1, dx2), min(dy1, dy2)
	minX1, minY1 := max(bounds1.Min.X, bounds1.Max.X-dx), bounds1.Min.Y
	minX2, minY2 := bounds2.Min.X, bounds2.Min.Y
	maxX1, maxY1 := minX1+dx, minY1+dy
	maxY2 := minY2 + dy
	maxOverlap = min(maxOverlap, dy)
	if col {
		minX1, minY1 = max(bounds1.Min.Y, bounds1.Max.Y-dy), bounds1.Min.X
		minX2, minY2 = bounds2.Min.Y, bounds2.Min.X
		maxX1, maxY1 = minY1+dy, minX1+dx
		maxY2 = minX2 + dx
		dy = dx
	}
	if len(gary1) == 0 {
		gary1 = make([]uint8, maxOverlap*dy)
	}
	// 遍历原始图像的每个像素并转换为灰度值
	var i int
	var c color.Color
	for x := maxX1 - maxOverlap; x < maxX1; x++ {
		for y := minY1; y < maxY1; y++ {
			if col {
				c = img1.At(y, x)
			} else {
				c = img1.At(x, y)
			}
			r, g, b, _ := c.RGBA()
			// 使用加权平均公式计算灰度值
			gary1[i] = uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
			i++
		}
	}
	if len(gary2) == 0 {
		gary2 = make([]uint8, maxOverlap*dy)
	}
	var j int
	for x := minX2; x < maxOverlap; x++ {
		for y := minY2; y < maxY2; y++ {
			if col {
				c = img2.At(y, x)
			} else {
				c = img2.At(x, y)
			}
			r, g, b, _ := c.RGBA()
			// 使用加权平均公式计算灰度值
			gary2[j] = uint8((19595*r + 38470*g + 7471*b + 1<<15) >> 24)
			j++
		}
	}
	n := len(gary1)
	minMean := math.MaxInt
	y := bounds2.Dy()
	var overlap int
	for o := minOverlap; o <= maxOverlap; o++ {
		var sum int
		m := o * y
		subimg1 := gary1[n-m:]
		subimg2 := gary2[:m]
		for i := range m {
			v := int(subimg1[i]) - int(subimg2[i])
			sum += v * v
		}
		mse := sum / m
		if mse < minMean {
			minMean = mse
			overlap = o
		}
	}

	return overlap
}

func RectRotateByCenter(x, y, l, w int, angle float64) []image.Point {
	rad := angle / 180 * math.Pi
	lSinYAxis := int(float64(l) / 2 * math.Sin(rad))
	lCosXAxis := int(float64(l) / 2 * math.Cos(rad))
	wSinXAxis := int(float64(w) / 2 * math.Sin(rad))
	wCosYAxis := int(float64(w) / 2 * math.Cos(rad))
	return []image.Point{
		{X: x - lCosXAxis - wSinXAxis, Y: y + lSinYAxis - wCosYAxis},
		{X: x + lCosXAxis - wSinXAxis, Y: y - lSinYAxis - wCosYAxis},
		{X: x + lCosXAxis + wSinXAxis, Y: y - lSinYAxis + wCosYAxis},
		{X: x - lCosXAxis + wSinXAxis, Y: y + lSinYAxis + wCosYAxis},
	}
}

type MergeImage struct {
	imgs                                 [][]image.Image
	horizontalOverlaps, verticalOverlaps []int
	width, height                        int
}
