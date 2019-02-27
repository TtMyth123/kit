package funs

import (
	"strconv"
)

func CONCATENATE(input interface{}) (string, error) {
	switch input.(type) {
	case nil:
		return "", nil
	case int:
		return strconv.Itoa(input.(int)), nil
	case string:
		return input.(string), nil
	//case int64:
	//	return strconv.FormatInt(input.(int64), 10)
	//case int8:
	//	return strconv.FormatInt(input.(int64), 10)
	//case int32:
	//	return strconv.FormatInt(input.(int64), 10)
	case float64:
		return strconv.FormatFloat(input.(float64), 'f', 2, 64), nil
		//case float32:
		//	return strconv.FormatFloat(input.(float64), 'f', 5, 32)
		//case string:
		return input.(string), nil

	default:
		return "", nil
		//return "", errors.New("CONCATENATE 不支持的数据类型:" + fmt.Sprint(input))
	}
}

func CONCATENATE2(input1 interface{}, input2 interface{}) (string, error) {
	return CONCATENATEn(input1, input2)
}
func CONCATENATE3(input1 interface{}, input2 interface{}, input3 interface{}) (string, error) {
	return CONCATENATEn(input1, input2, input3)
}
func CONCATENATE4(input1 interface{}, input2 interface{}, input3 interface{}, input4 interface{}) (string, error) {
	return CONCATENATEn(input1, input2, input3, input4)
}

func CONCATENATEn(input1 ...interface{}) (string, error) {
	allS := ""
	for _, input := range input1 {
		s, e := CONCATENATE(input)
		if e != nil {
			return allS, e
		}
		allS = allS + s
	}
	return allS, nil
}
