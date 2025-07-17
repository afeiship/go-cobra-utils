package logger

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() { log.SetOutput(os.Stderr) }()

	Info("test info: %s", "message")

	expected := "INFO: test info: message\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() { log.SetOutput(os.Stderr) }()

	Warn("test warn: %s", "message")

	expected := "WARN: test warn: message\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() { log.SetOutput(os.Stderr) }()

	Error("test error: %s", "message")

	expected := "ERROR: test error: message\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}

func TestSuccess(t *testing.T) {
	var buf bytes.Buffer
	// Redirect stdout to capture fmt.Printf output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Success("test success: %s", "message")

	w.Close()
	os.Stdout = oldStdout // Restore stdout

	buf.ReadFrom(r)

	expected := "✅ test success: message\n"
	if buf.String() != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}

func TestFail(t *testing.T) {
	var buf bytes.Buffer
	// Redirect stdout to capture fmt.Printf output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Fail("test fail: %s", "message")

	w.Close()
	os.Stdout = oldStdout // Restore stdout

	buf.ReadFrom(r)

	expected := "❌ test fail: message\n"
	if buf.String() != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}

func TestSetSuccessEmoji(t *testing.T) {
	defer SetSuccessEmoji("✅") // Reset to default after test

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	SetSuccessEmoji("👍")
	Success("custom success: %s", "message")

	w.Close()
	os.Stdout = oldStdout

	buf.ReadFrom(r)

	expected := "👍 custom success: message\n"
	if buf.String() != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}

func TestSetFailEmoji(t *testing.T) {
	defer SetFailEmoji("❌") // Reset to default after test

	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	SetFailEmoji("👎")
	Fail("custom fail: %s", "message")

	w.Close()
	os.Stdout = oldStdout

	buf.ReadFrom(r)

	expected := "👎 custom fail: message\n"
	if buf.String() != expected {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, buf.String())
	}
}