/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package buffer

import "go.uber.org/zap/buffer"

var (
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = buffer.NewPool().Get
)
