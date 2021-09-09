package HttpClientPoolKit

import (
	"net/http"
	"net/http/cookiejar"
)

type HttpClientData struct {
	CurCookieJar *cookiejar.Jar
	Cookies []*http.Cookie
}

func NewHttpClientData()  {
	aHttpClientData := new(HttpClientData)
	aHttpClientData.CurCookieJar, _ = cookiejar.New(nil)
}