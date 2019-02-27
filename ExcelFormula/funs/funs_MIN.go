package funs

import (
	"errors"
	"math"
)

func MIN(input interface{}) (float64, error) {
	switch input.(type) {
	case nil:
		return 0, nil
	case int:
		return float64(input.(int)), nil
	case float64:
		return input.(float64), nil
	case []float64:
		min := math.MaxFloat64

		for _, item := range input.([]float64) {
			if min > item {
				min = item
			}
		}
		return min, nil
	case [][]float64:
		min := math.MaxFloat64
		for _, items := range input.([][]float64) {
			for _, item := range items {
				if min > item {
					min = item
				}
			}
		}
		return min, nil
	case []interface{}:
		min := math.MaxFloat64
		for _, item := range input.([]interface{}) {
			switch item.(type) {
			case float64:
				v := item.(float64)
				if min > v {
					min = v
				}
				break
			default:
				break
			}
		}
		return min, nil
	case [][]interface{}:
		min := math.MaxFloat64
		var outer [][]interface{}
		var inner []interface{}

		outer = input.([][]interface{})
		for i := 0; i < len(outer); i++ {
			inner = outer[i]
			for j := 0; j < len(inner); j++ {
				v := inner[j].(float64)
				if min > v {
					min = v
				}
			}
		}

		return min, nil
	default:
		return 0, errors.New("数据格式不对")
	}
}

func MIN2(input1 interface{}, input2 interface{}) (float64, error) {
	a1, e := MIN(input1)
	if e != nil {
		return 0, e
	}

	a2, e := MIN(input1)
	if e != nil {
		return 0, e
	}
	if a1 > a2 {
		return a2, nil
	}
	return a1, nil
}
