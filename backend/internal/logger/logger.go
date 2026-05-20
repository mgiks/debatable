package logger

type Logger interface {
	Info(msg string, keysAndVals ...any)
	Error(msg string, keysAndVals ...any)
}
