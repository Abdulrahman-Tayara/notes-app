package logging

type ILogger interface {
	Info(message string, args ...any)

	Warning(message string, args ...any)

	Debug(message string, args ...any)

	Error(err error)

	Errorf(err string, args ...any)
}
