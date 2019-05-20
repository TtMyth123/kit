package ttLog

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"runtime"
	"strings"
)

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger

//运行方式
var runmode string

func InitLogs() {
	consoleLogs = logs.NewLogger(1)
	consoleLogs.SetLogger(logs.AdapterConsole)
	consoleLogs.Async() //异步
	consoleLogs.EnableFuncCallDepth(true)

	fileLogs = logs.NewLogger(10000)
	level := beego.AppConfig.String("logs::level")
	//fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/rms.log",
	//	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
	//	"level":`+level+`,
	//	"daily":true,
	//	"maxdays":10}`)
	fileLogs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/rms.log",
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
		"level":`+level+`,
		"daily":true,
		"maxdays":10}`)
	fileLogs.Async() //异步
	runmode = strings.TrimSpace(strings.ToLower(beego.AppConfig.String("runmode")))
	if runmode == "" {
		runmode = "dev"
	}
}

func LogEmergency(v interface{}) {
	log("emergency", v)
}

func LogAlert(v ...interface{}) {
	log("alert", v)
}

func LogCritical(v ...interface{}) {
	log("critical", v)
}

func LogError(v ...interface{}) {
	log("error", v)
}

func LogWarning(v ...interface{}) {
	log("warning", v)
}

func LogNotice(v ...interface{}) {
	log("notice", v)
}

func LogInfo(v ...interface{}) {
	log("info", v)
}

func LogDebug(v ...interface{}) {
	log("debug", v)
}

func LogTrace(v ...interface{}) {
	log("trace", v)
}

//Log 输出日志
func log(level string, v ...interface{}) {
	strInfos := fmt.Sprint(v)
	format := "%s"
	if level == "" {
		level = "debug"
	}
	if runmode == "dev" {
		switch level {
		case "emergency":
			fileLogs.Emergency(format, strInfos)
		case "alert":
			fileLogs.Alert(format, strInfos)
		case "critical":
			fileLogs.Critical(format, strInfos)
		case "error":
			fileLogs.Error(format, strInfos)
			buf := make([]byte,10000)
			runtime.Stack(buf,true)
			fileLogs.Error(format, string(buf))
		case "warning":
			fileLogs.Warning(format, strInfos)
		case "notice":
			fileLogs.Notice(format, strInfos)
		case "info":
			fileLogs.Info(format, strInfos)
		case "debug":
			fileLogs.Debug(format, strInfos)
		case "trace":
			fileLogs.Trace(format, strInfos)
		default:
			fileLogs.Debug(format, strInfos)
		}
	}

	switch level {
	case "emergency":
		consoleLogs.Emergency(format, strInfos)
	case "alert":
		consoleLogs.Alert(format, strInfos)
	case "critical":
		consoleLogs.Critical(format, strInfos)
	case "error":
		consoleLogs.Error(format, strInfos)
	case "warning":
		consoleLogs.Warning(format, strInfos)
	case "notice":
		consoleLogs.Notice(format, strInfos)
	case "info":
		consoleLogs.Info(format, strInfos)
	case "debug":
		consoleLogs.Debug(format, strInfos)
	case "trace":
		consoleLogs.Trace(format, strInfos)
	default:
		consoleLogs.Debug(format, strInfos)
	}
}
