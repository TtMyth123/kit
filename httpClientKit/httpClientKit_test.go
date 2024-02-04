package httpClientKit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fatih/structs"
)

func TestDoRequest(t *testing.T) {
	http := GetHttpClient("")
	reader := strings.NewReader("hello")
	str, e := http.DoRequest("POST", "http://www.abc.com", nil, reader)
	fmt.Println(str, e)
}
func TestPostRequestMap(t *testing.T) {
	http := GetHttpClient("")
	type BB struct {
	}
	aBB := BB{}
	mp := requestToMap(aBB)
	str, e := http.PostRequestMap("http://www.abc.com", mp)
	fmt.Println(str, e)
}

func requestToMap(request interface{}) map[string]interface{} {
	return structs.Map(request)
}
