package colour

import "fmt"

// Exported colours
const (
	Normal    = "\033[0;39m"
	Invisible = "\033[8m"
	Overwrite = "\033[2K\r"

	Red   = "\033[31m"
	Green = "\033[32m"
	Brown = "\033[33m"
	Blue  = "\033[34m"
	White = "\033[39m"

	RedBold   = "\033[1;31m"
	GreenBold = "\033[1;32m"
	BrownBold = "\033[1;33m"
	BlueBold  = "\033[1;34m"
	WhiteBold = "\033[1;39m"

	RedUnderlined   = "\033[4;31m"
	GreenUnderlined = "\033[4;32m"
	BrownUnderlined = "\033[4;33m"
	BlueUnderlined  = "\033[4;34m"
	WhiteUnderlined = "\033[4;39m"

	RedItalicised   = "\033[3;31m"
	GreenItalicised = "\033[3;32m"
	BrownItalicised = "\033[3;33m"
	BlueItalicised  = "\033[3;34m"
	WhiteItalicised = "\033[3;39m"

	RedStrikeThrough   = "\033[9;31m"
	GreenStrikeThrough = "\033[9;32m"
	BrownStrikeThrough = "\033[9;33m"
	BlueStrikeThrough  = "\033[9;34m"
	WhiteStrikeThrough = "\033[9;39m"

	RedBackground   = "\033[7;31m"
	GreenBackground = "\033[7;32m"
	BrownBackground = "\033[7;33m"
	BlueBackground  = "\033[7;34m"
	WhiteBackground = "\033[7;39m"

	RedBlinking   = "\033[5;31m"
	GreenBlinking = "\033[5;32m"
	BrownBlinking = "\033[5;33m"
	BlueBlinking  = "\033[5;34m"
	WhiteBlinking = "\033[5;39m"
)

func Info() string {
	return fmt.Sprintf(
		"%s%s[INFO]%s",
		BlueBackground,
		BlueBold,
		Normal,
	)
}

func Error() string {
	return fmt.Sprintf(
		"%s%s[ERROR]%s",
		RedBackground,
		RedBold,
		Normal,
	)
}

func Input() string {
	return fmt.Sprintf(
		"%s%s[INPUT]%s",
		GreenBackground,
		GreenBold,
		Normal,
	)
}

func Message(mesg string) string {
	return fmt.Sprintf(
		"%s%s%s",
		WhiteBold,
		Normal,
	)
}

func FileName(file string) string {
	return fmt.Sprintf(
		"%s%s%s",
		WhiteUnderlined,
		Normal,
	)
}
