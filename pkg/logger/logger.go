package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[string]zap.AtomicLevel{
	"debug": zap.NewAtomicLevelAt(zap.DebugLevel),
}

func NewLogger(path string, level string, runMode string) (*zap.SugaredLogger, error) {
	fmt.Printf("日志目录 %s ,日志等级 %s \n", path, level)

	var (
		logger *zap.Logger
		err    error
	)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time", // 这一堆只有在json格式才有用
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	config := zap.Config{
		Level:         levelMap[level],
		Encoding:      "console",
		OutputPaths:   []string{path},
		EncoderConfig: encoderConfig,
	}

	if runMode != "debug" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = config.Build()
	}

	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return logger.Sugar(), nil
}
