package shell

import (
	"os/exec"
)

// Exec executes a shell command and returns the combined output.
func Exec(cmd string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	return string(output), err
}

// Run executes a shell command without returning the output.
func Run(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	return c.Run()
}
