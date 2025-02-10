package logs

func NewNoopLogger(level Level, format Format) Interface {
	return &NoopLogger{}
}

type NoopLogger struct{}

func (logger NoopLogger) With(key string, value any) Interface       { return logger }
func (logger NoopLogger) WithFields(fields map[string]any) Interface { return logger }
func (logger NoopLogger) Print(format string, args ...any)           {}
func (logger NoopLogger) Debug(format string, args ...any)           {}
func (logger NoopLogger) Info(format string, args ...any)            {}
func (logger NoopLogger) Warn(format string, args ...any)            {}
func (logger NoopLogger) Error(format string, args ...any)           {}
func (logger NoopLogger) Fatal(format string, args ...any)           {}
