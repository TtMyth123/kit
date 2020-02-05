package requestKit

import (
	"net/http"
	"strconv"
	"strings"
)

/**
判断 是不是手机请求
 */
func IsMobileDevice(Request *http.Request) bool {
	userAgent := Request.Header.Get("user-agent")
	userAgent = strings.ToLower(userAgent)
	if userAgent == "" {
		return false
	}
	deviceArray := []string{"android", "mac os", "windows phone"}
	for _, v := range deviceArray {
		if strings.Index(userAgent, v) > 0 {
			return true
		}
	}
	return false
}

/**
获取请求的 IP和端口
 */
func GetIPInfo(Request *http.Request) (string, int) {
	remoteAddr := Request.RemoteAddr
	ipInfo := strings.Split(remoteAddr, ":")
	if len(ipInfo) == 2 {
		port, _ := strconv.Atoi(ipInfo[1])
		return ipInfo[0], port
	}

	return remoteAddr, 0
}
