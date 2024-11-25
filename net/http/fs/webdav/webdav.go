/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package webdav

import (
	"golang.org/x/net/webdav"
	"net/http"
)

func WebDav(prefix string, root webdav.FileSystem, lock webdav.LockSystem, logger func(*http.Request, error)) http.Handler {
	return &webdav.Handler{
		Prefix:     prefix,
		FileSystem: root,
		LockSystem: lock,
		Logger:     logger,
	}
}
