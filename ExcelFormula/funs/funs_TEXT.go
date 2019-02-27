package funs

import (
	"errors"
	"fmt"
	"strconv"
)

func TEXT(input interface{}, input2 interface{}) (string, error) {
	//f := ""
	switch input2.(type) {
	case nil:
	case string:
		//f = input2.(string)
	default:
		return "", errors.New("TEXT 函数的第2个参数必须是string类型")
	}

	switch input.(type) {
	case nil:
		return "", nil
	case int:
		return strconv.Itoa(input.(int)), nil
	case float64:
		return strconv.FormatFloat(input.(float64), 'f', 2, 64), nil
	default:
		return "", errors.New("TEXT 不支持的数据类型:" + fmt.Sprint(input))
	}
}
