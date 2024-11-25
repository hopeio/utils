/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import "testing"

func TestFetch(t *testing.T) {
	_, err := GetReader("")
	if err != nil {
		t.Log(err)
	}
}
