/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package pretty

import "github.com/tidwall/pretty"

func Pretty(json []byte) []byte { return pretty.PrettyOptions(json, nil) }
func PrettyOptions(json []byte, options *pretty.Options) []byte {
	return pretty.PrettyOptions(json, options)
}
