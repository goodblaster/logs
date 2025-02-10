package formatters

import (
	"encoding/json"

	"github.com/goodblaster/errors"
	"github.com/goodblaster/logs"
)

type jsonFormatter struct {
	cfg Config
}

func NewJsonFormatter(cfg Config) Formatter {
	return &jsonFormatter{cfg: cfg}
}

func (f jsonFormatter) Format(level logs.Level, msg string, fields Fields) string {
	type Entry struct {
		Level     string `json:"level"`
		Timestamp string `json:"timestamp"`
		Fields    Fields `json:"fields,omitempty"`
		Msg       string `json:"msg"`
	}

	entry := Entry{
		Level:     level.String(),
		Timestamp: f.cfg.Timestamp(),
		Msg:       msg,
		Fields:    fields,
	}

	b, err := json.Marshal(entry)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal log entry"))
	}

	return string(b)
}
