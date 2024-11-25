//go:build jsoniter

/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	Marshal = json.Marshal

	Unmarshal = json.Unmarshal

	MarshalIndent = json.MarshalIndent

	NewDecoder = json.NewDecoder

	NewEncoder = json.NewEncoder

	MarshalToString     = json.MarshalToString
	UnmarshalFromString = json.UnmarshalFromString
)

var Standard = jsoniter.ConfigCompatibleWithStandardLibrary

func SupportPrivateFields() {
	extra.SupportPrivateFields()
}

var WithPrivateField = jsoniter.Config{
	IndentionStep:                 4,
	MarshalFloatWith6Digits:       true,
	EscapeHTML:                    true,
	SortMapKeys:                   true,
	UseNumber:                     true,
	ObjectFieldMustBeSimpleString: true,
}.Froze()
