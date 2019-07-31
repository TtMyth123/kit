package stringKit

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	curI        int
	numGuidLock sync.RWMutex
	curDate     string
)

func GetBetweenStr(strS, beginStr, endStr string) string {
	bI := strings.Index(strS, beginStr)
	if bI >= 0 {
		s1 := strS[bI+len(beginStr):]
		eI := strings.Index(s1, endStr)
		if eI >= 0 {
			return s1[:eI]
		}

	}
	return ""
}

func GetLaterStrLast(strS, beginStr string) string {
	bI := strings.LastIndex(strS, beginStr)
	if bI >= 0 {
		s1 := strS[bI+len(beginStr):]
		return s1
	}
	return ""
}

func GetLaterStr(strS, beginStr string) string {
	bI := strings.Index(strS, beginStr)
	if bI >= 0 {
		s1 := strS[bI+len(beginStr):]
		return s1
	}
	return ""
}

func GetBeforeStr(strS, beginStr string) string {
	bI := strings.Index(strS, beginStr)
	if bI >= 0 {
		s1 := strS[:bI]
		return s1
	}
	return ""
}

func GetJsonStr(data interface{}) string {
	str, _ := json.Marshal(data)
	return string(str)
}

func GetNumGuid(i int) string {
	numGuidLock.Lock()
	defer numGuidLock.Unlock()
	t := time.Now()
	strT := t.Format("060102")
	if curDate != strT {
		curI = 0
		curDate = strT
	}

	curI++
	guid := fmt.Sprintf("%s%0"+strconv.Itoa(i)+"d", strT, curI)
	return guid
}
