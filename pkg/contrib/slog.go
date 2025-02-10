package contrib

import (
	"io"
	"log/slog"
	"strings"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/pkg/adapters"
)

func NewSLogLogger(level logs.Level, format logs.Format, writer io.Writer) *adapters.SLogAdapter {
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

	if format == logs.FormatJSON {
		handler = slog.NewJSONHandler(writer, options)
	} else {
		handler = slog.NewTextHandler(writer, options)
	}

	handler = &adapters.CustomSLogHandler{Handler: handler}
	return adapters.Slog(slog.New(handler))
}
