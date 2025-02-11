package adapters

import (
	"fmt"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/levels"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// PrintLevel is a custom log level for printing messages in same foramt as log messages regardless of level.
const PrintLevel zapcore.Level = 10

// CustomLevelEncoder ensures our level is properly encoded
func CustomLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	if l == PrintLevel {
		enc.AppendString("print")
	} else {
		enc.AppendString(l.String())
	}
}

// CustomCore wraps zapcore.Core to avoid stack traces for PrintLevel
type CustomCore struct {
	zapcore.Core
}

func (c CustomCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if entry.Level == PrintLevel {
		// Explicitly disable stack traces for PrintLevel
		entry.Stack = ""
	}
	return c.Core.Check(entry, ce)
}

func Zap(logger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{logger}
}

type ZapAdapter struct {
	logger *zap.Logger
}

func (adapter ZapAdapter) Level() levels.Level {
	return levels.Debug // TODO: Implement
}

func (adapter ZapAdapter) SetLevel(level levels.Level) {
	// TODO: Implement
}

func (adapter ZapAdapter) Flush() {
	_ = adapter.logger.Sync()
}

func (adapter ZapAdapter) With(key string, value any) logs.Interface {
	return &ZapAdapter{adapter.logger.With(zap.Any(key, value))}
}

func (adapter ZapAdapter) WithFields(fields map[string]any) logs.Interface {
	var zaps []zap.Field
	for k, v := range fields {
		zaps = append(zaps, zap.Any(k, v))
	}
	return &ZapAdapter{adapter.logger.With(zaps...)}
}

func (adapter ZapAdapter) Log(level levels.Level, format string, args ...any) {
	switch level {
	case levels.Debug:
		adapter.Debug(format, args...)
	case levels.Info:
		adapter.Info(format, args...)
	case levels.Warn:
		adapter.Warn(format, args...)
	case levels.Error:
		adapter.Error(format, args...)
	case levels.DPanic:
		adapter.DPanic(format, args...)
	case levels.Panic:
		adapter.Panic(format, args...)
	case levels.Fatal:
		adapter.Fatal(format, args...)
	case levels.Print:
		adapter.Print(format, args...)
	}
}

func (adapter ZapAdapter) LogFunc(level levels.Level, msg func() string) {
	if level > adapter.Level() {
		return
	}

	switch level {
	case levels.Debug:
		adapter.Debug(msg())
	case levels.Info:
		adapter.Info(msg())
	case levels.Warn:
		adapter.Warn(msg())
	case levels.Error:
		adapter.Error(msg())
	case levels.DPanic:
		adapter.DPanic(msg())
	case levels.Panic:
		adapter.Panic(msg())
	case levels.Fatal:
		adapter.Fatal(msg())
	case levels.Print:
		adapter.Print(msg())
	}
}

func (adapter ZapAdapter) Print(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Log(PrintLevel, msg)
}

func (adapter ZapAdapter) Debug(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Debug(msg)
}

func (adapter ZapAdapter) Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Info(msg)
}

func (adapter ZapAdapter) Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Warn(msg)
}

func (adapter ZapAdapter) Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Error(msg)
}

func (adapter ZapAdapter) Fatal(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Fatal(msg)
}

func (adapter ZapAdapter) Panic(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Panic(msg)
}

func (adapter ZapAdapter) DPanic(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.DPanic(msg)
}

func ToZapLevel(level levels.Level) zapcore.Level {
	switch level {
	case levels.Debug:
		return zap.DebugLevel
	case levels.Info:
		return zap.InfoLevel
	case levels.Warn:
		return zap.WarnLevel
	case levels.Error:
		return zap.ErrorLevel
	case levels.DPanic:
		return zap.DPanicLevel
	case levels.Panic:
		return zap.PanicLevel
	case levels.Fatal:
		return zap.FatalLevel
	case levels.Print:
		return PrintLevel
	}
	return zap.InfoLevel
}

func FromZapLevel(level zapcore.Level) levels.Level {
	switch level {
	case zap.DebugLevel:
		return levels.Debug
	case zap.InfoLevel:
		return levels.Info
	case zap.WarnLevel:
		return levels.Warn
	case zap.ErrorLevel:
		return levels.Error
	case zap.DPanicLevel:
		return levels.DPanic
	case zap.PanicLevel:
		return levels.Panic
	case zap.FatalLevel:
		return levels.Fatal
	case PrintLevel:
		return levels.Print
	}
	return levels.Info
}
