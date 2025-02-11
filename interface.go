package logs

import "github.com/goodblaster/logs/levels"

func init() {
	DefaultLogger = SimpleLogger{}
}

type Interface interface {
	SetLevel(level levels.Level)
	Log(level levels.Level, format string, args ...any)
	LogFunc(level levels.Level, f func() string)
	With(key string, value any) Interface
	WithFields(fields map[string]any) Interface
	Print(format string, args ...any)
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

var DefaultLogger Interface

func SetDefaultLogger(logger Interface) {
	DefaultLogger = logger
}

func SetLevel(level levels.Level) {
	DefaultLogger.SetLevel(level)
}

func With(key string, value any) Interface {
	return DefaultLogger.With(key, value)
}

func WithError(err error) Interface {
	return DefaultLogger.With("error", err)
}

func WithFields(fields map[string]any) Interface {
	return DefaultLogger.WithFields(fields)
}

func Log(level levels.Level, format string, args ...any) {
	DefaultLogger.Log(level, format, args...)
}

func LogFunc(level levels.Level, msg func() string) {
	DefaultLogger.LogFunc(level, msg)
}

func Print(format string, args ...any) {
	DefaultLogger.Print(format, args...)
}

func Debug(format string, args ...any) {
	DefaultLogger.Debug(format, args...)
}

func Info(format string, args ...any) {
	DefaultLogger.Info(format, args...)
}

func Warn(format string, args ...any) {
	DefaultLogger.Warn(format, args...)
}

func Error(format string, args ...any) {
	DefaultLogger.Error(format, args...)
}

func Fatal(format string, args ...any) {
	DefaultLogger.Fatal(format, args...)
}
