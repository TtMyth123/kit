package funs

import (
	"errors"
	"fmt"
	"strconv"
	"ttmyth123/kit/ExcelFormula/funs/k"
)

func REPLACE4(input interface{}, input2 interface{}, input3 interface{}, input4 interface{}) (string, error) {
	str := ""
	statIndex := 1
	iL := 0
	repStr := ""
	switch input.(type) {
	case error:
		//return "", input.(error)
		return "", nil
	case nil:
		str = ""
	case int:
		str = strconv.Itoa(input.(int))
	case float64:
		str = strconv.FormatFloat(input.(float64), 'f', 2, 64)
	case string:
		str = input.(string)
	default:
		return "", errors.New("REPLACE 函数的第1个参数必须是string类型" + fmt.Sprint(input))
	}

	switch input2.(type) {
	case error:
		//return "", input2.(error)
		return "", nil
	case nil:
		statIndex = 1
	case int:
		statIndex = input2.(int)
	case float64:
		statIndex = int(input2.(float64))
	case string:
		t1, e := strconv.Atoi(input2.(string))
		if e != nil {
			return "", errors.New("REPLACE 函数的第2个参数必须是int类型 " + fmt.Sprint(input2))
		}
		statIndex = t1
	default:
		return "", errors.New("REPLACE 函数的第2个参数必须是string类型" + fmt.Sprint(input2))
	}

	switch input3.(type) {
	case error:
		//return "", input3.(error)
		return "", nil
	case nil:
		iL = 0
	case int:
		iL = input3.(int)
	case float64:
		iL = int(input3.(float64))
	case string:
		t1, e := strconv.Atoi(input2.(string))
		if e != nil {
			return "", errors.New("REPLACE 函数的第3个参数必须是int类型 " + fmt.Sprint(input3))
		}
		iL = t1
	default:
		return "", errors.New("REPLACE 函数的第3个参数必须是string类型" + fmt.Sprint(input3))
	}

	switch input4.(type) {
	case error:
		//return "", input4.(error)
		return "", nil
	case nil:
		repStr = ""
	case string:
		repStr = input4.(string)
	default:
		return "", errors.New("REPLACE 函数的第4个参数必须是string类型" + fmt.Sprint(input4))
	}

	lth := k.LenCnString(str)
	statIndex = statIndex - 1
	if statIndex > lth {
		statIndex = lth
	}
	if statIndex < 0 {
		statIndex = 0
	}
	statIndex2 := statIndex + iL
	if statIndex2 > lth {
		statIndex2 = lth
	}

	//strR := str[:statIndex] + repStr + str[:statIndex2]
	strR := k.LeftCnString(str, statIndex) + repStr + k.IndexCnString(str, statIndex2)

	return strR, nil
}
