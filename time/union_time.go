package time

import (
	"time"
)

type UnionTime struct {
	time.Time
	Encoding
}

func (u UnionTime) MarshalJSON() ([]byte, error) {
	return u.marshalJSON(u.Time)
}

func (u *UnionTime) UnmarshalJSON(data []byte) error {
	return u.unmarshalJSON(&u.Time, data)
}
