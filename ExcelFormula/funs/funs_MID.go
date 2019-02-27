package funs

import (
	"errors"
	"fmt"
	"ttmyth123/kit/ExcelFormula/funs/k"
)

func MID(input interface{}, input2 interface{}, input3 interface{}) (string, error) {
	str := ""
	strIndex := 0
	strLen := 0
	switch input.(type) {
	case nil:
		str = ""
	case string:
		str = input.(string)
	default:
		return "", errors.New("MID 函数的第1个参数必须是string类型")
	}

	switch input2.(type) {
	case nil:
		strIndex = 0
	case int:
		strIndex = input.(int)
	case float64:
		strIndex = int(input2.(float64))
	default:
		return "", errors.New("MID 不支持的数据类型:" + fmt.Sprint(input2))
	}

	switch input3.(type) {
	case nil:
		strLen = 0
	case int:
		strLen = input3.(int)
	case float64:
		strLen = int(input3.(float64))
	default:
		return "", errors.New("MID 不支持的数据类型:" + fmt.Sprint(input3))
	}

	if strIndex < 0 {
		strIndex = 0
	}
	strIndex = strIndex - 1
	rs := []rune(str)
	lth := len(rs)
	if strLen+strIndex > lth {
		strLen = lth - strIndex
	}

	//return str[strIndex:endIndex], nil
	return k.SubCnString(str, strIndex, strLen), nil
}
