package excel

import "github.com/xuri/excelize/v2"

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

// xlsx cell style
var (
	noteCellStyle = &excelize.Style{
		Font: &excelize.Font{
			Family: "Arial",
			Bold:   true,
			Size:   10,
			Color:  "#FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#4F6228"},
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
			WrapText:   true,
		},
	}
	headerCellStyle = &excelize.Style{
		Font: &excelize.Font{
			Family: "Arial",
			Bold:   true,
			Size:   10,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#D3D3D3"},
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	}
	bodyCellStyle = &excelize.Style{
		Font: &excelize.Font{
			Family: "Arial",
			Bold:   false,
			Size:   10,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	}
	errorBodyCellStyle = &excelize.Style{
		Font: &excelize.Font{
			Family: "Arial",
			Bold:   false,
			Size:   10,
			Color:  "#FF0000",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
			WrapText:   true,
		},
	}
)
