package logger

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Error(err error, msg string)
	Fatal(err error, msg string)
}
