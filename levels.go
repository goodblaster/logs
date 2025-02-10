package logs

type Level int
type LevelName string

const (
	LevelDebug Level = iota - 1
	LevelInfo
	LevelWarn
	LevelError
	LevelDPanic
	LevelPanic
	LevelFatal
	LevelPrint = 10
)

// Levels - change if you like.
var Levels = []Level{
	LevelDebug,
	LevelInfo,
	LevelWarn,
	LevelError,
	LevelDPanic,
	LevelPanic,
	LevelFatal,
	LevelPrint,
}

// LevelNames - change if you like.
var LevelNames = map[Level]LevelName{
	LevelDebug:  "debug",
	LevelInfo:   "info",
	LevelWarn:   "warn",
	LevelError:  "error",
	LevelDPanic: "dpanic",
	LevelPanic:  "panic",
	LevelFatal:  "fatal",
	LevelPrint:  "print",
}

func (l Level) String() string {
	if name, ok := LevelNames[l]; ok {
		return string(name)
	}
	return "unknown"
}
