package gtLog

import (
	"fmt"
	"runtime"
)

//import "fmt"

const (
	Error = 0
	Warn  = 1
	Debug = 2
	Info  = 3
)

var (
	showLevel int
)

func Log(t int, tag string, msg ...interface{}) {
	//fmt.Println(t,tag,msg)
	if t <= showLevel {
		if t == Error {
			pc, file, line, ok := runtime.Caller(1)
			if ok {
				f := runtime.FuncForPC(pc)
				fmt.Println(file, f.Name(), line)
			}
			fmt.Println(t, tag, msg)
		} else {
			fmt.Println(t, tag, msg)
		}
	}
}

func LogSkip(t int, tag string, Skip int, msg ...interface{}) {
	//fmt.Println(t,tag,msg)
	if t <= showLevel {
		if t == Error {
			pc, file, line, ok := runtime.Caller(Skip)
			if ok {
				f := runtime.FuncForPC(pc)
				fmt.Println(file, f.Name(), line)
			}
			fmt.Println(t, tag, msg)
		} else {
			fmt.Println(t, tag, msg)
		}
	}
}

func SetLogLevel(l string) {
	if l == "e" || l == "E" {
		showLevel = Error
	} else if l == "w" || l == "W" {
		showLevel = Warn
	} else if l == "d" || l == "D" {
		showLevel = Debug
	} else if l == "i" || l == "I" {
		showLevel = Info
	} else {
		showLevel = Error
	}

	fmt.Println("日志级别为：", showLevel)
}

func init() {
	showLevel = Error
}
