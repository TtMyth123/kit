package funs

import (
	"errors"
	"strconv"
	"strings"
)

func FIND3(input1 interface{}, input2 interface{}, input3 interface{}) (int, error) {
	f := ""
	str := ""
	index := 0
	switch input1.(type) {
	case nil:
		f = ""
	case string:
		f = input1.(string)
	case int:
		f = strconv.Itoa(input1.(int))
	case float64:
		f = strconv.FormatFloat(input1.(float64), 'f', 0, 64)
	default:
		return 0, errors.New("FIND2_E 函数的第1个参数必须是string类型")
	}

	switch input2.(type) {
	case nil:
		str = ""
	case string:
		str = input2.(string)
	default:
		return 0, errors.New("FIND2_E 函数的第2个参数必须是string类型")
	}

	switch input3.(type) {
	case nil:
		index = 0
	case int:
		index = input3.(int)
	case float64:
		index = int(input3.(float64))
	default:
		return 0, errors.New("FIND2_E 函数的第3个参数必须是int类型")
	}

	iLen := len(str)
	if index > iLen {
		index = iLen
	}

	newIndex := strings.Index(str[index:], f) + index + 1
	if newIndex >= 0 {
		rs := []rune(str[:newIndex])
		lth := len(rs)
		return lth, nil
	} else {
		return -1, nil
	}
}

func FIND2(input1 interface{}, input2 interface{}) (int, error) {
	return FIND3(input1, input2, 0)
}
