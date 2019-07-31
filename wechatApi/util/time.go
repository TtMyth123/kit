package util

import "time"

func GetCurTimeStamp() int64 {
	t := time.Now().UnixNano()
	t1 := t / int64(time.Second)
	return t1
}
