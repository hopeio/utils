/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package msgpack

import "github.com/vmihailenco/msgpack/v5"

func Marshal(v any) ([]byte, error) {
	return msgpack.Marshal(v)
}
