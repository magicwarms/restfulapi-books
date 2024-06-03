package utils

import (
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields map[string]interface{}

type Logger interface {
	Info(ctx echo.Context, msg string, fields Fields)
	Error(ctx echo.Context, msg string, fields Fields)
}

type appLogger struct {
	logger *zap.Logger
}

func InitLogger() Logger {
	onlyProcessLogOnce := sync.Once{}

	var zapLogger *zap.Logger
	onlyProcessLogOnce.Do(func() {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		config.ConsoleSeparator = " "

		fileEncoder := zapcore.NewJSONEncoder(config)
		consoleEncoder := zapcore.NewConsoleEncoder(config)

		logFile, _ := os.OpenFile("./server/logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		writer := zapcore.AddSync(logFile)

		defaultLogLevel := zapcore.DebugLevel

		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)
		zapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})

	return &appLogger{
		logger: zapLogger,
	}
}

func (l *appLogger) Info(ctx echo.Context, msg string, fields Fields) {
	zapFields := l.getFields(ctx, fields)
	l.logger.Info(msg, zapFields...)
}

func (l *appLogger) Error(ctx echo.Context, msg string, fields Fields) {
	zapFields := l.getFields(ctx, fields)
	l.logger.Error(msg, zapFields...)
}

func (l *appLogger) getFields(ctx echo.Context, fields Fields) []zap.Field {
	zapFields := []zap.Field{}

	if fields == nil {
		return zapFields
	}

	fields["trace_id"] = ctx.Response().Header().Get(echo.HeaderXRequestID)
	fields["time"] = time.Now().Local()

	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}

	return zapFields
}
