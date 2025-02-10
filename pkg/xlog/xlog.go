package xlog

import (
	"fmt"
	"io"
	"sync"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/pkg/xlog/formatters"
)

type Fields = map[string]any

type Level int

// Default levels
var (
	LevelDebug  Level = 0
	LevelInfo   Level = 1
	LevelWarn   Level = 2
	LevelError  Level = 3
	LevelDPanic Level = 4
	LevelPanic  Level = 5
	LevelFatal  Level = 6
	LevelPrint  Level = 10
)

// LevelNames - change if you like.
var LevelNames map[Level]string

// LevelOrder - change if you like.
var LevelOrder []Level

// init - set defaults
func init() {
	LevelNames = map[Level]string{
		LevelDebug:  "debug",
		LevelInfo:   "info",
		LevelWarn:   "warn",
		LevelError:  "error",
		LevelDPanic: "dpanic",
		LevelPanic:  "panic",
		LevelFatal:  "fatal",
		LevelPrint:  "print",
	}

	LevelOrder = []Level{
		LevelDebug,
		LevelInfo,
		LevelWarn,
		LevelError,
		LevelDPanic,
		LevelPanic,
		LevelFatal,
		LevelPrint,
	}
}

type Logger struct {
	level     logs.Level
	formatter formatters.Formatter
	writer    io.Writer
	sync      *sync.Mutex
	fields    Fields
}

func NewLogger(level logs.Level, format logs.Format, writer io.Writer) logs.Interface {
	return &Logger{
		level:     level,
		formatter: formatters.NewFormatter(format),
		writer:    writer,
		sync:      &sync.Mutex{},
		fields:    nil,
	}
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

func (logger Logger) Log(level logs.Level, format string, args ...any) {
	if logger.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	line := logger.formatter.Format(level, msg, logger.fields)
	_, _ = fmt.Fprintln(logger.writer, line)
}

// LogFunc - use for expensive operations where you don't want to calculate the message if the level is not enabled.
func (logger Logger) LogFunc(level logs.Level, msg func() string) {
	if logger.level < level {
		return
	}
	line := logger.formatter.Format(level, msg(), logger.fields)
	_, _ = fmt.Fprintln(logger.writer, line)
}

func (logger Logger) Print(format string, args ...any) {
	logger.Log(logs.LevelPrint, format, args...)
}

func (logger Logger) Debug(format string, args ...any) {
	logger.Log(logs.LevelDebug, format, args...)
}

func (logger Logger) Info(format string, args ...any) {
	logger.Log(logs.LevelInfo, format, args...)
}

func (logger Logger) Warn(format string, args ...any) {
	logger.Log(logs.LevelWarn, format, args...)
}

func (logger Logger) Error(format string, args ...any) {
	logger.Log(logs.LevelError, format, args...)
}

func (logger Logger) Fatal(format string, args ...any) {
	logger.Log(logs.LevelFatal, format, args...)
	panic(fmt.Sprintf(format, args...))
}
