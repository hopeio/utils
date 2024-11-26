/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package geometry

import (
	mathi "github.com/hopeio/utils/math"
	"image"
	"math"
)

type Triangle struct {
	A, B, C Point
}

type Rectangle struct {
	Center Point
	Width  float64
	Height float64
	Angle  float64
}

func NewRect(center Point, width, height float64, angleDeg float64) *Rectangle {
	return &Rectangle{
		Center: center,
		Width:  width,
		Height: height,
		Angle:  angleDeg,
	}
}

func RectNoRotate(x0, y0, x1, y1 float64) *Rectangle {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return &Rectangle{
		Center: Point{(x0 + x1) / 2, (y0 + y1) / 2},
		Width:  x1 - x0,
		Height: y1 - y0,
	}
}

func RectFromImageRect(r image.Rectangle) *Rectangle {
	return RectNoRotate(float64(r.Min.X), float64(r.Min.Y), float64(r.Max.X), float64(r.Max.Y))
}

func (rect *Rectangle) Bounds() *Rectangle {
	if rect.Angle == 0 {
		return &*rect
	}
	corners := rect.Corners()
	minx, maxx := mathi.MinAndMax(corners[0][0], corners[1][0], corners[2][0], corners[3][0])
	miny, maxy := mathi.MinAndMax(corners[0][1], corners[1][1], corners[2][1], corners[3][1])
	return RectNoRotate(minx, miny, maxx, maxy)
}

func (rect *Rectangle) Corners() [][]float64 {
	angleRad := rect.Angle * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	halfW, halfH := rect.Width/2, rect.Height/2
	// 计算矩形四个角的坐标 (A左下-B右下-C右上-D左上)
	dx := rect.Center.X + halfW*cosA - halfH*sinA
	dy := rect.Center.Y + halfW*sinA + halfH*cosA
	ax := rect.Center.X - halfW*cosA - halfH*sinA
	ay := rect.Center.Y - halfW*sinA + halfH*cosA
	bx := rect.Center.X - halfW*cosA + halfH*sinA
	by := rect.Center.Y - halfW*sinA - halfH*cosA
	cx := rect.Center.X + halfW*cosA + halfH*sinA
	cy := rect.Center.Y + halfW*sinA - halfH*cosA
	return [][]float64{{ax, ay}, {bx, by}, {cx, cy}, {dx, dy}}
}

// 图片就是第四象限,角度90+θ
func (rect *Rectangle) ContainsPoint(p Point) bool {

	// 射线法判断点是否在矩形内
	inside := false
	intersections := 0
	corners := rect.Corners()

	for i := 0; i < len(corners); i++ {
		x1, y1 := corners[i][0], corners[i][1]
		x2, y2 := corners[(i+1)%len(corners)][0], corners[(i+1)%len(corners)][1]

		if y1 == y2 { // 水平边
			continue
		}
		if p.Y < min(y1, y2) || p.Y > max(y1, y2) { // 在边的外部
			continue
		}

		xIntersect := x1 + (p.Y-y1)*(x2-x1)/(y2-y1)
		if p.X < xIntersect {
			intersections++
		}
	}

	if intersections%2 == 1 {
		inside = true
	}

	return inside
}
