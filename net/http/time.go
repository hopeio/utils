package http

import (
	timei "github.com/hopeio/utils/time"
	"time"
)

type RequestAt struct {
	Time       time.Time
	TimeStamp  int64
	TimeString string
}

func (r *RequestAt) String() string {
	return r.TimeString
}

func NewRequestAt() *RequestAt {
	now := time.Now()
	return &RequestAt{
		Time:       now,
		TimeStamp:  now.Unix(),
		TimeString: now.Format(timei.LayoutTimeMacro),
	}
}

func NewRequestAtFromTime(t time.Time) *RequestAt {
	return &RequestAt{
		Time:       t,
		TimeStamp:  t.Unix(),
		TimeString: t.Format(timei.LayoutTimeMacro),
	}
}
