package shell

import (
	"testing"
)

func TestExec(t *testing.T) {
	output, err := Exec("echo 'hello world'")
	if err != nil {
		t.Fatal(err)
	}

	if output != "hello world\n" {
		t.Fatalf("Expected 'hello world', got '%s'", output)
	}
}

func TestRun(t *testing.T) {
	err := Run("echo 'hello world'")
	if err != nil {
		t.Fatal(err)
	}
}
