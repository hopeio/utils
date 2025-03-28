/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package log

import (
	"cmp"
	"time"
)

func ValueNotify[T cmp.Ordered](msg string, v T, rangeMin, rangeMax T) {
	if v > rangeMin || v < rangeMax {
		CallerSkipLogger(1).Warnf("%s except: %v - %v,but got %s", msg, rangeMin, rangeMax, v)
	}
}

func DurationNotify(msg string, v time.Duration, std time.Duration) {
	if v > 0 && v < std {
		CallerSkipLogger(1).Warnf("%s except: %s level,but got %s", msg, std, v)
	}
}
