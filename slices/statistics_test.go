/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package slices

import (
	"fmt"
	"testing"
)

func TestMean(t *testing.T) {
	data := []float64{1, 2, 2, 2, 3, 10, 2, 2, 1, 2, 3, 2, 100}
	fmt.Printf("Original data: %v\n", data)
	result := RemoveOutliersMean(data)
	fmt.Printf("Mean after removing outliers: %f\n", result)
}
