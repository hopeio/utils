/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package svg

import (
	"fmt"
)

type Path struct {
	X        int
	Y        int
	Commands []interface{}
	Fill     string
	Attr     map[string]string
}

type PathArc struct {
	RadiusX  int
	RadiusY  int
	LargeArc int
	Sweep    int
	X        int
	Y        int

	CenterX int
	CenterY int
}

type PathLine struct {
	X int
	Y int
}

type Circle struct {
	X      int
	Y      int
	Radius int
	Fill   string
	Attr   map[string]string
}

type Line struct {
	XStart int
	YStart int
	XEnd   int
	YEnd   int
	Width  int
	Stroke string
	Attr   map[string]string
}

type Rectangle struct {
	X        int
	Y        int
	Width    int
	Height   int
	XCenter  int // rotation center
	YCenter  int
	Rotation float64
	RX       int
	RY       int
	Fill     string
	Attr     map[string]string
}

func FormatAttr(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}

	s := ""
	for k, v := range m {
		s += fmt.Sprintf(` %s="%s"`, k, v)
	}
	return s
}
