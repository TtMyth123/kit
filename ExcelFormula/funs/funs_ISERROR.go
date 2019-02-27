package funs

import "math"

func ISERROR(input interface{}) bool {
	switch input.(type) {
	case float64:
		return math.IsInf(input.(float64), 1) || math.IsInf(input.(float64), -1) || math.IsNaN(input.(float64))
	case error:
		return true
	default:
		return false
	}
}
