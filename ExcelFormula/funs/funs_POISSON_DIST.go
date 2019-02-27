package funs

import (
	"errors"
	"math"
)

func POISSON_DIST3(input1 interface{}, input2 interface{}, input3 interface{}) (float64, error) {
	var x, r float64
	switch input1.(type) {
	case nil:
		x = 0
	case int:
		x = float64(input1.(int))
	case float64:
		x = input1.(float64)
	default:
		return 0, errors.New("POISSON_DIST 参数1数据类型不对")
	}

	switch input2.(type) {
	case nil:
		r = 0
	case int:
		r = float64(input2.(int))
	case float64:
		r = input2.(float64)
	default:
		return 0, errors.New("POISSON_DIST 参数2数据类型不对")
	}

	r1 := math.Pow(r, x) * math.Pow(math.E, -r)
	r2 := float64(X1(int(x)))
	return r1 / r2, nil
}

func X1(x int) int {
	s := 1
	for i := 2; i <= x; i++ {
		s = s * i
	}
	return s
}
