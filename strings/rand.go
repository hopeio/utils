/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package strings

import "math/rand/v2"

func String() string {
	runes := make([]rune, 5)
	for i := range runes {
		runes[i] = Rune()
	}
	return string(runes)
}

func Rune() rune {
	r := rand.N(500)
	return rune(r + 19968)
}
