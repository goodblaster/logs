package logs

type Format int
type FormatName string

const (
	FormatJSON Format = iota
	FormatText
	FormatConsole
)

var Formats = []Format{
	FormatJSON,
	FormatText,
	FormatConsole,
}

var FormatNames = map[Format]FormatName{
	FormatJSON:    "FormatJSON",
	FormatText:    "TEXT",
	FormatConsole: "CONSOLE",
}
