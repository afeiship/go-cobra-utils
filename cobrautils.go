package cobrautils

import (
	"os/exec"
)

// Exec executes a shell command.
func Exec(cmd string) (string, error) {
	c := exec.Command("sh", "-c", cmd)
	output, err := c.CombinedOutput()
	return string(output), err
}
