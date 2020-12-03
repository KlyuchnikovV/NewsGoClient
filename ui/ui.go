package ui

import "fmt"

const (
	clearEscape  = "\033[H\033[2J"
	infoColor    = "\033[1;34m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
)

func ClearScreen() {
	print(clearEscape)
}

func PrintfNotice(format string, args ...interface{}) {
	print(format, args...)
}

func PrintfInfo(format string, args ...interface{}) {
	print(infoColor, fmt.Sprintf(format, args...))
}
func PrintfWarning(format string, args ...interface{}) {
	print(warningColor, fmt.Sprintf(format, args...))
}

func PrintfError(format string, args ...interface{}) {
	print(errorColor, fmt.Sprintf(format, args...))
}

func print(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
