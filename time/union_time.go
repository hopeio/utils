package time

import (
	"time"
)

type EncodeTime struct {
	time.Time
	Encoding
}

func (u EncodeTime) MarshalJSON() ([]byte, error) {
	return u.marshalJSON(u.Time)
}

func (u *EncodeTime) UnmarshalJSON(data []byte) error {
	return u.unmarshalJSON(&u.Time, data)
}
