/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package sjson

import "github.com/tidwall/sjson"

func Set(json, path string, value interface{}) (string, error) {
	return sjson.SetOptions(json, path, value, nil)
}
