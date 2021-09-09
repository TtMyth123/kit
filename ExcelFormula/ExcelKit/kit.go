package ExcelKit

import (
	"fmt"
	"strings"
)

func GetColNum(strCol string)int  {
	strCol = strings.ToUpper(strCol)
	arr := []byte(strCol)

	s := 0
	for _,n := range arr{
		n1 := n-'A'+1
		s=s*26+int(n1)
	}
	return s
}
func GetFullCellXY2Str(sheetName string,row, col int)string  {
	strC := GetColNum2Str(col)
	return fmt.Sprintf("%s!%s%d",sheetName,strC,row)
}

func GetCellXY2Str(row, col int)string  {
	strC := GetColNum2Str(col)
	return fmt.Sprintf("%s%d",strC,row)
}
func GetColNum2Str(col int) string {
	a:='A'
	b := int(a)+col-1
	s := fmt.Sprintf("%s",string(b))
	return s
}
