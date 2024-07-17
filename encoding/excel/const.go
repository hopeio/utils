package excel

import "github.com/xuri/excelize/v2"

type ColumnNumber int

const (
	A ColumnNumber = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	AA
	AB
	AC
)

// 只拓展到两位列ZZ
func (c ColumnNumber) Sting() string {
	if c < 26 {
		return string(rune(c + 'A'))
	}

	return (c/26 - 1).Sting() + (c % 26).Sting()
}

var ColumnLetter = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC"}

var Style = &excelize.Style{
	Alignment: &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	},
}

var HeaderStyle = &excelize.Style{
	Border: []excelize.Border{{
		Type:  "left",
		Color: "000000",
		Style: 1,
	},
		{
			Type:  "top",
			Color: "000000",
			Style: 1,
		},
		{
			Type:  "bottom",
			Color: "000000",
			Style: 1,
		},
		{
			Type:  "right",
			Color: "000000",
			Style: 1,
		}},
	Fill: excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#bfbfbf"},
	},
	Font: &excelize.Font{Bold: true},
	Alignment: &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	},
	Protection:   nil,
	NumFmt:       0,
	CustomNumFmt: nil,
	NegRed:       false,
}
