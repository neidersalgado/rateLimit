// pkg/Logger/zaplogger.go

package logger

import (
	"go.uber.org/zap"
	"os"
)

type ZapLogger struct {
	Logger *zap.Logger
	env    string
}

func NewZapLogger() *ZapLogger {
	var logger *zap.Logger
	var err error
	env := os.Getenv("ENV")

	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err) // No se puede proceder sin un Logger
	}

	return &ZapLogger{
		Logger: logger,
		env:    env,
	}
}

func (zl *ZapLogger) Debug(msg string, fields ...zap.Field) {
	if zl.env != "production" {
		zl.Logger.Debug(msg, fields...)
	}
}

func (zl *ZapLogger) Info(msg string, fields ...zap.Field) {
	zl.Logger.Info(msg, fields...)
}

func (zl *ZapLogger) Error(msg string, fields ...zap.Field) {
	zl.Logger.Error(msg, fields...)
}
