package logger

import (
	"fmt"
	"github.com/wenit/logrus"
	golog "log"
	"os"
	"regexp"
	"runtime"
	"sync"
)

// text格式日志：DEBU[2020-01-09 16:39:06]main.main:16 aa
// json格式日志：{"func":"main.main:16","level":"debug","msg":"aa","time":"2020-01-09 16:22:26"}
type LogConfig struct {
	Formatter    string `yaml:"formatter" json:"formatter"`         // 日志格式：text-文本格式日志，json-json格式日志
	ReportCaller bool   `yaml:"report_caller" json:"report_caller"` // 是否打印源码文件路径和行号
	SkipCaller   int    `yaml:"skip_caller" json:"skip_caller"`     // 打印源码文件路径和行号时跳过堆栈信息，已便正确显示源码文件和行号
	LogFile      string `yaml:"log_file" json:"log_file"`           // 日志文件路径
	LogLevel     string `yaml:"log_level" json:"log_level"`         // 日志级别：trace，debug，info，warning，error，fatal，panic
	ForceColors  bool   `yaml:"force_colors" json:"force_colors"`   //是否根据不同的报警级别显示颜色
}

var lock sync.Mutex

var log *logrus.Logger

var (
	TextFormatter = "text"
	JSONFormatter = "json"
)

//
//var timestampFormat = "2006-01-02 15:04:05"
var timestampFormat = "2006-01-02 15:04:05.000"

var reg = regexp.MustCompile(`\(.*\)\.`)

func SetLogger(config *LogConfig) *logrus.Logger {

	lock.Lock()
	defer lock.Unlock()
	_log := logrus.New()

	if config.LogFile == "" {
		_log.SetOutput(os.Stdout)
	} else {
		//写入文件
		src, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			golog.Panicln("设置日志文件", err)
		} else {
			//设置输出
			_log.Out = src
		}
	}

	// 设置是否源码文件路径和行号
	_log.SetReportCaller(config.ReportCaller)
	if config.SkipCaller > 0 {
		_log.SkipCaller = config.SkipCaller
	}

	// 设置日志级别
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		// 设置默认日志级别为info级别
		_log.SetLevel(logrus.InfoLevel)
	} else {
		_log.SetLevel(level)
	}

	//设置json格式日志
	if config.Formatter == JSONFormatter {

		_log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:  timestampFormat,
			CallerPrettyfier: callerPrettyfier,
		})
	} else {
		//设置text格式日志
		_log.SetFormatter(&logrus.TextFormatter{
			ForceColors:      config.ForceColors,
			TimestampFormat:  timestampFormat,
			FullTimestamp:    true,
			CallerPrettyfier: callerPrettyfier,
		})
	}
	log = _log

	return log
}

func AddHook(hook logrus.Hook) {
	log.AddHook(hook)
}

func callerPrettyfier(frame *runtime.Frame) (function string, file string) {
	// 替换
	tempName := reg.ReplaceAllString(frame.Function, "")
	fileName := fmt.Sprintf("%s:%d", tempName, frame.Line)
	return "", fileName
}

func GetLogger() *logrus.Logger {
	return log
}

func Trace(args ...interface{}) {
	log.Trace(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}
func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Traceln(args ...interface{}) {
	log.Traceln(args...)
}

func Debugln(args ...interface{}) {
	log.Debugln(args...)
}
func Infoln(args ...interface{}) {
	log.Infoln(args...)
}
func Println(args ...interface{}) {
	log.Println(args...)
}
func Warnln(args ...interface{}) {
	log.Warnln(args...)
}
func Warningln(args ...interface{}) {
	log.Warningln(args...)
}
func Errorln(args ...interface{}) {
	log.Errorln(args...)
}
func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}
func Panicln(args ...interface{}) {
	log.Panicln(args...)
}
