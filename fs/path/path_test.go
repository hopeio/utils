/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package path

import (
	"github.com/stretchr/testify/assert"
	"path"
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	dir := "https://a\\video/a.jpg"
	t.Log(path.Split(dir))
	t.Log(filepath.Split(dir))
	t.Log(filepath.Dir(dir), filepath.Base(dir))
}

func TestClean(t *testing.T) {
	assert.Equal(t, "", FileCleanse(`--......++`))
	assert.Equal(t, "1...1", FileCleanse(`--1...1...++`))
}

func TestRune(t *testing.T) {
	t.Log('，')
}
