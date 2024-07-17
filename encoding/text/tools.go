package text

import (
	"encoding"
	"golang.org/x/net/html/charset"
	tencoding "golang.org/x/text/encoding"
)

func DetermineEncoding(content []byte, contentType string) (e tencoding.Encoding, name string, certain bool) {
	return charset.DetermineEncoding(content, contentType)
}

func Unmarshal[T any](str string) error {
	var t T
	v, vp := any(t), any(&t)
	itv, ok := v.(encoding.TextUnmarshaler)
	if !ok {
		itv, ok = vp.(encoding.TextUnmarshaler)
	}
	if ok {
		err := itv.UnmarshalText([]byte(str))
		if err != nil {
			return err
		}
	}
	return nil
}
