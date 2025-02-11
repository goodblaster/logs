package logs

import (
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
)

func NewNoopLogger(level levels.Level, format formats.Format) Interface {
	return &NoopLogger{}
}

type NoopLogger struct{}

func (logger NoopLogger) With(key string, value any) Interface               { return logger }
func (logger NoopLogger) WithFields(fields map[string]any) Interface         { return logger }
func (logger NoopLogger) Log(level levels.Level, format string, args ...any) {}
func (logger NoopLogger) LogFunc(level levels.Level, f func() string)        {}
func (logger NoopLogger) Print(format string, args ...any)                   {}
func (logger NoopLogger) Debug(format string, args ...any)                   {}
func (logger NoopLogger) Info(format string, args ...any)                    {}
func (logger NoopLogger) Warn(format string, args ...any)                    {}
func (logger NoopLogger) Error(format string, args ...any)                   {}
func (logger NoopLogger) Fatal(format string, args ...any)                   {}
