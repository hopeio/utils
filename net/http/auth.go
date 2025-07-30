/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package http

import (
	"encoding/base64"
	"github.com/hopeio/gox/net/http/consts"
	"net/http"
)

func SetBasicAuth(header http.Header, username, password string) {
	header.Set(consts.HeaderAuthorization, "Basic "+BasicAuth(username, password))
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func GetToken(r *http.Request) string {
	if token := r.Header.Get(consts.HeaderAuthorization); token != "" {
		return token
	}
	if cookie, _ := r.Cookie(consts.HeaderCookieValueToken); cookie != nil {
		return cookie.Value
	}
	return ""
}
