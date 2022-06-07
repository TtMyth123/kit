package funs

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/kit/ExcelFormula/funs/k"
)

func MEDIAN(input interface{}) (float64, error) {
	switch input.(type) {
	case nil:
		return 0, nil
	case int:
		return float64(input.(int)), nil
	case []int:
		aData := input.([]int)
		sortData := k.SotrInt(aData, true)
		iLen := len(sortData)
		if iLen%2 == 0 {
			i := iLen / 2
			return float64((sortData[i-1] + sortData[i])) / 2, nil
		} else {
			i := iLen / 2
			return float64(sortData[i]), nil
		}
	case [][]int:
		aData := input.([][]int)

		sortData := k.Sotr2DInt(aData, true)
		iLen := len(sortData)
		if iLen%2 == 0 {
			i := iLen / 2
			return float64(sortData[i-1]+sortData[i]) / 2, nil
		} else {
			i := iLen / 2
			return float64(sortData[i]), nil
		}
	case float64:
		return input.(float64), nil
	case []float64:
		aData := input.([]float64)

		sortData := k.Sotr(aData, true)
		iLen := len(sortData)
		if iLen%2 == 0 {
			i := iLen / 2
			return (sortData[i-1] + sortData[i]) / 2, nil
		} else {
			i := iLen/2 + 1
			return sortData[i], nil
		}

	case [][]float64:
		aData := input.([][]float64)

		sortData := k.Sotr2D(aData, true)
		iLen := len(sortData)
		if iLen%2 == 0 {
			i := iLen / 2
			return (sortData[i-1] + sortData[i]) / 2, nil
		} else {
			i := iLen / 2
			return sortData[i], nil
		}
	case []interface{}:
		aData := input.([]interface{})
		fData := make([]float64, 0)
		for _, v := range aData {
			fv := float64(0)

			switch v.(type) {
			case int:
				fv = float64(v.(int))
				fData = append(fData, fv)
			case float64:
				fv = v.(float64)
				fData = append(fData, fv)
			}
		}
		sortData := k.Sotr(fData, true)
		iLen := len(sortData)
		if iLen%2 == 0 {
			i := iLen / 2
			return (sortData[i-1] + sortData[i]) / 2, nil
		} else {
			i := iLen / 2
			return sortData[i], nil
		}
	case [][]interface{}:
		aData := input.([][]interface{})
		fData := make([]float64, 0)
		for _, vi := range aData {
			for _, v := range vi {
				fv := float64(0)
				switch v.(type) {
				case int:
					fv = float64(v.(int))
					fData = append(fData, fv)
				case float64:
					fv = v.(float64)
					fData = append(fData, fv)
				}
			}
		}
		sortData := k.Sotr(fData, true)
		iLen := len(sortData)
		if iLen%2 == 0 {
			i := iLen / 2
			return (sortData[i-1] + sortData[i]) / 2, nil
		} else {
			i := iLen / 2
			return sortData[i], nil
		}
	default:
		return 0, errors.New("MEDIAN 函数的第1个参数必须是 数值 类型" + fmt.Sprint(input))
	}
}

func MEDIAN2(input1 interface{}, input2 interface{}) (float64, error) {
	i1 := float64(0)
	i2 := float64(0)
	switch input1.(type) {
	case nil:
		i1 = 0
	case int:
		i1 = float64(input1.(int))
	case float64:
		i1 = float64(input1.(int))
	default:
		return 0, errors.New("MEDIAN 函数的第1个参数必须是 数值 类型")
	}

	switch input2.(type) {
	case nil:
		i2 = 0
	case int:
		i2 = float64(input2.(int))
	case float64:
		i2 = float64(input2.(int))
	default:
		return 0, errors.New("MEDIAN 函数的第2个参数必须是 数值 类型")
	}

	return (i1 + i2) / 2, nil
}
