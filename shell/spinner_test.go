package shell

import (
	"testing"
	"time"
)

func TestRunWithSpinner(t *testing.T) {
	// Test with default config (nil)
	err := RunWithSpinner("echo 'hello world'", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Test with custom config
	customConfig := &SpinnerConfig{
		CharSetIndex: 14, // A different character set
		Speed:        200 * time.Millisecond,
	}
	err = RunWithSpinner("echo 'hello world'", customConfig)
	if err != nil {
		t.Fatal(err)
	}
}
