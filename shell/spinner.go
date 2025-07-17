package shell

import (
	"time"

	"github.com/briandowns/spinner"
)

// SpinnerConfig holds the configuration for the spinner.
type SpinnerConfig struct {
	CharSetIndex int
	Speed        time.Duration
}

// RunWithSpinner executes a shell command with a spinner.
// It accepts an optional SpinnerConfig. If nil, default values are used.
func RunWithSpinner(cmd string, config *SpinnerConfig) error {
	cs := spinner.CharSets[9]
	speed := 100 * time.Millisecond

	if config != nil {
		if config.CharSetIndex >= 0 && config.CharSetIndex < len(spinner.CharSets) {
			cs = spinner.CharSets[config.CharSetIndex]
		}
		if config.Speed > 0 {
			speed = config.Speed
		}
	}

	s := spinner.New(cs, speed)
	s.Start()
	defer s.Stop()

	return Run(cmd)
}
