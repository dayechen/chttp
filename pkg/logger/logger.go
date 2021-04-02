package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Engine struct{}

func NewLogger(path string, l string) (*Engine, error) {
	log.SetReportCaller(true)
	str, _ := os.Getwd()
	path = str + "/" + path
	var logFile *os.File
	var err error
	logFile, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	level := map[string]log.Level{
		"debug": log.DebugLevel,
		"info":  log.InfoLevel,
		"warn":  log.WarnLevel,
		"error": log.ErrorLevel,
		"fatal": log.FatalLevel,
		"panic": log.PanicLevel,
	}
	log.SetLevel(level[l])
	log.SetOutput(io.MultiWriter(logFile))
	log.Debug("草")
	return &Engine{}, nil
}

func (l *Engine) Debug(args ...interface{}) {
	log.Debug(args...)
}
func (l *Engine) Info(args ...interface{}) {
	log.Info(args...)
}
func (l *Engine) Warn(args ...interface{}) {
	log.Warn(args...)
}
func (l *Engine) Error(args ...interface{}) {
	log.Error(args...)
}
