package adapters

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/levels"
)

const (
	LevelFatal = slog.Level(10)
	LevelPanic = slog.Level(12)
	LevelPrint = slog.Level(14)
)

func Slog(logger *slog.Logger) *SLogAdapter {
	return &SLogAdapter{logger}
}

type SLogAdapter struct {
	logger *slog.Logger
}

func (adapter SLogAdapter) Level() levels.Level {
	return levels.Debug // TODO: Implement
}

func (adapter SLogAdapter) With(key string, value any) logs.Interface {
	return &SLogAdapter{adapter.logger.With(key, value)}
}

func (adapter SLogAdapter) WithFields(fields map[string]any) logs.Interface {
	var params []any
	for k, v := range fields {
		params = append(params, k, v)
	}
	return &SLogAdapter{adapter.logger.With(params...)}
}

func (adapter SLogAdapter) Log(level levels.Level, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Log(context.Background(), ToSLogLevel(level), msg)
}

func (adapter SLogAdapter) LogFunc(level levels.Level, msg func() string) {
	if level > adapter.Level() {
		return
	}
	adapter.logger.Log(context.Background(), ToSLogLevel(level), msg())
}

func (adapter SLogAdapter) Print(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Log(context.Background(), LevelPrint, msg)
}

func (adapter SLogAdapter) Debug(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Debug(msg)
}

func (adapter SLogAdapter) Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Info(msg)
}

func (adapter SLogAdapter) Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Warn(msg)
}

func (adapter SLogAdapter) Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Error(msg)
}

func (adapter SLogAdapter) Fatal(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Log(context.Background(), LevelFatal, msg)
	os.Exit(-1)
}

func (adapter SLogAdapter) Panic(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Log(context.Background(), LevelPanic, msg)
	panic(msg)
}

// CustomSLogHandler - Custom Handler to Override Level Formatting
type CustomSLogHandler struct {
	slog.Handler
}

func (h *CustomSLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level == LevelPrint {
		r = slog.Record{
			Time:    r.Time,
			Level:   r.Level,
			Message: r.Message,
		}
		// Override level format
		r.AddAttrs(slog.String("level", "print"))
	}
	return h.Handler.Handle(ctx, r)
}

// WithAttrs returns a new [JSONHandler] whose attributes consists
// of h's attributes followed by attrs.
func (h *CustomSLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CustomSLogHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *CustomSLogHandler) WithGroup(name string) slog.Handler {
	return &CustomSLogHandler{Handler: h.Handler.WithGroup(name)}
}

func ToSLogLevel(level levels.Level) slog.Level {
	switch level {
	case levels.Debug:
		return slog.LevelDebug
	case levels.Info:
		return slog.LevelInfo
	case levels.Warn:
		return slog.LevelWarn
	case levels.Error:
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
