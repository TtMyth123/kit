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

func Test_RSAEncryptWith1024Key(t *testing.T) {
	PublicKey := `-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgH1ufjmfY7tAt14i4z1ba8HvjCEf
HeGdG8+9ABVbZhPphrW15dyb/Q6HfTCyBwDLdGZViiPcXli3/guMfv83o6nLZPyz
gg8f7fblKF9vB5l9AmuzKSVL/8DCBaIXIr1aRmMBV04f1v4xXPyUdNXQkzDkXobX
wPpk3jyb/K/XKefJAgMBAAE=
-----END PUBLIC KEY-----
`
	password := "a"

	encrypted, err := RSAEncryptWith1024Key(password, PublicKey)
	if err != nil {
		fmt.Printf("加密失败: %v\n", err)
		return
	}

	fmt.Println("加密结果:", encrypted)

}
