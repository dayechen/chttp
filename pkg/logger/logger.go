package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct{}

func NewLogger(path string, l string) (*Logger, error) {
	log.SetReportCaller(true)
	str, _ := os.Getwd()
	path = str + "/" + path
	var logFile *os.File
	var err error
	logFile, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	level := map[string]log.Level{
		"debug": log.DebugLevel,
		"info":  log.InfoLevel,
		"warn":  log.WarnLevel,
		"error": log.ErrorLevel,
	}
	log.SetLevel(level[l])
	log.SetOutput(io.MultiWriter(logFile))
	log.Info("日志初始化成功")
	return &Logger{}, nil
}

func (l *Logger) Debug(args ...interface{}) {
	log.Debug(args...)
}
func (l *Logger) Info(args ...interface{}) {
	log.Info(args...)
}
func (l *Logger) Warn(args ...interface{}) {
	log.Warn(args...)
}
func (l *Logger) Error(args ...interface{}) {
	log.Error(args...)
}
