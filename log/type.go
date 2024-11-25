/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package log

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
)

type RawJson json.RawMessage

func (b RawJson) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return nil
}

func (b RawJson) MarshalJSON() ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}
	return b, nil
}
func (b *RawJson) UnmarshalJSON(raw []byte) error {
	*b = raw
	return nil
}
