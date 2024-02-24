package colour

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	InfoColor     = color.New(color.FgBlue, color.Bold).SprintFunc()
	ErrorColor    = color.New(color.FgRed, color.Bold).SprintFunc()
	InputColor    = color.New(color.FgGreen, color.Bold).SprintFunc()
	MessageColor  = color.New(color.Bold).SprintFunc()
	FileNameColor = color.New(color.Underline).SprintFunc()
	TimeColor     = color.New(color.FgYellow, color.Bold).SprintFunc()
	Info          = InfoColor("[INFO]")
	Error         = ErrorColor("[ERROR]")
	Input         = InputColor("[INPUT]")
)

// Message formats the mesg string
func Message(mesg string) string {
	return MessageColor(mesg)
}

// FileName formats the file string
func FileName(file string) string {
	return FileNameColor(file)
}

// Time returns the time in a formatted string
func Time(time float64) string {
	return TimeColor(fmt.Sprintf("%.2f", time))
}
