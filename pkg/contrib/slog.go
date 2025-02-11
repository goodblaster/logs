package contrib

import (
	"io"
	"log/slog"
	"strings"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
	"github.com/goodblaster/logs/pkg/adapters"
)

func NewSLogLogger(level levels.Level, format formats.Format, writer io.Writer) logs.Interface {
	var handler slog.Handler

	options := &slog.HandlerOptions{
		Level: adapters.ToSLogLevel(level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove time from the output for predictable test output.
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			// Customize the name of the level key and the output string, including
			// custom level values.
			if a.Key == slog.LevelKey {
				a.Value = slog.StringValue(strings.ToLower(a.Value.String()))
			}

			return a
		},
	}

	if format == formats.JSON {
		handler = slog.NewJSONHandler(writer, options)
	} else {
		handler = slog.NewTextHandler(writer, options)
	}

	handler = &adapters.CustomSLogHandler{Handler: handler}
	return adapters.Slog(slog.New(handler))
}
