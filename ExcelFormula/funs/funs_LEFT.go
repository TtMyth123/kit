package funs

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/kit/ExcelFormula/funs/k"
	"strconv"
)

func LEFT1(input interface{}) (string, error) {
	return LEFT2(input, 1)
}

func LEFT2(input interface{}, input2 interface{}) (string, error) {
	str := ""
	i := 0
	switch input2.(type) {
	case nil:
		i = 0
	case int:
		i = input2.(int)
	case float64:
		i = int(input2.(float64))
	default:
		return "", errors.New("LEFT 函数的第2个参数必须是string类型")
	}

	switch input.(type) {
	case nil:
		str = ""
	case int:
		str = strconv.Itoa(input.(int))
	case float64:
		str = strconv.FormatFloat(input.(float64), 'f', 2, 64)
	case string:
		str = input.(string)
	default:
		return "", errors.New("LEFT 不支持的数据类型:" + fmt.Sprint(input))
	}

	rs := []rune(str)
	iLen := len(rs)

	if i < 0 {
		i = 0
	}
	if i > iLen {
		i = iLen
	}

	r := k.SubCnString(str, 0, i)
	return r, nil
}
