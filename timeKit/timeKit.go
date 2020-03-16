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

/**
根据Java的LongTime生成一个时间
longTime: Java的LongTime。如1551513615000=>2019-03-02 16:00:15 +0800 CST
 */
func NewTimeByJavaTimeLong(longTime int64) time.Time {
	t := time.Unix(0, 0)
	t = t.Add(time.Duration(longTime) * time.Millisecond)
	return t
}

/**
把时间转换成Java的LongTime
 */
func GetJavaTimeLong(t time.Time) int64 {
	return t.Unix() * 1000
}

/**
获取日期 格式为:yyyy-mm-dd
 */
func GetDate(t time.Time) string {
	return t.Format(DateLayout)
}

const DateLayout = "2006-01-02"

func GetDateForTime(strTime string) (time.Time, error) {
	return time.ParseInLocation(DateLayout, strTime,time.Local)
}
func GetTime(strTime string) (time.Time, error)  {
	return time.ParseInLocation("2006-01-02 15:04:05", strTime,time.Local)
}
