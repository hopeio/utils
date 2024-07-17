package math

import (
	"math"
	"sort"
)

// Calculate the median of a slice of floats
func median(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

// Calculate the mean of a slice of floats
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// Remove outliers using the MAD method and calculate the mean of the remaining data
func removeOutliersAndMean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	med := median(data)

	// Calculate absolute deviations from the median
	absDevs := make([]float64, len(data))
	for i, value := range data {
		absDevs[i] = math.Abs(value - med)
	}

	// Calculate the median of the absolute deviations
	mad := median(absDevs)

	// Define a threshold using the MAD; here we use 3 times the MAD
	threshold := 3.0 * mad

	// Filter out outliers
	filteredData := make([]float64, 0)
	for _, value := range data {
		if math.Abs(value-med) <= threshold {
			filteredData = append(filteredData, value)
		}
	}

	// Calculate the mean of the remaining data
	return mean(filteredData)
}
