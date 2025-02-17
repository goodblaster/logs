package contrib

import (
	"io"
	"os"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
	"github.com/goodblaster/logs/pkg/adapters"
	"github.com/sirupsen/logrus"
)

func NewLogrusLogger(level levels.Level, format formats.Format, writer io.Writer) logs.Interface {
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
