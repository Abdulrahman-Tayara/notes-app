package logging

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	zapLogger *zap.SugaredLogger
}

func NewDevelopment() ILogger {
	logger, _ := zap.NewDevelopment(zap.AddCallerSkip(1))

	return &ZapLogger{
		zapLogger: logger.Sugar(),
	}
}

func NewProduction() ILogger {
	logger, _ := zap.NewProduction(zap.AddCallerSkip(1))

	return &ZapLogger{
		zapLogger: logger.Sugar(),
	}
}

func (z *ZapLogger) Info(message string, args ...any) {
	z.zapLogger.Infof(message, args...)
}

func (z *ZapLogger) Warning(message string, args ...any) {
	z.zapLogger.Warnf(message, args)
}

func (z *ZapLogger) Debug(message string, args ...any) {
	z.zapLogger.Debugf(message, args...)
}

func (z *ZapLogger) Error(err error) {
	z.zapLogger.Error(err)
}

func (z *ZapLogger) Errorf(err string, args ...any) {
	z.zapLogger.Errorf(err, args...)
}

func (z *ZapLogger) Fatalf(err string, args ...any) {
	z.zapLogger.Fatalf(err, args...)
}
