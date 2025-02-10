package adapters

import (
	"fmt"

	"github.com/goodblaster/logs"
	"github.com/sirupsen/logrus"
)

// LogrusAdapter wraps a logrus.Logger to implement the logs.Interface interface.
type LogrusAdapter struct {
	logger *logrus.Entry
}

func Logrus(logger *logrus.Logger) *LogrusAdapter {
	return &LogrusAdapter{logrus.NewEntry(logger)}
}

func (adapter LogrusAdapter) With(key string, value any) *LogrusAdapter {
	return &LogrusAdapter{adapter.logger.WithField(key, value)}
}

func (adapter LogrusAdapter) WithFields(fields map[string]any) *LogrusAdapter {
	return &LogrusAdapter{adapter.logger.WithFields(logrus.Fields(fields))}
}

func (adapter LogrusAdapter) Print(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Println(msg)
}

func (adapter LogrusAdapter) Debug(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Debug(msg)
}

func (adapter LogrusAdapter) Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Info(msg)
}

func (adapter LogrusAdapter) Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Warn(msg)
}

func (adapter LogrusAdapter) Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Error(msg)
}

func (adapter LogrusAdapter) Fatal(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Fatal(msg)
}

func (adapter LogrusAdapter) Panic(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	adapter.logger.Panic(msg)
}

func ToLogrusLevel(level logs.Level) logrus.Level {
	switch level {
	case logs.LevelDebug:
		return logrus.DebugLevel
	case logs.LevelInfo:
		return logrus.InfoLevel
	case logs.LevelWarn:
		return logrus.WarnLevel
	case logs.LevelError:
		return logrus.ErrorLevel
	default:
		return logrus.DebugLevel
	}
}

// CustomFormatter forces uppercase log levels
type CustomFormatter struct {
	logrus.Formatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return f.Formatter.Format(entry)
}
