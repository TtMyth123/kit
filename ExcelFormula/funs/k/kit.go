package k

import (
	"errors"
)

func interface2() {

}
func SotrInt(arrData []int, b bool) []int {
	datas := arrData
	ilen := len(datas)
	for i := 0; i < ilen; i++ {
		for j := 0; j < ilen-i-1; j++ {
			if b {
				if datas[j] > datas[j+1] {
					t := datas[j]
					datas[j] = datas[j+1]
					datas[j+1] = t
				}
			} else {
				if datas[j] < datas[j+1] {
					t := datas[j]
					datas[j] = datas[j+1]
					datas[j+1] = t
				}
			}

		}
	}
	return datas
}

func Sotr2DInt(arrData [][]int, b bool) []int {
	ilen := len(arrData)

	datas := make([]int, 0)
	for i := 0; i < ilen; i++ {
		for j := 0; j < len(arrData[i]); j++ {
			datas = append(datas, arrData[i][j])
		}
	}

	return SotrInt(datas, b)
}

func Sotr(arrData []float64, b bool) []float64 {
	datas := arrData
	ilen := len(datas)
	for i := 0; i < ilen; i++ {
		for j := 0; j < ilen-i-1; j++ {
			if b {
				if datas[j] > datas[j+1] {
					t := datas[j]
					datas[j] = datas[j+1]
					datas[j+1] = t
				}
			} else {
				if datas[j] < datas[j+1] {
					t := datas[j]
					datas[j] = datas[j+1]
					datas[j+1] = t
				}
			}

		}
	}
	return datas
}

func Sotr2D(arrData [][]float64, b bool) []float64 {
	ilen := len(arrData)

	datas := make([]float64, 0)
	for i := 0; i < ilen; i++ {
		for j := 0; j < len(arrData[i]); j++ {
			datas = append(datas, arrData[i][j])
		}
	}

	return Sotr(datas, b)
}

func Inter1DTo1DFloat(input1 interface{}, def float64) ([]float64, error) {
	rData := make([]float64, 0)
	switch input1.(type) {
	case []interface{}:
		for _, item := range input1.([]interface{}) {
			switch item.(type) {
			case float64:
				v := item.(float64)
				rData = append(rData, v)
			case int:
				v := float64(item.(int))
				rData = append(rData, v)
			default:
				rData = append(rData, def)
			}
		}
	default:
		return rData, errors.New("类型不对")
	}

	return rData, nil
}

func Inter2DTo1DFloat(input1 interface{}, def float64) ([]float64, error) {
	rData := make([]float64, 0)
	switch input1.(type) {
	case [][]interface{}:
		for _, itemV := range input1.([][]interface{}) {
			for _, item := range itemV {
				switch item.(type) {
				case error:
					rData = append(rData, 0)
				case int:
					v := float64(item.(int))
					rData = append(rData, v)
				case float64:
					v := item.(float64)
					rData = append(rData, v)
				default:
					rData = append(rData, def)
				}
			}
		}
	default:
		return rData, errors.New("类型不对")
	}

	return rData, nil
}

/**
截取中英字符串
*/
func SubCnString(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}
func LeftCnString(str string, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if length > lth {
		length = lth
	}

	return string(rs[:length])
}

func RightCnString(str string, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if length > lth {
		length = lth
	}

	statIndex := lth - length
	if statIndex < 0 {
		statIndex = 0
	}

	return string(rs[statIndex:])
}

func IndexCnString(str string, index int) string {
	rs := []rune(str)
	lth := len(rs)
	if index > lth {
		index = lth
	}

	return string(rs[index:])
}

/**
获取中英字符串长度
如:a了工2
为:4
*/
func LenCnString(str string) int {

	rs := []rune(str)
	lth := len(rs)
	return lth
}
