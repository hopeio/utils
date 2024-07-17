package math

import (
	"math"
)

func DecimalPlaces(value float64, prec int) float64 {
	multiplier := math.Pow(10, float64(prec))
	return float64(int(value*multiplier)) / multiplier
}

// 四舍五入
func DecimalPlacesRound(value float64, rank int) float64 {
	multiplier := math.Pow(10, float64(rank))
	return math.Round(value*multiplier) / multiplier
}
