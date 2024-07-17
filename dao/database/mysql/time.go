package mysql

import (
	timei "github.com/hopeio/utils/time"
	"time"
)

func Now() string {
	return time.Now().Format(timei.LayoutTimeMacro)
}
