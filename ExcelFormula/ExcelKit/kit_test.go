package ExcelKit

import (
	"fmt"
	"testing"
)

func TestGetColNum2Str(t *testing.T) {
	s := GetColNum2Str(2)
	fmt.Println(s)
	if s!="B" {
		t.Fail()
	}

}

func TestGetCellXY2Str(t *testing.T) {
	s := GetCellXY2Str(1,2)
	fmt.Println(s)
	if s!="B1" {
		t.Fail()
	}
}

func TestGetColNum(t *testing.T) {
	str := "BG"
	i :=GetColNum(str)
	fmt.Println(str,":",i)
}