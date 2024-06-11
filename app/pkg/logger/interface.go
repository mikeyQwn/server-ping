package logger

type Logger interface {
	Info(msg string)
	Infof(format string, v any)
	Debug(msg string)
	Debugf(format string, v any)
	Warn(msg string)
	Warnf(format string, v any)
	Error(err error, msg string)
	Errorf(err error, format string, v any)
	Fatal(err error, msg string)
	Fatalf(err error, format string, v any)
}
