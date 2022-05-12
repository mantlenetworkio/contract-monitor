package logger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func getHook(filename string, maxAge int) (io.Writer, error) {
	hook := &lumberjack.Logger{
		Filename:   filename, // 日志文件名
		MaxSize:    100,      // megabytes
		MaxBackups: 1,
		MaxAge:     maxAge, // days
		Compress:   false,  // disabled by default
	}

	return hook, nil
}
