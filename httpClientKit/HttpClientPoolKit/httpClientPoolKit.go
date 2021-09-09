package HttpClientPoolKit

import (
	"net/http"
	"net/http/cookiejar"
)

type HttpClientPool struct {
	client        *http.Client

	gCurCookieJar *cookiejar.Jar
	gCurCookies   []*http.Cookie
}