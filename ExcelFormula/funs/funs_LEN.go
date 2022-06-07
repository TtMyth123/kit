package funs

import "github.com/TtMyth123/kit/ExcelFormula/funs/k"

func LEN(input interface{}) (int, error) {
	f := ""
	switch input.(type) {
	case nil:
		f = ""
	case string:
		f = input.(string)
	}
	return k.LenCnString(f), nil
}
