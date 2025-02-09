package formats

type Format int
type FormatName string

const (
	JSON Format = iota
	Text
	Console
)

var Formats = []Format{
	JSON,
	Text,
	Console,
}

var FormatNames = map[Format]FormatName{
	JSON:    "JSON",
	Text:    "TEXT",
	Console: "CONSOLE",
}
