/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package geometry

import (
	"golang.org/x/exp/constraints"
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
	return math.Hypot(dx, dy)
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

// PointInt 结构体用于表示一个点
type PointInt[T constraints.Integer] struct {
	X T
	Y T
}

func (l *PointInt[T]) ToFloat64(factor float64) *Point {
	return &Point{
		X: float64(l.X) / factor,
		Y: float64(l.Y) / factor,
	}
}

type Point3DInt[T constraints.Integer] struct {
	X T
	Y T
	Z T
}

func (l *Point3DInt[T]) ToFloat64(factor float64) *Point3D {
	return &Point3D{
		X: float64(l.X) / factor,
		Y: float64(l.Y) / factor,
		Z: float64(l.Z) / factor,
	}
}

type VectorInt[T constraints.Integer] struct {
	X T
	Y T
}

func (l *VectorInt[T]) ToFloat64(factor float64) *Vector {
	return &Vector{
		X: float64(l.X) / factor,
		Y: float64(l.Y) / factor,
	}
}
