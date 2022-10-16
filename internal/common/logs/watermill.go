package logs

import (
	"github.com/ThreeDotsLabs/watermill"
	"go.uber.org/zap"
)

type watermillLogger struct {
	logger *zap.SugaredLogger
}

func NewWatermillLogger(logger *zap.SugaredLogger) watermillLogger { //nolint:revive
	return watermillLogger{logger: logger}
}

func NewWatermillNopLogger() watermillLogger { //nolint:revive
	return NewWatermillLogger(NewNopLogger())
}

func (l watermillLogger) Error(msg string, err error, fields watermill.LogFields) {
	var logFields []interface{}

	logFields = append(logFields, "err", err.Error())

	for k, v := range fields {
		logFields = append(logFields, k, v)
	}

	l.logger.Errorw(msg, logFields...)
}

func (l watermillLogger) Info(msg string, fields watermill.LogFields) {
	var logFields []interface{}

	for k, v := range fields {
		logFields = append(logFields, k, v)
	}

	l.logger.Infow(msg, logFields...)
}

func (l watermillLogger) Debug(msg string, fields watermill.LogFields) {
	var logFields []interface{}

	for k, v := range fields {
		logFields = append(logFields, k, v)
	}

	l.logger.Debugw(msg, logFields...)
}

func (l watermillLogger) Trace(msg string, fields watermill.LogFields) {
	var logFields []interface{}

	for k, v := range fields {
		logFields = append(logFields, k, v)
	}

	l.logger.Infow(msg, logFields...)
}

func (l watermillLogger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	var logFields []interface{}

	for k, v := range fields {
		logFields = append(logFields, k, v)
	}

	return watermillLogger{logger: l.logger.With(logFields...)}
}
