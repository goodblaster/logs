package levels

type Level int
type LevelName string

const (
	Debug Level = iota - 1
	Info
	Warn
	Error
	DPanic
	Panic
	Fatal
	Print = 10
)

var Levels = []Level{
	Debug,
	Info,
	Warn,
	Error,
	DPanic,
	Panic,
	Fatal,
	Print,
}

var LevelNames = map[Level]LevelName{
	Debug:  "DEBUG",
	Info:   "INFO",
	Warn:   "WARN",
	Error:  "ERROR",
	DPanic: "DPANIC",
	Panic:  "PANIC",
	Fatal:  "FATAL",
	Print:  "PRINT",
}
