package colors

type TextColor string
type BgColor string

const (
	Reset TextColor = "\033[0m"

	TextRed     TextColor = "\033[31m"
	TextYellow  TextColor = "\033[33m"
	TextGreen   TextColor = "\033[32m"
	TextBlue    TextColor = "\033[34m"
	TextMagenta TextColor = "\033[35m"
	TextCyan    TextColor = "\033[36m"
	TextWhite   TextColor = "\033[37m"
	TextPurple  TextColor = "\033[35m"

	BgRed     BgColor = "\033[41m"
	BgYellow  BgColor = "\033[43m"
	BgGreen   BgColor = "\033[42m"
	BgBlue    BgColor = "\033[44m"
	BgMagenta BgColor = "\033[45m"
	BgCyan    BgColor = "\033[46m"
	BgWhite   BgColor = "\033[47m"
	BgBlack   BgColor = "\033[40m"
)
