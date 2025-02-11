package formats

type Format int

func (f Format) String() string {
	if name, ok := FormatNames[f]; ok {
		return name
	}
	return "unknown"
}

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

var FormatNames = map[Format]string{
	JSON:    "JSON",
	Text:    "TEXT",
	Console: "CONSOLE",
}
