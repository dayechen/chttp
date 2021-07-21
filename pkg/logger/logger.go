package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// var levelMap = map[string]zap.AtomicLevel{
// 	"debug": zap.NewAtomicLevelAt(zap.DebugLevel),
// 	"info":  zap.NewAtomicLevelAt(zap.InfoLevel),
// 	"warn":  zap.NewAtomicLevelAt(zap.WarnLevel),
// 	"error": zap.NewAtomicLevelAt(zap.ErrorLevel),
// 	"fatal": zap.NewAtomicLevelAt(zap.FatalLevel),
// }

func NewLogger(path string, level string, runMode string) (*zap.SugaredLogger, error) {

	levelMap := map[string]zapcore.Level{
		"debug": zap.DebugLevel,
		"info":  zap.InfoLevel,
		"warn":  zap.WarnLevel,
		"error": zap.ErrorLevel,
		"fatal": zap.FatalLevel,
	}

	hook := lumberjack.Logger{
		Filename:   path, // 日志文件路径
		MaxSize:    128,  // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,   // 日志文件最多保存多少个备份
		MaxAge:     7,    // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(levelMap[level])

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),                                        // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development)

	return logger.Sugar(), nil
}

// func NewLogger1(path string, level string, runMode string) (*zap.SugaredLogger, error) {

// 	var (
// 		logger *zap.Logger
// 		err    error
// 	)

// 	encoderConfig := zapcore.EncoderConfig{
// 		TimeKey:        "time", // 这一堆只有在json格式才有用
// 		LevelKey:       "level",
// 		NameKey:        "logger",
// 		CallerKey:      "caller",
// 		MessageKey:     "msg",
// 		StacktraceKey:  "stacktrace",
// 		LineEnding:     zapcore.DefaultLineEnding,
// 		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
// 		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
// 		EncodeDuration: zapcore.SecondsDurationEncoder,
// 		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
// 	}

// 	config := zap.Config{
// 		Level:         levelMap[level],
// 		Encoding:      "console",
// 		OutputPaths:   []string{path},
// 		EncoderConfig: encoderConfig,
// 	}

// 	if runMode != "debug" {
// 		logger, err = zap.NewDevelopment()
// 	} else {
// 		logger, err = config.Build()
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer logger.Sync()

// 	return logger.Sugar(), nil
// }
