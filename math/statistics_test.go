package math

import (
	"fmt"
	"testing"
)

func TestMean(t *testing.T) {
	data := []float64{1, 2, 2, 2, 3, 10, 2, 2, 1, 2, 3, 2, 100}
	fmt.Printf("Original data: %v\n", data)
	result := removeOutliersAndMean(data)
	fmt.Printf("Mean after removing outliers: %f\n", result)
}
