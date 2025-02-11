package contrib

import (
	"io"

	"github.com/goodblaster/logs"
	"github.com/goodblaster/logs/formats"
	"github.com/goodblaster/logs/levels"
	"github.com/goodblaster/logs/pkg/adapters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(level levels.Level, format formats.Format, writer io.Writer) logs.Interface {
	encoding := func(format formats.Format) string {
		switch format {
		case formats.JSON:
			return "json"
		case formats.Console, formats.Text:
			return "console"
		}
		return "json"
	}

	zapLevel := zap.DebugLevel
	switch level {
	case levels.Debug:
		zapLevel = zap.DebugLevel
	case levels.Info:
		zapLevel = zap.InfoLevel
	case levels.Warn:
		zapLevel = zap.WarnLevel
	case levels.Error:
		zapLevel = zap.ErrorLevel
	case levels.Fatal:
		zapLevel = zap.FatalLevel
	case levels.Panic:
		zapLevel = zap.PanicLevel
	}

	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapLevel),
		Development: level == levels.Debug,
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
	case formats.JSON:
		encoder = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	case formats.Console, formats.Text:
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
