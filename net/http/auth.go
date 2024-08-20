package http

import (
	"encoding/base64"
	"net/http"
	"net/url"
)

func SetBasicAuth(header http.Header, username, password string) {
	header.Set(HeaderAuthorization, "Basic "+BasicAuth(username, password))
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

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
