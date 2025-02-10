package adapters

import (
	"fmt"

	"github.com/goodblaster/logs"
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
