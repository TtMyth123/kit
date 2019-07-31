package kit

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strconv"

	"encoding/json"
)

const (
	//BASE64字符表,不要有重复
	base64Table        = "<>:;',./?~!@#$CDVWX%^&*ABYZabcghijklmnopqrstuvwxyz01EFGHIJKLMNOPQRSTU2345678(def)_+|{}[]9/"
	hashFunctionHeader = "zh.ife.iya"
	hashFunctionFooter = "09.O25.O20.78"
)

const (
	PT_Goods         = ""
	PT_WJ            = ""
	PT_LoginGift     = ""
	PT_MallGoods     = ""
	PT_TurnplateGift = ""
	PT_MailGift      = ""
)

func StrToBytes(strData string) []byte {
	buffer := &bytes.Buffer{}
	buffer.WriteString(strData)
	return buffer.Bytes()
}

func BytesToStr(b []byte) string {
	buffer := &bytes.Buffer{}
	buffer.Write(b)
	return buffer.String()
}

func GetMapStr(mp map[string]interface{}, k string) string {
	r := ""
	if mp[k] != nil {
		if r1, ok := mp[k].(string); ok {
			return r1
		}
	}

	return r
}

func GetInterface2Str(mp interface{}, k string) string {
	r := k
	if mp != nil {
		if r1, ok := mp.(string); ok {
			return r1
		}
	}

	return r
}

func GetInterface2Int(mp interface{}, k int) int {
	if r1, ok := mp.(int); ok {
		return r1
	}
	if r1, ok := mp.(int32); ok {
		return int(r1)
	}
	if r1, ok := mp.(int64); ok {
		return int(r1)
	}
	if r1, ok := mp.(float64); ok {
		return int(r1)
	}
	if r1, ok := mp.(string); ok {
		r, e := strconv.Atoi(r1)
		if e == nil {
			return r
		} else {
			return k
		}
	}

	return k
}

func GetInterface2Int64(mp interface{}, k int64) int64 {
	if r1, ok := mp.(int64); ok {
		return r1
	}
	if r1, ok := mp.(int); ok {
		return int64(r1)
	}
	if r1, ok := mp.(int32); ok {
		return int64(r1)
	}
	if r1, ok := mp.(float64); ok {
		return int64(r1)
	}

	if r1, ok := mp.(string); ok {
		r, e := strconv.ParseInt(r1, 10, 64)
		if e == nil {
			return int64(r)
		} else {
			return k
		}
	}

	return k
}

func GetInterface2Int32(mp interface{}, k int32) int32 {
	if r1, ok := mp.(int32); ok {
		return r1
	}
	if r1, ok := mp.(int); ok {
		return int32(r1)
	}
	if r1, ok := mp.(int64); ok {
		return int32(r1)
	}
	if r1, ok := mp.(float64); ok {
		return int32(r1)
	}

	if r1, ok := mp.(string); ok {
		r, e := strconv.ParseInt(r1, 10, 64)
		if e == nil {
			return int32(r)
		} else {
			return k
		}
	}

	return k
}

func GetInterface2Float64(mp interface{}, k float64) float64 {
	if r1, ok := mp.(float64); ok {
		return r1
	}
	if r1, ok := mp.(float32); ok {
		return float64(r1)
	}
	if r1, ok := mp.(int); ok {
		return float64(r1)
	}
	if r1, ok := mp.(int64); ok {
		return float64(r1)
	}
	if r1, ok := mp.(float64); ok {
		return float64(r1)
	}
	if r1, ok := mp.(int); ok {
		return float64(r1)
	}
	if r1, ok := mp.(int32); ok {
		return float64(r1)
	}

	if r1, ok := mp.(string); ok {
		r, e := strconv.ParseFloat(r1, 64)
		if e == nil {
			return float64(r)
		} else {
			return k
		}
	}

	return k
}
func GetInterface2Bool(mp interface{}, k bool) bool {
	if r1, ok := mp.(bool); ok {
		return r1
	}
	if r1, ok := mp.(int); ok {
		return r1 == 1
	}
	if r1, ok := mp.(int64); ok {
		return r1 == 1
	}
	if r1, ok := mp.(float64); ok {
		return r1 == 1
	}

	if r1, ok := mp.(string); ok {
		r, e := strconv.ParseBool(r1)
		if e == nil {
			return r
		} else {
			return k
		}
	}

	return k
}

/**
 * 获取一个Guid值
 */
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s)) //使用zhifeiya名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}

/**
将aData数据转换json
*/
func ToJsonMarshaltr(aData interface{}) (string, error) {
	bData, e := json.Marshal(aData)
	if e != nil {
		return "", e
	}
	strData := BytesToStr(bData)
	return strData, e
}

/**
将Json字符串数据转换 interface
err := ToJsonUnmarshal("{[]}",&aData)
*/
func ToJsonUnmarshal(strJson string, aData interface{}) error {
	bData := StrToBytes(strJson)
	return json.Unmarshal(bData, aData)
}
