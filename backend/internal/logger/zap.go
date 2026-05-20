package logger

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func (l zapLogger) Info(msg string, keysAndVals ...any) {
	l.logger.Infow(msg, keysAndVals...)
}

func (l zapLogger) Error(msg string, keysAndVals ...any) {
	l.logger.Errorw(msg, keysAndVals...)
}

func (l zapLogger) Fatal(msg string, keysAndVals ...any) {
	l.logger.Fatalw(msg, keysAndVals...)
}

func NewZapLogger(env string) Logger {
	loggerConfig := zap.NewProductionConfig()
	if env == "development" {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	logger := zap.Must(loggerConfig.Build(zap.AddCallerSkip(1))).Sugar()

	return zapLogger{
		logger: logger,
	}
}
