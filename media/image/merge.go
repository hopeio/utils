package image

import (
	colori "github.com/hopeio/utils/media/image/color"
	"image"
	"image/color"
	"image/draw"
)

// 有一定重合的固定大小的图片拼图
func MergeUniformBoundsImagesByOverlap(imgIdxs [][]int, getImage func(int) image.Image, imgWidth, imgHeight int,
	horizontalOverlaps, verticalOverlaps []int) image.Image {
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

	// 创建一个新的 RGBA 图片，用于存储合并后的图片
	result := image.NewRGBA(image.Rect(0, 0, resultWidth, resultHeight))
	slideWin := image.Rect(0, 0, imgWidth, imgHeight)
	var img image.Image
	// 将 img1 复制到结果图片中
	for i, rowimgs := range imgIdxs {
		for j, imgIdx := range rowimgs {
			img = getImage(imgIdx)
			draw.Draw(result, slideWin, img, image.Point{}, draw.Src)
			if j < len(horizontalOverlaps) {
				slideWin.Min.X += slideWin.Dx() - horizontalOverlaps[j]
				slideWin.Max.X += slideWin.Dx() + slideWin.Min.X
			}
		}
		if i < len(verticalOverlaps) {
			slideWin.Min.Y += slideWin.Dy() - verticalOverlaps[i]
			slideWin.Max.Y += slideWin.Dy() + slideWin.Min.Y
			slideWin.Min.X = 0
			slideWin.Max.X = slideWin.Dx()
		}
	}

	return result
}

func MergeUniformBoundsImagesByOverlapReuseMemory(imgIdxs [][]int, getImage func(int) image.Image, imgWidth, imgHeight int,
	horizontalOverlaps, verticalOverlaps []int, result *image.RGBA) {
	if result == nil {
		panic("result is nil")
	}
	if len(result.Pix) == 0 {
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

		// 创建一个新的 RGBA 图片，用于存储合并后的图片
		result = image.NewRGBA(image.Rect(0, 0, resultWidth, resultHeight))
	}
	slideWin := image.Rect(0, 0, imgWidth, imgHeight)
	var img image.Image
	// 将 img1 复制到结果图片中
	for i, rowimgs := range imgIdxs {
		for j, imgIdx := range rowimgs {
			img = getImage(imgIdx)
			draw.Draw(result, slideWin, img, image.Point{}, draw.Src)
			if j < len(horizontalOverlaps) {
				slideWin.Min.X += slideWin.Dx() - horizontalOverlaps[j]
				slideWin.Max.X += slideWin.Dx() + slideWin.Min.X
			}
		}
		if i < len(verticalOverlaps) {
			slideWin.Min.Y += slideWin.Dy() - verticalOverlaps[i]
			slideWin.Max.Y += slideWin.Dy() + slideWin.Min.Y
			slideWin.Min.X = 0
			slideWin.Max.X = slideWin.Dx()
		}
	}
}

type MergeImg struct {
	Pixes                           [][]image.Image
	stride                          int
	effectiveWidth, effectiveHeight []int
	cacheXIdx, cacheYIdx            int
	Rect                            image.Rectangle
}

func (m *MergeImg) ColorModel() color.Model {
	return m.Pixes[0][0].ColorModel()
}

func (m *MergeImg) Bounds() image.Rectangle {
	return m.Rect
}

func (m *MergeImg) ImgOffset(x, y int) image.Image {
	if m.effectiveWidth[m.cacheXIdx] == x {
		m.cacheXIdx += 1
	} else {
		if m.effectiveWidth[m.cacheXIdx] < x {
			m.cacheXIdx = findImgIdx(m.effectiveWidth, m.cacheXIdx+1, len(m.effectiveWidth), x)
		} else if m.cacheXIdx-1 >= 0 && m.effectiveWidth[m.cacheXIdx-1] > x {
			m.cacheXIdx = findImgIdx(m.effectiveWidth, 0, m.cacheXIdx, x)
		}
	}
	if m.effectiveHeight[m.cacheYIdx] == y {
		m.cacheYIdx += 1
	} else {
		if m.effectiveHeight[m.cacheYIdx] < y {
			m.cacheYIdx = findImgIdx(m.effectiveHeight, m.cacheYIdx+1, len(m.effectiveHeight), y)
		} else if m.cacheYIdx-1 >= 0 && m.effectiveHeight[m.cacheYIdx-1] > y {
			m.cacheYIdx = findImgIdx(m.effectiveHeight, 0, m.cacheYIdx, y)
		}
	}
	return m.Pixes[m.cacheYIdx][m.cacheXIdx]
}

func findImgIdx(idx []int, start, end, x int) int {
	for i := start; i < end; i++ {
		if idx[i] > x && (i-1 < 0 || idx[i-1] <= x) {
			return i
		}
	}
	return len(idx) - 1
}

func (m *MergeImg) At(x, y int) color.Color {
	if !(image.Point{X: x, Y: y}.In(m.Rect)) {
		return colori.RGB{}
	}
	pix := m.ImgOffset(x, y)
	if m.cacheXIdx > 0 {
		x -= m.effectiveWidth[m.cacheXIdx-1]
	}
	if m.cacheYIdx > 0 {
		y -= m.effectiveHeight[m.cacheYIdx-1]
	}
	return pix.At(x, y)
}

func NewMergeImg(imgs [][]image.Image, width, height int, horizontalOverlaps, verticalOverlaps []int) *MergeImg {
	effectiveWidth := make([]int, len(imgs[0]))
	effectiveHeight := make([]int, len(imgs))
	var resultWidth, resultHeight int
	for i := range imgs[0] {
		resultWidth += width
		if i < len(horizontalOverlaps) {
			resultWidth -= horizontalOverlaps[i]
		}
		effectiveWidth[i] = resultWidth
	}
	for i := range imgs {
		resultHeight += height
		if i < len(verticalOverlaps) {
			resultHeight -= verticalOverlaps[i]
		}
		effectiveHeight[i] = resultHeight
	}
	return &MergeImg{
		Pixes:           imgs,
		stride:          width * 3,
		effectiveWidth:  effectiveWidth,
		effectiveHeight: effectiveHeight,
		Rect:            image.Rect(0, 0, resultWidth, resultHeight),
	}
}
