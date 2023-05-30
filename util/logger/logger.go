package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func Init() error {
	conf := zap.Config{
		Level:    zap.NewAtomicLevel(),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "name",
			CallerKey:     "caller",
			FunctionKey:   "function",
			StacktraceKey: "stackTrace",
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},
	}

	var err error
	Log, err = conf.Build()

	return err
}
