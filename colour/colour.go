package colour

import "fmt"

const esc = 27 // \033

// Exported colours
var (
	Normal        = []byte{esc, '[', '0', 'm'}
	Bold          = []byte{esc, '[', '1', 'm'}
	Italic        = []byte{esc, '[', '3', 'm'}
	Underline     = []byte{esc, '[', '4', 'm'}
	Blink         = []byte{esc, '[', '5', 'm'}
	Invert        = []byte{esc, '[', '7', 'm'}
	Invisible     = []byte{esc, '[', '8', 'm'}
	StrikeThrough = []byte{esc, '[', '9', 'm'}
	Overwrite     = []byte{esc, '[', '2', 'K', '\r'}

	Black   = []byte{esc, '[', '3', '0', 'm'}
	Red     = []byte{esc, '[', '3', '1', 'm'}
	Green   = []byte{esc, '[', '3', '2', 'm'}
	Yellow  = []byte{esc, '[', '3', '3', 'm'}
	Blue    = []byte{esc, '[', '3', '4', 'm'}
	Magenta = []byte{esc, '[', '3', '5', 'm'}
	Cyan    = []byte{esc, '[', '3', '6', 'm'}
	White   = []byte{esc, '[', '3', '7', 'm'}

	BlackBackground   = []byte{esc, '[', '4', '0', 'm'}
	RedBackground     = []byte{esc, '[', '4', '1', 'm'}
	GreenBackground   = []byte{esc, '[', '4', '2', 'm'}
	YellowBackground  = []byte{esc, '[', '4', '3', 'm'}
	BlueBackground    = []byte{esc, '[', '4', '4', 'm'}
	MagentaBackground = []byte{esc, '[', '4', '5', 'm'}
	CyanBackground    = []byte{esc, '[', '4', '6', 'm'}
	WhiteBackground   = []byte{esc, '[', '4', '7', 'm'}
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
