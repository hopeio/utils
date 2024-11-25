/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package time

import (
	"time"
)

type EncodingTime struct {
	time.Time
	Encoding
}

func (u EncodingTime) MarshalJSON() ([]byte, error) {
	return u.marshalJSON(u.Time)
}

func (u *EncodingTime) UnmarshalJSON(data []byte) error {
	return u.unmarshalJSON(&u.Time, data)
}

type GlobETime time.Time

func (u GlobETime) MarshalJSON() ([]byte, error) {
	return encoding.marshalJSON(time.Time(u))
}

func (u *GlobETime) UnmarshalJSON(data []byte) error {
	return encoding.unmarshalJSON((*time.Time)(u), data)
}
