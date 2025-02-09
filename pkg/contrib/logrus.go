package contrib

import (
	"io"
	"os"

	"github.com/goodblaster/logs/pkg/adapters"
	"github.com/goodblaster/logs/pkg/formats"
	"github.com/goodblaster/logs/pkg/levels"
	"github.com/sirupsen/logrus"
)

func NewLogrusLogger(level levels.Level, format formats.Format, writer io.Writer) *adapters.LogrusAdapter {
	var formatter logrus.Formatter
	if format == formats.JSON {
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
