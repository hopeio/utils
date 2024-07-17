package http

import (
	"net/http"
	"net/url"
)

func GetToken(r *http.Request) string {
	if token := r.Header.Get(HeaderAuthorization); token != "" {
		return token
	}
	if cookie, _ := r.Cookie(HeaderCookieValueToken); cookie != nil {
		value, _ := url.QueryUnescape(cookie.Value)
		return value
	}
	return ""
}
