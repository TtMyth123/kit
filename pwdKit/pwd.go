package pwdKit

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

/*
*
奖pwd字符串进行Sha1加密并转换成string
*/
func Sha1ToStr(pwd string) string {
	h := sha1.New()
	io.WriteString(h, pwd)
	newPwd := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return newPwd
}

/*
*
奖pwd字符串进行Sha256加密并转换成string
*/
func Sha256ToStr(pwd string) string {
	h := sha256.New()
	io.WriteString(h, pwd)
	newPwd := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return newPwd
}
func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

/*
*
奖pwd字符串进行Md5加密并转换成string
*/
func Md5ToStr(pwd string) string {
	h := md5.New()
	io.WriteString(h, pwd)
	newPwd := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return newPwd
}

/*
*
标准将pwd字符串进行Md5加密并转换成string
*/
func StdMd5ToStr(pwd string) string {
	md5data := md5.Sum([]byte(pwd))
	str := fmt.Sprintf("%x", md5data)
	return str
}

/*
*
给str编码，并加密
*/
func TtEncodeToString(str, pwd string) string {
	a := base64.StdEncoding.EncodeToString([]byte(pwd + str))
	return a
}

/*
*
给str解码并 解密
*/
func TtDecodeString(str, pwd string) (string, error) {
	bb, e := base64.StdEncoding.DecodeString(str)
	if e != nil {
		return "", e
	}
	i := len(pwd)
	str1 := string(bb)
	if len(str1) >= i {
		return str1[i:], nil
	}
	return "", errors.New("秘钥不对")

}
