package math

import (
	"math"
	"sort"
)

// Calculate the Median of a slice of floats
func Median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

// Calculate the Mean of a slice of floats
func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Remove outliers using the MAD method and calculate the Mean of the remaining data
func RemoveOutliersAndMean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	med := Median(data)

	// Calculate absolute deviations from the Median
	absDevs := make([]float64, len(data))
	for i, value := range data {
		absDevs[i] = math.Abs(value - med)
	}

	// Calculate the Median of the absolute deviations
	mad := Median(absDevs)

	// Define a threshold using the MAD; here we use 3 times the MAD
	threshold := 3.0 * mad

	// Filter out outliers
	filteredData := make([]float64, 0)
	for _, value := range data {
		if math.Abs(value-med) <= threshold {
			filteredData = append(filteredData, value)
		}
	}

	// Calculate the Mean of the remaining data
	return Mean(filteredData)
}
