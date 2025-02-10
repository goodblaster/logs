package contrib

import (
	"io"
	"os"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/pkg/adapters"
	"github.com/sirupsen/logrus"
)

func NewLogrusLogger(level logs.Level, format logs.Format, writer io.Writer) *adapters.LogrusAdapter {
	var formatter logrus.Formatter
	if format == logs.FormatJSON {
		formatter = new(logrus.JSONFormatter)
	} else {
		formatter = new(logrus.TextFormatter)
	}

	// Set custom formatter
	formatter = &adapters.CustomFormatter{
		Formatter: formatter,
	}

	logger := &logrus.Logger{
		Out:          writer,
		Formatter:    formatter,
		Hooks:        make(logrus.LevelHooks),
		Level:        adapters.ToLogrusLevel(level),
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	return adapters.Logrus(logger)
}
