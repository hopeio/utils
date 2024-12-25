/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package fiber

import (
	"github.com/gofiber/fiber/v3"
	httpi "github.com/hopeio/utils/net/http"
	stringsi "github.com/hopeio/utils/strings"
	"net/url"
)

func GetToken(ctx fiber.Ctx) string {
	req := ctx.Request()
	if token := stringsi.BytesToString(req.Header.Peek(httpi.HeaderAuthorization)); token != "" {
		return token
	}
	if cookie := stringsi.BytesToString(req.Header.Cookie("token")); len(cookie) > 0 {
		token, _ := url.QueryUnescape(cookie)
		return token
	}
	return ""
}
