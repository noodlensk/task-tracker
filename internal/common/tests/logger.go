package tests

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.DisableCaller = true
	zapConfig.DisableStacktrace = true
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.OutputPaths = []string{}
	zapConfig.ErrorOutputPaths = []string{}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return zapLogger.Sugar()
}
