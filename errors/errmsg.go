/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package errors

type ErrMsg string

func (e ErrMsg) Error() string {
	return string(e)
}
