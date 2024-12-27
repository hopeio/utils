/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package text

import (
	"encoding"
	reflecti "github.com/hopeio/utils/reflect/converter"
	stringsi "github.com/hopeio/utils/strings"
	"golang.org/x/net/html/charset"
	tencoding "golang.org/x/text/encoding"
	"strconv"
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

func StringConvertFor[T any](str string) (T, error) {
	var t T
	a, ap := any(t), any(&t)
	itv, ok := a.(encoding.TextUnmarshaler)
	if !ok {
		itv, ok = ap.(encoding.TextUnmarshaler)
	}
	if ok {
		var err error
		str, err = strconv.Unquote(str)
		err = itv.UnmarshalText(stringsi.ToBytes(str))
		if err != nil {
			return t, err
		}
		return t, nil
	}

	v, err := reflecti.StringConvertFor[T](str)
	if err != nil {
		return t, err
	}
	return v, nil
}
