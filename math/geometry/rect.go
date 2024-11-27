/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package geometry

import (
	mathi "github.com/hopeio/utils/math"
	"golang.org/x/exp/constraints"
	"image"
	"math"
)

type Triangle struct {
	A, B, C Point
}

type Rectangle struct {
	CenterX float64
	CenterY float64
	Width   float64
	Height  float64
	Angle   float64
}

func NewRect(centerX, centerY, width, height float64, angleDeg float64) *Rectangle {
	return &Rectangle{
		CenterX: centerX,
		CenterY: centerY,
		Width:   width,
		Height:  height,
		Angle:   angleDeg,
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
		CenterX: (x0 + x1) / 2,
		CenterY: (y0 + y1) / 2,
		Width:   x1 - x0,
		Height:  y1 - y0,
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
	minx, maxx := mathi.MinAndMax(corners[0].X, corners[1].X, corners[2].X, corners[3].X)
	miny, maxy := mathi.MinAndMax(corners[0].Y, corners[1].Y, corners[2].Y, corners[3].Y)
	return RectNoRotate(minx, miny, maxx, maxy)
}

func (rect *Rectangle) Corners() [4]Point {
	if rect.Angle == 0 {
		return [4]Point{{rect.CenterX - rect.Width/2, rect.CenterY - rect.Height/2},
			{rect.CenterX + rect.Width/2, rect.CenterY - rect.Height/2},
			{rect.CenterX + rect.Width/2, rect.CenterY + rect.Height/2},
			{rect.CenterX - rect.Width/2, rect.CenterY + rect.Height/2}}
	}
	angleRad := rect.Angle * math.Pi / 180.0
	// Calculate cosine and sine of the angle
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	halfW, halfH := rect.Width/2, rect.Height/2
	// 计算矩形四个角的坐标 (A左下-B右下-C右上-D左上)
	dx := rect.CenterX + halfW*cosA - halfH*sinA
	dy := rect.CenterY + halfW*sinA + halfH*cosA
	ax := rect.CenterX - halfW*cosA - halfH*sinA
	ay := rect.CenterY - halfW*sinA + halfH*cosA
	bx := rect.CenterX - halfW*cosA + halfH*sinA
	by := rect.CenterY - halfW*sinA - halfH*cosA
	cx := rect.CenterX + halfW*cosA + halfH*sinA
	cy := rect.CenterY + halfW*sinA - halfH*cosA
	return [4]Point{{ax, ay}, {bx, by}, {cx, cy}, {dx, dy}}
}

// 图片就是第四象限,角度90+θ
func (rect *Rectangle) ContainsPoint(p Point) bool {

	// 射线法判断点是否在矩形内
	inside := false
	intersections := 0
	corners := rect.Corners()

	for i := 0; i < len(corners); i++ {
		x1, y1 := corners[i].X, corners[i].Y
		x2, y2 := corners[(i+1)%len(corners)].X, corners[(i+1)%len(corners)].Y

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

type RectangleInt[T constraints.Integer] struct {
	Center PointInt[T]
	Width  T
	Height T
	Angle  float64
}

func (rect *RectangleInt[T]) ToFloat64(factor float64) *Rectangle {
	if factor == 0 {
		factor = 1
	}
	return &Rectangle{
		CenterX: float64(rect.Center.X) / factor,
		CenterY: float64(rect.Center.Y) / factor,
		Width:   float64(rect.Width) / factor,
		Height:  float64(rect.Height) / factor,
		Angle:   rect.Angle,
	}
}

func NewRectInt[T constraints.Integer](center PointInt[T], width, height T, angle float64) *RectangleInt[T] {
	return &RectangleInt[T]{center, width, height, angle}
}

func RectIntFromFloat64[T constraints.Integer](e *Rectangle, factor float64) *RectangleInt[T] {
	if factor == 0 {
		factor = 1
	}
	return &RectangleInt[T]{
		Center: PointInt[T]{
			X: T(math.Round(e.CenterX * factor)),
			Y: T(math.Round(e.CenterY * factor)),
		},
		Width:  T(math.Round(e.Width * factor)),
		Height: T(math.Round(e.Angle * factor)),
		Angle:  e.Angle,
	}
}
