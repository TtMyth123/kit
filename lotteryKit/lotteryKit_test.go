package lotteryKit


import (
	"fmt"
	"testing"
)

func TestGetNumList(t *testing.T) {
	strNum := "1,2,3,5,7"
	arr:= GetStrNum2Arr(strNum)
	fmt.Println(arr)
}

func TestGetArrNum2String(t *testing.T) {
	arr := []int{1,5,6,7}
	r:= GetArrNum2String(arr, 2)
	fmt.Println(r)
}
