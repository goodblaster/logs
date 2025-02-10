package formatters

import (
	"fmt"
	"slices"
	"strings"

	"github.com/goodblaster/logs"
)

// Just like the text formatter, but with colors
type consoleFormatter struct {
	cfg Config
}

func NewConsoleFormatter(cfg Config) Formatter {
	return &consoleFormatter{cfg: cfg}
}

func (f consoleFormatter) Format(level logs.Level, msg string, fields Fields) string {
	var tuples []string
	for key, value := range fields {
		tuples = append(tuples, fmt.Sprintf("%s=%v", key, value))
	}
	slices.Sort(tuples)

	// ANSI color codes
	reset := "\033[0m" // Reset color to default
	var color string
	switch level {
	case logs.LevelError:
		color = "\033[31m" // Red
	case logs.LevelWarn:
		color = "\033[33m" // Yellow
	case logs.LevelInfo:
		color = "\033[32m" // Green
	case logs.LevelDebug:
		color = "\033[34m" // Blue
	default:
		color = reset // Reset
	}

	return fmt.Sprintf("%s\t%s%s%s\t%s\t%s", f.cfg.Timestamp(), color, level.String(), reset, strings.Join(tuples, " "), msg)
}
