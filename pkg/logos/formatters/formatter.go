package formatters

import (
	"time"

	"github.com/goodblaster/logs"
)

type Fields = map[string]any

type Formatter interface {
	Format(level logs.Level, msg string, fields Fields) string
}

type Config struct {
	Timestamp func() string
}

var DefaultConfig = Config{
	Timestamp: DefaultTimestamp,
}

func NewFormatter(format logs.Format) Formatter {
	return NewFormatterWithConfig(format, DefaultConfig)
}

func NewFormatterWithConfig(format logs.Format, cfg Config) Formatter {
	switch format {
	case logs.FormatJSON:
		return NewJsonFormatter(cfg)
	case logs.FormatText:
		return NewTextFormatter(cfg)
	case logs.FormatConsole:
		return NewConsoleFormatter(cfg)
	}
	return &jsonFormatter{}
}

const DefaultTimestampFormat = "2006-01-02T15:04:05"

func DefaultTimestamp() string {
	return time.Now().Local().Format(DefaultTimestampFormat)
}
