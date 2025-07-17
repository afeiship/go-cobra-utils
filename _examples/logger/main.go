package main

import (
	"log"

	"github.com/afeiship/cobrautils/logger"
)

func main() {
	logger.Info("This is an info message: %s", "hello")
	logger.Warn("This is a warning message: %d", 123)
	logger.Error("This is an error message: %v", "something went wrong")

	logger.Success("Operation completed successfully!")
	logger.Fail("Operation failed!")

	// 使用默认表情符号
	logger.Success("默认成功消息")
	logger.Fail("默认失败消息")

	// 自定义表情符号
	logger.SetSuccessEmoji("👍")
	logger.SetFailEmoji("👎")
	logger.Success("自定义成功消息")
	logger.Fail("自定义失败消息")

	// --- Timestamp Format Examples ---
	logger.Info("\n--- Timestamp Format Examples ---")

	// Set to show only date
	logger.SetTimestampFormat(log.Ldate)
	logger.Info("Log with only date")

	// Set to show only time
	logger.SetTimestampFormat(log.Ltime)
	logger.Info("Log with only time")

	// Set to show date and time (standard)
	logger.SetTimestampFormat(log.LstdFlags)
	logger.Info("Log with standard date and time")

	// Set to show long form of time (microseconds)
	logger.SetTimestampFormat(log.Ltime | log.Lmicroseconds)
	logger.Info("Log with microseconds")

	// Set to show short file name and line number
	logger.SetTimestampFormat(log.Lshortfile)
	logger.Info("Log with short file and line number")

	// Set to no timestamp
	logger.SetTimestampFormat(0)
	logger.Info("Log with no timestamp")

	// Reset to default (LstdFlags) for other logs if any
	logger.SetTimestampFormat(log.LstdFlags)
}