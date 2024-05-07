package observability

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type LogContext struct {
	parameters []zap.Field
}

func (context *LogContext) Add(fields ...zap.Field) {
	for _, field := range fields {
		context.parameters = append(context.parameters, field)
	}
}

func Log[T any](message string, block func(context *LogContext) (T, error)) (T, error) {
	logger := logger()
	defer logger.Sync()

	context := &LogContext{parameters: []zap.Field{}}

	startTime := time.Now()
	result, err := block(context)
	duration := time.Since(startTime)

	context.Add(zap.Int64("elapsed_time", duration.Milliseconds()))

	if err != nil {
		logger.Error(message, context.parameters...)
		return result, err
	}

	logger.Info(message, context.parameters...)
	return result, err
}

func logger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	return logger
}
