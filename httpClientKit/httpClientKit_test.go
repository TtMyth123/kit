package httpClientKit

func TestDoRequest(t *testing.T) {
	http := GetHttpClient("")
	reader := strings.NewReader("hello")
	str, e := http.DoRequest("POST", "http://www.abc.com", nil, reader)
	fmt.Println(str, e)
}
