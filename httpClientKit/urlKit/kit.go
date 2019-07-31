package urlKit

func GetUrl(baseUrl string, param map[string]string) string {
	strParam := ""
	for k, v := range param {
		strParam = strParam + "&" + k + "=" + v
	}

	if len(strParam) > 0 {
		strParam = strParam[1:]
	}

	if len(strParam) > 0 {
		baseUrl = baseUrl + "?" + strParam
	}

	return baseUrl
}
