package pwdKit

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/big"

	"crypto/rand"
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

// StdMd5ToStr /*
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

/*
*
1.25 go 不支持1024位RSA密钥 所以添加支持1024位RSA方法
*/
func RSAEncryptWith1024Key(plainText string, pubKeyPEM string) (string, error) {
	// 解析 PEM
	block, _ := pem.Decode([]byte(pubKeyPEM))
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block")
	}

	// 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse public key: %v", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("not an RSA public key")
	}

	fmt.Printf("使用 %d 位 RSA 密钥进行加密\n", rsaPub.N.BitLen())

	// 使用自定义加密函数
	ciphertext, err := customRSAEncryptPKCS1v15(rsaPub, []byte(plainText))
	if err != nil {
		return "", fmt.Errorf("encryption failed: %v", err)
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
func customRSAEncryptPKCS1v15(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if len(msg) > k-11 {
		return nil, fmt.Errorf("message too long")
	}

	// 构造 EM = 0x00 || 0x02 || PS || 0x00 || M
	em := make([]byte, k)
	em[0] = 0x00
	em[1] = 0x02

	// 生成随机填充
	psLen := k - len(msg) - 3
	_, err := rand.Read(em[2 : 2+psLen])
	if err != nil {
		return nil, err
	}

	// 设置分隔符
	em[2+psLen] = 0x00

	// 复制消息
	copy(em[3+psLen:], msg)

	// 将 EM 转换为大整数
	m := new(big.Int).SetBytes(em)

	// 计算密文: c = m^e mod n
	c := new(big.Int).Exp(m, big.NewInt(int64(pub.E)), pub.N)

	// 将密文转换回字节数组
	ciphertext := c.Bytes()

	// 补齐到 k 字节
	if len(ciphertext) < k {
		padded := make([]byte, k)
		copy(padded[k-len(ciphertext):], ciphertext)
		ciphertext = padded
	}

	return ciphertext, nil
}
