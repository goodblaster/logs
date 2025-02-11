package formatters

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/goodblaster/logs/colors"
	"github.com/goodblaster/logs/levels"
)

// Just like the text formatter, but with colors
type consoleFormatter struct {
	cfg Config
}

func NewConsoleFormatter(cfg Config) Formatter {
	return &consoleFormatter{cfg: cfg}
}

func (f consoleFormatter) Format(level levels.Level, msg string, fields map[string]any) string {
	var tuples []string
	for key, value := range fields {
		b, _ := json.Marshal(value)
		tuples = append(tuples, fmt.Sprintf("%s=%v", key, string(b)))
	}
	slices.Sort(tuples)

	// ANSI color codes
	textColor := levels.LevelColors[level]

	return fmt.Sprintf("%s\t%s%s%s\t%s\t%s", f.cfg.Timestamp(), textColor, level.String(), colors.Reset, strings.Join(tuples, " "), msg)
}
