/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package fiber

import (
	"github.com/hopeio/utils/net/http/consts"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/valyala/fasthttp"
	"net/url"
)

func GetToken(req *fasthttp.Request) string {
	if token := stringsi.BytesToString(req.Header.Peek(consts.HeaderAuthorization)); token != "" {
		return token
	}
	if cookie := stringsi.BytesToString(req.Header.Cookie(consts.HeaderCookieValueToken)); len(cookie) > 0 {
		token, _ := url.QueryUnescape(cookie)
		return token
	}
	return ""
}
