package logos

import (
	"fmt"
	"io"
	"sync"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
	"github.com/goodblaster/logs/pkg/logos/formatters"
)

type Fields = map[string]any

type Logger struct {
	level     *levels.Level
	formatter formatters.Formatter
	writer    io.Writer
	sync      *sync.Mutex
	fields    Fields
}

func NewLogger(level levels.Level, format formats.Format, writer io.Writer) logs.Interface {
	return &Logger{
		level:     &level,
		formatter: formatters.NewFormatter(format),
		writer:    writer,
		sync:      &sync.Mutex{},
		fields:    nil,
	}
}

func (logger Logger) SetLevel(level levels.Level) {
	*logger.level = level
}

func (logger Logger) Copy() Logger {
	logger.sync.Lock()
	defer logger.sync.Unlock()

	newLogger := logger

	if newLogger.fields != nil {
		newLogger.fields = make(Fields)
		for key, value := range logger.fields {
			newLogger.fields[key] = value
		}
	}

	return newLogger
}

func (logger Logger) With(key string, value any) logs.Interface {
	return logger.WithFields(Fields{key: value})
}

func (logger Logger) WithFields(fields Fields) logs.Interface {
	newLogger := logger.Copy()

	if newLogger.fields == nil {
		newLogger.fields = make(Fields)
	}

	for key, value := range fields {
		newLogger.fields[key] = value
	}

	return &newLogger
}

func (logger Logger) Log(level levels.Level, format string, args ...any) {
	if *logger.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	line := logger.formatter.Format(level, msg, logger.fields)
	_, _ = fmt.Fprintln(logger.writer, line)
}

// LogFunc - use for expensive operations where you don't want to calculate the message if the level is not enabled.
func (logger Logger) LogFunc(level levels.Level, msg func() string) {
	if *logger.level > level {
		return
	}
	logger.Log(level, msg())
}

func (logger Logger) Print(format string, args ...any) {
	logger.Log(levels.Print, format, args...)
}

func (logger Logger) Debug(format string, args ...any) {
	logger.Log(levels.Debug, format, args...)
}

func (logger Logger) Info(format string, args ...any) {
	logger.Log(levels.Info, format, args...)
}

func (logger Logger) Warn(format string, args ...any) {
	logger.Log(levels.Warn, format, args...)
}

func (logger Logger) Error(format string, args ...any) {
	logger.Log(levels.Error, format, args...)
}

func (logger Logger) Fatal(format string, args ...any) {
	logger.Log(levels.Fatal, format, args...)
	panic(fmt.Sprintf(format, args...))
}
