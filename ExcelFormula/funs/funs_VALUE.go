package funs

import (
	"strconv"
	"strings"
)

func VALUE(input interface{}) (float64, error) {
	switch input.(type) {
	case nil:
		return 0, nil
	case string:
		str := input.(string)
		str = strings.Replace(str, ",", "", 999)
		f, e := strconv.ParseFloat(str, 32)
		return f, e
	case float64:
		return input.(float64), nil
	case int:
		return float64(input.(int)), nil
	}
	return 0, nil
}
