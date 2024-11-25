/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package interfaces

type TaskRetry interface {
	Do(times uint) (retry bool)
}