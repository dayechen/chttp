package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
)

// Level 错误等级
type Level int8

// Fields 错误字段
type Fields map[string]interface{}

// Debug 错误等级
const (
	Debug Level = iota
	Info
	Warn
	Error
)

// String 将错误码转换成字符串
func (l Level) String() string {
	switch l {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	}
	return ""
}

// Logger 日志结构体
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

// NewLogger 创建日志
func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

// clone 复制日志
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// WithLevel 设置日志等级
func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

// WithFields 设置日志公用字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

// WithContext 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller 设置当前的调用信息（文件信息行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}
	return ll
}

// WithCallersFrames 设置当前整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 7 // 最大追踪的层级
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// JSONFormat 格式化日志信息
func (l *Logger) JSONFormat(message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = l.level.String()
	// data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.WithCallersFrames().callers[4:]
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

// Output 输出日志
func (l *Logger) Output(message string) {
	body, _ := json.Marshal(l.JSONFormat(message))
	content := string(body)
	switch l.level {
	case Debug:
		l.newLogger.Print(content)
	case Info:
		l.newLogger.Print(content)
	case Warn:
		l.newLogger.Print(content)
	case Error:
		l.newLogger.Print(content)
	}
}

// Debugf 错误日志
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.WithLevel(Debug).Output(fmt.Sprintf(format, v...))
}

// Infof 详情日志
func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithLevel(Info).Output(fmt.Sprintf(format, v...))
}

// Warnf 警告日志
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WithLevel(Warn).Output(fmt.Sprintf(format, v...))
}

// Errorf 错误日志
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WithLevel(Error).Output(fmt.Sprintf(format, v...))
}
