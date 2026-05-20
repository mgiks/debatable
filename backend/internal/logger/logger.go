package logger

type Logger interface {
	Info(msg string, keysAndVals ...any)
	Error(msg string, keysAndVals ...any)
	Fatal(msg string, keysAndVals ...any)
}
