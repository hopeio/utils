package image

import (
	"image"
	"image/color"
	"image/draw"
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
func CalculateOverlap(img1, img2 image.Image, row bool, minOverlap, maxOverlap int) int {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	dx1, dy1 := bounds1.Dx(), bounds1.Dy()
	dx2, dy2 := bounds2.Dx(), bounds2.Dy()
	dx, dy := min(dx1, dx2), min(dy1, dy2)
	minX1, minY1 := max(bounds1.Min.X, bounds1.Max.X-dx), bounds1.Min.Y
	minX2, minY2 := bounds2.Min.X, bounds2.Min.Y
	maxX1, maxY1 := minX1+dx, minY1+dy
	maxY2 := minY2 + dy
	maxOverlap = min(maxOverlap, dy)
	if row {
		minX1, minY1 = max(bounds1.Min.Y, bounds1.Max.Y-dy), bounds1.Min.X
		minX2, minY2 = bounds2.Min.Y, bounds2.Min.X
		maxX1, maxY1 = minY1+dy, minX1+dx
		maxY2 = minX2 + dx
		dy = dx
	}
	data1 := make([]uint8, 0, maxOverlap*dy)
	// 遍历原始图像的每个像素并转换为灰度值
	for x := maxX1 - maxOverlap; x < maxX1; x++ {
		for y := minY1; y < maxY1; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data1 = append(data1, gray)
		}
	}
	data2 := make([]uint8, 0, maxOverlap*dy)
	for x := minX2; x <= maxOverlap; x++ {
		for y := minY2; y < maxY2; y++ {
			r, g, b, _ := img2.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data2 = append(data2, gray)
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

func ColorToGray(c color.Color) uint8 {
	if g, ok := c.(color.Gray); ok {
		return g.Y
	}
	r, g, b, _ := c.RGBA()

	// These coefficients (the fractions 0.299, 0.587 and 0.114) are the same
	// as those given by the JFIF specification and used by func RGBToYCbCr in
	// ycbcr.go.
	//
	// Note that 19595 + 38470 + 7471 equals 65536.
	//
	// The 24 is 16 + 8. The 16 is the same as used in RGBToYCbCr. The 8 is
	// because the return value is 8 bit color, not 16 bit color.
	y := (19595*r + 38470*g + 7471*b + 1<<15) >> 24

	return uint8(y)
}

func RGBAToGray(c color.RGBA) uint8 {
	return uint8(0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B))
}

func RGBToGray(c color.RGBA) uint8 {
	return uint8(0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B))
}

func MergeImages(imgs [][]int, getImage func(int) image.Image, bounds image.Rectangle, horizontalOverlaps,
	verticalOverlaps []int) image.Image {
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
			resultWidth -= verticalOverlaps[i]
		}
	}

	// 创建一个新的 RGBA 图片，用于存储合并后的图片
	result := image.NewRGBA(image.Rect(0, 0, resultWidth, resultHeight))
	var rbounds = bounds

	// 将 img1 复制到结果图片中
	for i, rimg := range imgs {
		for j, imgIdx := range rimg {
			img := getImage(imgIdx)
			draw.Draw(result, rbounds, img, image.Point{}, draw.Src)
			if j < len(horizontalOverlaps) {
				rbounds.Min.X += bounds.Dx() - horizontalOverlaps[j]
				rbounds.Max.X += bounds.Dx() + rbounds.Min.X
			}
		}
		if i < len(verticalOverlaps) {
			rbounds.Min.Y += bounds.Dy() - verticalOverlaps[i]
			rbounds.Max.Y += bounds.Dy() + rbounds.Min.Y
			rbounds.Min.X = 0
			rbounds.Max.X = bounds.Dx()
		}
	}

	return result
}
