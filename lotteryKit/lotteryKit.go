package lotteryKit

import (
	"fmt"
	"strconv"
	"strings"
	"ttmyth123/kit/strconvEx"
)

func GetStrNum2Arr(strNums string)[]int {
	arrStrN := strings.Split(strNums, ",")
	iLen := len(arrStrN)
	numList := make([]int, iLen)
	for i := 0; i < iLen; i++ {
		numList[i] = strconvEx.StrTry2Int(arrStrN[i], 0)
	}

	return numList
}

func GetStrNum2ArrStr(strNums string, s int)[]string {
	arrStrN := strings.Split(strNums, ",")
	iLen := len(arrStrN)
	numList := make([]string, iLen)
	for i := 0; i < iLen; i++ {
		strFormat := "%0"+strconv.Itoa(s)+"d"
		numList[i] = fmt.Sprintf(strFormat, strconvEx.StrTry2Int(arrStrN[i], 0))
	}

	return numList
}

func  GetArrNum2String(arr []int, d int)string {
	strNums :=""
 	iLen := len(arr)
 	fm := ",%0"+strconv.Itoa(d)+"d"
	for i := 0; i < iLen; i++ {
		strNums = strNums +fmt.Sprintf(fm, arr[i])
	}
	if iLen>0 {
		return strNums[1:]
	}

	return 	strNums
}