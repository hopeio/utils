/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package math

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	fmt.Printf("%.2f", 123.123)
}

func TestDecimalPlaces(t *testing.T) {
	assert.Equal(t, 3.14, DecimalPlaces(3.1415926, 2))
	assert.Equal(t, 3.141, DecimalPlaces(3.1415926, 3))
	assert.Equal(t, 3.1415, DecimalPlaces(3.1415926, 4))
	assert.Equal(t, 3.14159, DecimalPlaces(3.1415926, 5))

	assert.Equal(t, 3.14, DecimalPlacesRound(3.1415926, 2))
	assert.Equal(t, 3.142, DecimalPlacesRound(3.1415926, 3))
	assert.Equal(t, 3.1416, DecimalPlacesRound(3.1415926, 4))
	assert.Equal(t, 3.14159, DecimalPlacesRound(3.1415926, 5))

}
