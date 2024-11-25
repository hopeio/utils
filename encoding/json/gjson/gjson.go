/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package gjson

import "github.com/tidwall/gjson"

func Get(json, path string) gjson.Result {
	return gjson.Get(json, path)
}
