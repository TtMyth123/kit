package funs

import "errors"

func HARMEAN(input interface{}) (float64, error) {
	switch input.(type) {
	case nil:
		return 0, nil
	case int:
		return float64(input.(int)), nil
	case float64:
		return input.(float64), nil
	case []float64:
		s := float64(0)
		l := float64(len(input.([]float64)))
		for _, item := range input.([]float64) {
			s = s + 1/item
		}
		return l / s, nil
	case [][]float64:
		s := float64(0)
		l := 0
		for _, items := range input.([][]float64) {
			for _, item := range items {
				l++
				s = s + 1/item
			}
		}
		return float64(l) / s, nil
	case []interface{}:
		s := float64(0)
		l := float64(len(input.([]interface{})))
		for _, item := range input.([]interface{}) {
			switch item.(type) {
			case float64:
				v := item.(float64)
				s = s + 1/v
				break
			case int:
				v := float64(item.(int))
				s = s + 1/v
				break
			default:
				break
			}
		}
		return l / s, nil
	case [][]interface{}:
		s := float64(0)
		var outer [][]interface{}
		var inner []interface{}

		outer = input.([][]interface{})
		l := 0
		for i := 0; i < len(outer); i++ {
			inner = outer[i]
			for j := 0; j < len(inner); j++ {
				l++
				v := inner[j].(float64)
				s = s + 1/v
			}
		}

		return float64(l) / s, nil
	default:
		return 0, errors.New("数据格式不对.")
	}
}

func HARMEAN2(input1 interface{}, input2 interface{}) (float64, error) {
	f1 := float64(0)
	f2 := float64(0)
	switch input1.(type) {
	case nil:
		f1 = 0
	case int:
		f1 = float64(input1.(int))
	case float64:
		f1 = input1.(float64)
	}

	switch input2.(type) {
	case nil:
		f2 = 0
	case int:
		f2 = float64(input2.(int))
	case float64:
		f2 = input2.(float64)
	}
	data := []float64{f1, f2}

	return HARMEAN(data)
}
