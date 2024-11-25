/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package style

import (
	"fmt"
	"testing"
)

func TestCustom(t *testing.T) {
	fmt.Println(Custom("红色", 31, 39))
	fmt.Println(Custom("红色", 47, 0))
	fmt.Println(Custom("红色", DcItalic, 0))
	fmt.Println(Custom("红色", DcUnderline, 0))
	for i := range 256 {
		fmt.Println(Color256("红色", byte(i)))
	}
}
