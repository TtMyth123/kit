package funs

import (
	"errors"
	"ttmyth123/kit/ExcelFormula/funs/k"
)

func SUMPRODUCT1(input1 interface{}) (float64, error) {
	switch input1.(type) {
	case nil:
		return 0, nil
	case []float64:
		fInput1 := input1.([]float64)
		s := float64(0)
		lr := len(fInput1)
		for i := 0; i < lr; i++ {
			s = s + fInput1[i]
		}
		return s, nil

	case [][]float64:
		fInput1 := input1.([][]float64)
		lr := len(fInput1)
		s := float64(0)
		for i := 0; i < lr; i++ {
			lc := len(fInput1[i])
			for j := 0; j < lc; j++ {
				s = s + fInput1[i][j]
			}
		}
		return s, nil
	case []interface{}:
		s := float64(0)
		for _, item := range input1.([]interface{}) {
			switch item.(type) {
			case float64:
				v := item.(float64)
				s = s + v
			case int:
				v := float64(item.(int))
				s = s + v
			default:
				return 0, errors.New("类型不对")
			}
		}
		return s, nil
	case [][]interface{}:
		s := float64(0)
		for _, itemV := range input1.([][]interface{}) {
			for _, item := range itemV {
				switch item.(type) {
				case float64:
					v := item.(float64)
					s = s + v
				case int:
					v := float64(item.(int))
					s = s + v
				default:
					return 0, errors.New("类型不对")
				}
			}
		}
		return s, nil
	default:
		return 0, errors.New("类型不对")
	}
	return 0, errors.New("参数有问题")
}

func SUMPRODUCT2(input1 interface{}, input2 interface{}) (float64, error) {
	switch input1.(type) {
	case nil:
		return 0, nil
	case []float64:
		fInput1 := input1.([]float64)
		switch input2.(type) {
		case []float64:
			fInput2 := input2.([]float64)
			l := len(fInput1)
			if l == len(fInput2) {
				s := float64(0)
				for i := 0; i < l; i++ {
					s = s + fInput1[i]*fInput2[i]
				}
				return s, nil
			}
		}
	case [][]float64:
		fInput1 := input1.([][]float64)
		switch input2.(type) {
		case [][]float64:
			fInput2 := input2.([][]float64)
			lr := len(fInput1)
			if lr == len(fInput2) {
				s := float64(0)
				for i := 0; i < lr; i++ {
					lc := len(fInput1[i])
					if lc != len(fInput2[i]) {
						return 0, errors.New("参数有问题 长度不一致")
					}
					for j := 0; j < lc; j++ {
						s = s + fInput1[i][j]*fInput2[i][j]
					}
				}
				return s, nil
			}
		}
	case []interface{}:
		fInput1, e := k.Inter1DTo1DFloat(input1, 0)
		if e != nil {
			return 0, e
		}
		fInput2, e := k.Inter1DTo1DFloat(input2, 0)
		if e != nil {
			return 0, e
		}
		iL1 := len(fInput1)
		iL2 := len(fInput2)
		if iL1 != iL2 {
			return 0, errors.New("参数有问题 长度不一致")
		}
		s := float64(0)
		for i := 0; i < iL1; i++ {
			s = s + fInput1[i]*fInput2[i]
		}
		return s, nil
	case [][]interface{}:
		fInput1, e := k.Inter2DTo1DFloat(input1, 0)
		if e != nil {
			return 0, e
		}
		fInput2, e := k.Inter2DTo1DFloat(input2, 0)
		if e != nil {
			return 0, e
		}
		iL1 := len(fInput1)
		iL2 := len(fInput2)
		if iL1 != iL2 {
			return 0, errors.New("参数有问题 长度不一致")
		}
		s := float64(0)
		for i := 0; i < iL1; i++ {
			s = s + fInput1[i]*fInput2[i]
		}
		return s, nil
	default:
		return 0, errors.New("类型不对")
	}
	return 0, errors.New("参数有问题")
}
