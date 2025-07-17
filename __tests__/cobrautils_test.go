package cobrautils_test

import (
	"github.com/afeiship/cobrautils"
	"testing"
)

func TestExec(t *testing.T) {
	output, err := cobrautils.Exec("echo 'hello world'")
	if err != nil {
		t.Fatal(err)
	}

	if output != "hello world\n" {
		t.Fatalf("Expected 'hello world', got '%s'", output)
	}
}
