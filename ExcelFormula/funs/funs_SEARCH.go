package funs

import (
	"errors"
	"strings"
	"ttmyth123/kit/ExcelFormula/funs/k"
)

func SEARCH3(input interface{}, input2 interface{}, input3 interface{}) (int, error) {
	str1 := ""
	str2 := ""
	starIndex := 0
	switch input.(type) {
	case nil:
		str1 = ""
	case string:
		str1 = input.(string)
	default:
		return 0, errors.New("SEARCH 函数的第1个参数必须是string类型")
	}

	switch input2.(type) {
	case nil:
		str2 = ""
	case string:
		str2 = input2.(string)
	default:
		return 0, errors.New("SEARCH 函数的第2个参数必须是string类型")
	}

	switch input3.(type) {
	case nil:
		starIndex = 0
	case int:
		starIndex = input3.(int) - 1
	case float64:
		starIndex = int(input3.(float64)) - 1
	default:
		return 0, errors.New("SEARCH 函数的第1个参数必须是string类型")
	}
	iIndex := strings.Index(str2[starIndex:], str1)
	if iIndex < 0 {
		return 0, errors.New("-1")
	}
	return k.LenCnString(str2[:iIndex]), nil
}

func SEARCH2(input interface{}, input2 interface{}) (int, error) {
	return SEARCH3(input, input2, 1)
}
