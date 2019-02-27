package pwdKit

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

/**
奖pwd字符串进行Sha1加密并转换成string
*/
func Sha1ToStr(pwd string) string {
	h := sha1.New()
	io.WriteString(h, pwd)
	newPwd := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return newPwd
}

/**
奖pwd字符串进行Sha256加密并转换成string
*/
func Sha256ToStr(pwd string) string {
	h := sha256.New()
	io.WriteString(h, pwd)
	newPwd := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return newPwd
}

/**
奖pwd字符串进行Md5加密并转换成string
*/
func Md5ToStr(pwd string) string {
	h := md5.New()
	io.WriteString(h, pwd)
	newPwd := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return newPwd
}
