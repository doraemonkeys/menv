package color

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Bold variants
	BoldRed     = "\033[1;31m"
	BoldGreen   = "\033[1;32m"
	BoldYellow  = "\033[1;33m"
	BoldBlue    = "\033[1;34m"
	BoldMagenta = "\033[1;35m"
	BoldCyan    = "\033[1;36m"
	BoldWhite   = "\033[1;37m"
)

func init() {
	enableVirtualTerminal()
}

// enableVirtualTerminal enables ANSI escape sequences on Windows
func enableVirtualTerminal() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(stdout, &mode); err == nil {
		_ = windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
}

// Print functions with colors

// func printColored(color, format string, args ...any) {
// 	fmt.Printf("%s%s%s", color, fmt.Sprintf(format, args...), Reset)
// }

func printColoredLn(color, format string, args ...any) {
	fmt.Printf("%s%s%s\n", color, fmt.Sprintf(format, args...), Reset)
}

// Success prints a green success message
func Success(format string, args ...any) {
	printColoredLn(Green, format, args...)
}

// Error prints a red error message
func Error(format string, args ...any) {
	printColoredLn(Red, format, args...)
}

// Warning prints a yellow warning message
func Warning(format string, args ...any) {
	printColoredLn(Yellow, format, args...)
}

// Info prints a cyan info message
func Info(format string, args ...any) {
	printColoredLn(Cyan, format, args...)
}

// Highlight prints a magenta highlighted message
func Highlight(format string, args ...any) {
	printColoredLn(Magenta, format, args...)
}

// Sprintf returns a colored string
func Sprintf(c, format string, args ...any) string {
	return fmt.Sprintf("%s%s%s", c, fmt.Sprintf(format, args...), Reset)
}
