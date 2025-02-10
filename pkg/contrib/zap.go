package contrib

import (
	"io"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/pkg/adapters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(level logs.Level, format logs.Format, writer io.Writer) *adapters.ZapAdapter {
	encoding := func(format logs.Format) string {
		switch format {
		case logs.FormatJSON:
			return "json"
		case logs.FormatConsole, logs.FormatText:
			return "console"
		}
		return "json"
	}

	zapLevel := zap.DebugLevel
	switch level {
	case logs.LevelDebug:
		zapLevel = zap.DebugLevel
	case logs.LevelInfo:
		zapLevel = zap.InfoLevel
	case logs.LevelWarn:
		zapLevel = zap.WarnLevel
	case logs.LevelError:
		zapLevel = zap.ErrorLevel
	case logs.LevelFatal:
		zapLevel = zap.FatalLevel
	case logs.LevelPanic:
		zapLevel = zap.PanicLevel
	}

	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapLevel),
		Development: level == logs.LevelDebug,
		Encoding:    encoding(format),
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    adapters.CustomLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zapConfig.EncoderConfig.EncodeLevel = adapters.CustomLevelEncoder

	var encoder zapcore.Encoder
	var enabler zapcore.LevelEnabler

	switch format {
	case logs.FormatJSON:
		encoder = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	case logs.FormatConsole, logs.FormatText:
		encoder = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	}

	enabler = zapLevel

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(writer),
		enabler,
	)

	logger := zap.New(core,
		zap.WithCaller(false),
		zap.AddStacktrace(adapters.PrintLevel+1),
	)

	return adapters.Zap(logger)
}
