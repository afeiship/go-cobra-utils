package logger

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	"time"
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

func TestSetTimestampFormat(t *testing.T) {
	// Save original log flags and output
	originalFlags := log.Flags()
	originalOutput := log.Writer()
	defer func() {
		log.SetFlags(originalFlags)
		log.SetOutput(originalOutput)
	}()

	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Test with only date
	SetTimestampFormat(log.Ldate)
	Info("test date format")
	output := buf.String()
	// Extract the timestamp part (e.g., "2006/01/02 ")
	timestampPart := strings.Split(output, " INFO:")[0]
	expectedDate := time.Now().Format("2006/01/02")
	if !strings.HasPrefix(timestampPart, expectedDate) || strings.Contains(timestampPart, ":") {
		t.Errorf("Expected date only, got: %s", timestampPart)
	}
	buf.Reset()

	// Test with only time
	SetTimestampFormat(log.Ltime)
	Info("test time format")
	output = buf.String()
	timestampPart = strings.Split(output, " INFO:")[0]
	expectedTime := time.Now().Format("15:04:05")
	if !strings.HasPrefix(timestampPart, expectedTime) || strings.Contains(timestampPart, "/") {
		t.Errorf("Expected time only, got: %s", timestampPart)
	}
	buf.Reset()

	// Test with no flags
	SetTimestampFormat(0)
	Info("test no format")
	output = buf.String()
	// In this case, there should be no timestamp part before " INFO:"
	if strings.HasPrefix(output, time.Now().Format("2006/01/02")) || strings.HasPrefix(output, time.Now().Format("15:04:05")) {
		t.Errorf("Expected no timestamp, got: %s", output)
	}
	// Also check that the output starts directly with the log level prefix
	if !strings.HasPrefix(output, "INFO:") {
		t.Errorf("Expected output to start with \"INFO:\", got: %s", output)
	}
}