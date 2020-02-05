package sqlKit

import (
	"bytes"
	"fmt"
	"strings"
)

/**
获取 分页的偏移量。
AllCount：总数量
onePageCount：一页的数量
page：第几面。1：为第1页

pageCount
*/
func GetOffset(AllCount, onePageCount, page int) (offset,pageCount int) {
	//4/2
	c := AllCount / onePageCount
	if AllCount%onePageCount != 0 {
		c++
	}
	if page > c {
		page = c
	}

	offset = onePageCount * (page - 1)
	pageCount = c
	return offset, pageCount
}

func GetWhereInStr(fieldName, str string) string {
	arrStr := strings.Split(str, ",")
	a := bytes.Buffer{}
	for _, s1 := range arrStr {
		a.WriteString(",'")
		a.WriteString(s1)
		a.WriteString("'")
	}

	strValue := a.String()
	if len(strValue) > 0 {
		strValue = strValue[1:]
	}

	sql := fmt.Sprintf("%s in (%s)", fieldName, strValue)

	return sql
}