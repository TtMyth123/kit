package strconvEx

import "strconv"

func StrTry2Int(str string, v int) int {
	r, e := strconv.Atoi(str)
	if e != nil {
		r = v
	}
	return r
}

func StrTry2Int32(str string, v uint32) int32 {
	r := StrTry2Uint64(str, uint64(v))

	return int32(r)
}

func StrTry2Int64(str string, v int64) int64 {
	r, e := strconv.ParseInt(str, 10, 0)
	if e != nil {
		r = v
	}
	return r
}
func StrTry2Uint64(str string, v uint64) uint64 {
	r, e := strconv.ParseUint(str, 10, 0)
	if e != nil {
		r = v
	}
	return r
}

func StrTry2Byte(str string, v byte) byte {
	r := StrTry2Int(str, int(v))

	return byte(r)
}

func StrTry2Uint32(str string, v uint32) uint32 {
	r := StrTry2Int64(str, int64(v))

	return uint32(r)
}

func StrTry2Uint16(str string, v uint16) uint16 {
	r := StrTry2Int(str, int(v))

	return uint16(r)
}

func StrTry2Int16(str string, v int16) int16 {
	r := StrTry2Int(str, int(v))

	return int16(r)
}

func StrTry2Float64(str string, v float64) float64 {
	r, e := strconv.ParseFloat(str, 10)
	if e != nil {
		r = v
	}
	return r
}
