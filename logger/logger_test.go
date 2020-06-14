package logger

import (
	"testing"
)

func init() {
	config := &LogConfig{
		Formatter:    TextFormatter,
		ReportCaller: true,
		SkipCaller:   1,
		//ForceColors:  true,
		//LogFile:      "/tmp/test.log",
		LogLevel: "trace",
	}
	SetLogger(config)

}

func TestLogger(t *testing.T) {

	// AddHook(hook.NewTraceIdHook(strings.ShortUUID()))

	Trace("这是Trace日志")
	Debug("这是debug日志")
	Info("这是info日志")
	Warn("这是Warn日志")
	Error("这是info日志")

	Traceln("这是Trace日志")
	Debugln("这是debug日志")
	Infoln("这是info日志")
	Warnln("这是Warn日志")
	Errorln("这是info日志")

	//Fatal("这是Fatal日志")
	//Panic("这是Panic日志")
}

//func BenchmarkTest(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Info("这是info日志dddddddddddddddddd")
//	}
//}
