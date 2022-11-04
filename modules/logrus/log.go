package logrus

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	// 日志文件的保存时间
	maxAge       = 30 * time.Hour * 24
	rotationTime = time.Hour * 24
	timeFormat   = "2006-01-02 15:04:05"
)

var L logrus.Logger

func init() {
	//控制台logger
	L = logrus.Logger{
		Out:   os.Stdout,
		Hooks: make(logrus.LevelHooks),
		Formatter: &logrus.TextFormatter{
			ForceColors:               true,
			ForceQuote:                true,
			EnvironmentOverrideColors: true,
			FullTimestamp:             true,
			TimestampFormat:           timeFormat,
			DisableSorting:            false,
			PadLevelText:              true,
			QuoteEmptyFields:          true,
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				return frame.Function[strings.LastIndex(frame.Function, "/")+1:],
					fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
			},
		},
		ReportCaller: true,
		Level:        logrus.InfoLevel,
		ExitFunc: func(i int) {
			fmt.Println("the app will exit...")
			os.Exit(i)
		},
	}
	filesMap := lfshook.WriterMap{}
	for k, v := range map[logrus.Level]string{
		logrus.InfoLevel:  "info",
		logrus.WarnLevel:  "warning",
		logrus.ErrorLevel: "error",
		logrus.FatalLevel: "fatal",
	} {

		f, err := rotatelogs.New(
			fmt.Sprintf("logs/%s-%s.log", "%Y-%m-%d", v),
			rotatelogs.WithMaxAge(maxAge),
			rotatelogs.WithRotationTime(rotationTime),
		)
		if err != nil {
			log.Fatal(err)
		}
		filesMap[k] = f
	}

	filesMap[logrus.TraceLevel] = filesMap[logrus.InfoLevel]
	filesMap[logrus.DebugLevel] = filesMap[logrus.InfoLevel]

	//文件日志
	L.AddHook(lfshook.NewHook(filesMap, &TextFormatter{
		ForceQuote:       false,
		DisableQuote:     true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
		TimestampFormat:  timeFormat,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fmt.Println(frame.File, frame.Line, frame.Function)
			return frame.Function[strings.LastIndex(frame.Function, "/")+1:],
				fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		},
	}))
}
