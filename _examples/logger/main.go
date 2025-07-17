package main

import (
	"github.com/afeiship/cobrautils/logger"
)

func main() {
	logger.Info("This is an info message: %s", "hello")
	logger.Warn("This is a warning message: %d", 123)
	logger.Error("This is an error message: %v", "something went wrong")

	logger.Success("Operation completed successfully!")
	logger.Fail("Operation failed!")

	// 使用默认表情符号logger.Success"默认成功消息")
	logger.Fail("默认失败消息")

	// 自定义表情符号
	logger.SetSuccessEmoji("👍")
	logger.SetFailEmoji("👎")
	logger.Success("自定义成功消息")
	logger.Fail("自定义失败消息")
}
