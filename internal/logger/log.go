package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func NewLogger(cfg Config) *zap.Logger {
	var lvl zapcore.Level

	if err := lvl.Set(cfg.Level); err != nil {
		log.Printf("cannot parse log level %s: %s", cfg.Level, err)

		lvl = zapcore.WarnLevel
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), lvl)
	cores := []zapcore.Core{
		defaultCore,
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger
}
