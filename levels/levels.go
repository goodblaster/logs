package levels

import "github.com/goodblaster/logs/colors"

type Level int

func (level Level) String() string {
	if name, ok := LevelNames[level]; ok {
		return name
	}
	return "unknown"
}

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

// LevelNames - change if you like.
var LevelNames map[Level]string

var LevelColors map[Level]colors.TextColor

// init - set defaults
func init() {
	LevelNames = map[Level]string{
		Debug:  "debug",
		Info:   "info",
		Warn:   "warn",
		Error:  "error",
		DPanic: "dpanic",
		Panic:  "panic",
		Fatal:  "fatal",
		Print:  "print",
	}

	LevelColors = map[Level]colors.TextColor{
		Debug:  colors.TextColorBlue,
		Info:   colors.TextColorGreen,
		Warn:   colors.TextColorYellow,
		Error:  colors.TextColorRed,
		DPanic: colors.TextColorPurple,
		Panic:  colors.TextColorPurple,
		Fatal:  colors.TextColorPurple,
		Print:  colors.TextColorReset,
	}
}
