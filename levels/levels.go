package levels

import (
	"math"

	"github.com/goodblaster/logs/colors"
)

type Level int

func (level Level) String() string {
	if name, ok := LevelNames[level]; ok {
		return name
	}

	// Keep "print" as a default option.
	if level == Print {
		return "print"
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
	Print = math.MaxInt
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
		Debug:  colors.TextBlue,
		Info:   colors.TextGreen,
		Warn:   colors.TextYellow,
		Error:  colors.TextRed,
		DPanic: colors.TextPurple,
		Panic:  colors.TextPurple,
		Fatal:  colors.TextPurple,
		Print:  colors.Reset,
	}
}
