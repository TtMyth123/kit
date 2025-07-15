package pwdKit

import (
	"fmt"
	"testing"
)

func TestCreateFile(t *testing.T) {

	fmt.Println()
}
func TestMd5ToStr(t *testing.T) {
	str := "a"
	str1 := Md5ToStr(str)
	fmt.Println(str1, str)
}
