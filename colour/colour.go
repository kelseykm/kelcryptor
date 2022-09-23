package colour

import "fmt"

// Exported colours
const (
	Normal        = "\033[0m"
	Bold          = "\033[1m"
	Italic        = "\033[3m"
	Underline     = "\033[4m"
	Blink         = "\033[5m"
	Invert        = "\033[7m"
	Invisible     = "\033[8m"
	StrikeThrough = "\033[9m"
	Overwrite     = "\033[2K" + "\r"

	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	BlackBackground   = "\033[40m"
	RedBackground     = "\033[41m"
	GreenBackground   = "\033[42m"
	YellowBackground  = "\033[43m"
	BlueBackground    = "\033[44m"
	MagentaBackground = "\033[45m"
	CyanBackground    = "\033[46m"
	WhiteBackground   = "\033[47m"
)

// Info returns a formatted info banner
func Info() string {
	return fmt.Sprintf(
		"%s%s[INFO]%s",
		BlueBackground,
		Bold,
		Normal,
	)
}

// Error returns a formatted error banner
func Error() string {
	return fmt.Sprintf(
		"%s%s[ERROR]%s",
		RedBackground,
		Bold,
		Normal,
	)
}

// Input returns a formatted input banner
func Input() string {
	return fmt.Sprintf(
		"%s%s[INPUT]%s",
		GreenBackground,
		Bold,
		Normal,
	)
}

// Message formats the mesg string
func Message(mesg string) string {
	return fmt.Sprintf(
		"%s%s%s",
		Bold,
		mesg,
		Normal,
	)
}

// FileName formats the file string
func FileName(file string) string {
	return fmt.Sprintf(
		"%s%s%s",
		Underline,
		file,
		Normal,
	)
}

// Time returns the time in a formatted string
func Time(time float64) string {
	return fmt.Sprintf(
		"%s%s%.2f%s",
		YellowBackground,
		Bold,
		time,
		Normal,
	)
}
