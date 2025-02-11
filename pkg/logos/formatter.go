package logos

import (
	"time"

	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
)

type Formatter interface {
	Format(level levels.Level, msg string, fields Fields) string
}

type Config struct {
	Timestamp func() string
}

var DefaultConfig = Config{
	Timestamp: DefaultTimestamp,
}

func NewFormatter(format formats.Format) Formatter {
	return NewFormatterWithConfig(format, DefaultConfig)
}

func NewFormatterWithConfig(format formats.Format, cfg Config) Formatter {
	switch format {
	case formats.JSON:
		return NewJsonFormatter(cfg)
	case formats.Text:
		return NewTextFormatter(cfg)
	case formats.Console:
		return NewConsoleFormatter(cfg)
	}
	panic("unknown format")
}

const DefaultTimestampFormat = "2006-01-02T15:04:05"

func DefaultTimestamp() string {
	return time.Now().Local().Format(DefaultTimestampFormat)
}
