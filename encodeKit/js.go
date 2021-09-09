package encodeKit

import
(
"net/url"
"strings"
)
func JsEncodeURI(strUrl string) string {
	newUrl := url.PathEscape(strUrl)
	newUrl = strings.Replace(newUrl, "%2F", `/`, 100)
	newUrl = strings.Replace(newUrl, "%3F", `?`, 100)
	return newUrl
}
