package stringKit

import "strings"

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
