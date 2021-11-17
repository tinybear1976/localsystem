package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var LogContainer map[string]*zap.Logger

func init() {
	LogContainer = make(map[string]*zap.Logger)
}

func NewLogger(tag, logFilenameWithPath, loglevel string) { //*zap.Logger {
	hook := lumberjack.Logger{
		Filename:   logFilenameWithPath,
		MaxSize:    32,
		MaxBackups: 30,
		MaxAge:     90,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	switch loglevel {
	case "debug":
		atomicLevel.SetLevel(zap.DebugLevel)
	case "info":
		atomicLevel.SetLevel(zap.InfoLevel)
	case "error":
		atomicLevel.SetLevel(zap.ErrorLevel)
	case "warn":
		atomicLevel.SetLevel(zap.WarnLevel)
	default:
		atomicLevel.SetLevel(zap.DebugLevel)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
		atomicLevel,
	)

	caller := zap.AddCaller()
	development := zap.Development()
	Log := zap.New(core, caller, development)
	LogContainer[tag] = Log
	//Log.Info("Log initialization successful")
	//return Log
}
