/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package fiber

import (
	httpi "github.com/hopeio/utils/net/http"
	stringsi "github.com/hopeio/utils/strings"
	"github.com/valyala/fasthttp"
	"net/url"
)

func GetToken(req *fasthttp.Request) string {
	if token := stringsi.BytesToString(req.Header.Peek(httpi.HeaderAuthorization)); token != "" {
		return token
	}
	if cookie := stringsi.BytesToString(req.Header.Cookie(httpi.HeaderCookieValueToken)); len(cookie) > 0 {
		token, _ := url.QueryUnescape(cookie)
		return token
	}
	return ""
}
