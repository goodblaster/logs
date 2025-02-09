package logs

func init() {
	DefaultLogger = SimpleLogger{}
}

type Logger interface {
	With(key string, value any) Logger
	WithFields(fields map[string]any) Logger
	Print(format string, args ...any)
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

var DefaultLogger Logger

func SetDefaultLogger(logger Logger) {
	DefaultLogger = logger
}

func With(key string, value any) Logger {
	return DefaultLogger.With(key, value)
}

func WithError(err error) Logger {
	return DefaultLogger.With("error", err)
}

func WithFields(fields map[string]any) Logger {
	return DefaultLogger.WithFields(fields)
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
