package logger

import (
	"fmt"
	"log"
)

var (
	successEmoji = "✅"
	failEmoji    = "❌"
)

func init() {
	// Default log flags to include date and time
	log.SetFlags(log.LstdFlags)
}

// Info prints an informational log message.
func Info(format string, v ...interface{}) {
	log.Printf("INFO: "+format, v...)
}

// Warn prints a warning log message.
func Warn(format string, v ...interface{}) {
	log.Printf("WARN: "+format, v...)
}

// Error prints an error log message.
func Error(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}

// Success prints a success log message with an emoji.
func Success(format string, v ...interface{}) {
	fmt.Printf(successEmoji + " " + format + "\n", v...)
}

// Fail prints a failure log message with an emoji.
func Fail(format string, v ...interface{}) {
	fmt.Printf(failEmoji + " " + format + "\n", v...)
}

// SetSuccessEmoji sets the emoji for success messages.
func SetSuccessEmoji(emoji string) {
	successEmoji = emoji
}

// SetFailEmoji sets the emoji for failure messages.
func SetFailEmoji(emoji string) {
	failEmoji = emoji
}

// SetTimestampFormat sets the flags for the standard logger to control timestamp format.
// Use constants from the log package, e.g., log.LstdFlags, log.Ldate, log.Ltime.
func SetTimestampFormat(flags int) {
	log.SetFlags(flags)
}