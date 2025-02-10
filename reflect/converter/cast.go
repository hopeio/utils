package converter

import "github.com/spf13/cast"

func CastInt64(v any) int64 {
	return cast.ToInt64(v)
}
