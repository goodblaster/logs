package colors

type TextColor string

const (
	TextColorReset   TextColor = "\033[0m"
	TextColorRed     TextColor = "\033[31m"
	TextColorYellow  TextColor = "\033[33m"
	TextColorGreen   TextColor = "\033[32m"
	TextColorBlue    TextColor = "\033[34m"
	TextColorMagenta TextColor = "\033[35m"
	TextColorCyan    TextColor = "\033[36m"
	TextColorWhite   TextColor = "\033[37m"
	TextColorPurple  TextColor = "\033[35m"
)
