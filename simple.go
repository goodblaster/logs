package logs

import (
	"fmt"
	"os"

	"github.com/goodblaster/logs/pkg/formats"
	"github.com/goodblaster/logs/pkg/levels"
)

func NewSimpleLogger(level levels.Level, format formats.Format) Logger {
	return &SimpleLogger{}
}

type SimpleLogger struct{}

func (logger SimpleLogger) With(key string, value any) Logger {
	fmt.Println("WITH", key, value)
	return logger
}
func (logger SimpleLogger) WithFields(fields map[string]any) Logger {
	fmt.Println("WITH_FIELDS", fields)
	return logger
}

func (logger SimpleLogger) Print(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("PRINT", msg)
}

func (logger SimpleLogger) Debug(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("DEBUG", msg)
}

func (logger SimpleLogger) Info(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("INFO", msg)
}

func (logger SimpleLogger) Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("WARN", msg)
}

func (logger SimpleLogger) Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("ERROR", msg)
}

func (logger SimpleLogger) Fatal(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("FATAL", msg)
	os.Exit(-1)
}
