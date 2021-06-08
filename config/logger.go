package config

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	// Log Log
	Log *logrus.Logger
)

// InitLogger InitLogger
func InitLogger() {
	Log = logrus.New()
	hook := NewLfsHook(filepath.Join(Cfg.Logger.Path, Cfg.App.Name), 0, 5)
	Log.AddHook(hook)
	Log.SetFormatter(formatter(true))
	Log.SetReportCaller(true)
}

func formatter(isConsole bool) *nested.Formatter {
	fmtter := &nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerFirst:     true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error during runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			return fmt.Sprintf(" [%v:%v]", filepath.Base(fullPath), line)
		},
	}
	if runtime.GOOS == "windows" {
		fmtter.NoColors = true
	} else {
		fmtter.NoColors = false
	}
	return fmtter
}

// NewLfsHook NewLfsHook
func NewLfsHook(logName string, rotationTime time.Duration, leastDay uint) logrus.Hook {
	writer, err := rotatelogs.New(
		// 日志文件
		logName+".%Y%m%d",
		// 日志周期(默认每86400秒/一天旋转一次)
		//rotatelogs.WithRotationTime(rotationTime),
		// 清除历史 (WithMaxAge和WithRotationCount只能选其一)
		//rotatelogs.WithMaxAge(time.Hour*24*7), //默认每7天清除下日志文件
		rotatelogs.WithRotationCount(leastDay), //只保留最近的N个日志文件
	)
	if err != nil {
		log.Fatal(err)
	}
	// 可设置按不同level创建不同的文件名
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
		//}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	}, formatter(false))
	return lfsHook
}
