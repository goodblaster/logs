package adapters

import (
	"fmt"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/levels"
	"github.com/sirupsen/logrus"
)

// LogrusAdapter wraps a logrus.Logger to implement the logs.Interface interface.
type LogrusAdapter struct {
	logger *logrus.Entry
}

func (adapter LogrusAdapter) Level() levels.Level {
	return levels.Level(adapter.logger.Level)
}

func Logrus(logger *logrus.Logger) *LogrusAdapter {
	return &LogrusAdapter{logrus.NewEntry(logger)}
}

func (adapter LogrusAdapter) With(key string, value any) logs.Interface {
	return &LogrusAdapter{adapter.logger.WithField(key, value)}
}

func (adapter LogrusAdapter) WithFields(fields map[string]any) logs.Interface {
	return &LogrusAdapter{adapter.logger.WithFields(logrus.Fields(fields))}
}

func (adapter LogrusAdapter) Log(level levels.Level, format string, args ...any) {
	switch level {
	case levels.Print:
		adapter.Print(format, args...)
	case levels.Debug:
		adapter.Debug(format, args...)
	case levels.Info:
		adapter.Info(format, args...)
	case levels.Warn:
		adapter.Warn(format, args...)
	case levels.Error:
		adapter.Error(format, args...)
	case levels.Panic:
		adapter.Panic(format, args...)
	case levels.Fatal:
		adapter.Fatal(format, args...)
	}
}

func (adapter LogrusAdapter) LogFunc(level levels.Level, msg func() string) {
	if level > adapter.Level() {
		return
	}
	switch level {
	case levels.Print:
		adapter.Print(msg())
	case levels.Debug:
		adapter.Debug(msg())
	case levels.Info:
		adapter.Info(msg())
	case levels.Warn:
		adapter.Warn(msg())
	case levels.Error:
		adapter.Error(msg())
	case levels.Panic:
		adapter.Panic(msg())
	case levels.Fatal:
		adapter.Fatal(msg())
	}
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

func ToLogrusLevel(level levels.Level) logrus.Level {
	switch level {
	case levels.Debug:
		return logrus.DebugLevel
	case levels.Info:
		return logrus.InfoLevel
	case levels.Warn:
		return logrus.WarnLevel
	case levels.Error:
		return logrus.ErrorLevel
	case levels.Fatal:
		return logrus.FatalLevel
	case levels.Panic:
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

func FromLogrusLevel(level logrus.Level) levels.Level {
	switch level {
	case logrus.DebugLevel:
		return levels.Debug
	case logrus.InfoLevel:
		return levels.Info
	case logrus.WarnLevel:
		return levels.Warn
	case logrus.ErrorLevel:
		return levels.Error
	case logrus.FatalLevel:
		return levels.Fatal
	case logrus.PanicLevel:
		return levels.Panic
	default:
		return levels.Debug
	}
}

// CustomFormatter forces uppercase log levels
type CustomFormatter struct {
	logrus.Formatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return f.Formatter.Format(entry)
}
