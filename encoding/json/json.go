// Copyright 2017 Bo-Yi Wu. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !jsoniter && !(sonic && amd64) && !go_json

package json

import (
	"encoding/json"
	"github.com/hopeio/utils/strings"
)

var (
	Marshal = json.Marshal

	Unmarshal = json.Unmarshal

	MarshalIndent = json.MarshalIndent

	NewDecoder = json.NewDecoder

	NewEncoder = json.NewEncoder
)

func MarshalToString(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return strings.FromBytes(data), nil
}

func UnmarshalFromString(str string, v any) error {
	return json.Unmarshal(strings.ToBytes(str), v)
}
