package logs

import (
	"github.com/goodblaster/logs/pkg/formats"
	"github.com/goodblaster/logs/pkg/levels"
)

func NewNoopLogger(level levels.Level, format formats.Format) Logger {
	return &NoopLogger{}
}

type NoopLogger struct{}

func (logger NoopLogger) With(key string, value any) Logger       { return logger }
func (logger NoopLogger) WithFields(fields map[string]any) Logger { return logger }
func (logger NoopLogger) Print(format string, args ...any)        {}
func (logger NoopLogger) Debug(format string, args ...any)        {}
func (logger NoopLogger) Info(format string, args ...any)         {}
func (logger NoopLogger) Warn(format string, args ...any)         {}
func (logger NoopLogger) Error(format string, args ...any)        {}
func (logger NoopLogger) Fatal(format string, args ...any)        {}
