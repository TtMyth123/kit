package timeKit

import (
	"fmt"
	"math/rand"
	"time"
)

var GlobaRand *rand.Rand

//判断两个时间是不是同一天
func IsSameDayByUnix(t1Unix, t2Unix int64) bool {
	t1 := time.Unix(t1Unix, 0)
	t2 := time.Unix(t2Unix, 0)
	return IsSameDay(t1, t2)
}

//判断两个时间是不是同一天
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	if y1 == y2 && m1 == m2 && d1 == d2 {
		return true
	} else {
		return false
	}
}

func GetStrDay(t1 time.Time) string {
	return fmt.Sprintf("%04d%02d%02d", t1.Year(), t1.Month(), t1.Day())
}

func GetStrTime(t1 time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t1.Hour(), t1.Minute(), t1.Second())
}

func init() {
	GlobaRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}
