package fs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	tmpDir := t.TempDir()
	dummyFile := filepath.Join(tmpDir, "dummy.txt")

	// Create a dummy file
	f, err := os.Create(dummyFile)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if !Exists(dummyFile) {
		t.Fatal("Expected file to exist")
	}
}

func TestIsFile(t *testing.T) {
	tmpDir := t.TempDir()
	dummyFile := filepath.Join(tmpDir, "dummy.txt")

	// Create a dummy file
	f, err := os.Create(dummyFile)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if !IsFile(dummyFile) {
		t.Fatal("Expected path to be a file")
	}
}

func TestIsDir(t *testing.T) {
	tmpDir := t.TempDir()
	dummyDir := filepath.Join(tmpDir, "dummy_dir")

	// Create a dummy directory
	err := os.Mkdir(dummyDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	if !IsDir(dummyDir) {
		t.Fatal("Expected path to be a directory")
	}
}

func TestCopy(t *testing.T) {
	tmpDir := t.TempDir()
	dummyFile := filepath.Join(tmpDir, "dummy.txt")
	dummyCopy := filepath.Join(tmpDir, "dummy_copy.txt")

	// Create a dummy file
	f, err := os.Create(dummyFile)
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("hello world")
	f.Close()

	err = Copy(dummyFile, dummyCopy)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the copied file exists
	if !Exists(dummyCopy) {
		t.Fatal("Expected copied file to exist")
	}

	// Check if the content is the same
	content, err := os.ReadFile(dummyCopy)
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "hello world" {
		t.Fatalf("Expected 'hello world', got '%s'", string(content))
	}
}

func TestEnsureDir(t *testing.T) {
	tmpDir := t.TempDir()
	dummyDir := filepath.Join(tmpDir, "dummy_dir")

	// Ensure a directory that does not exist
	err := EnsureDir(dummyDir)
	if err != nil {
		t.Fatal(err)
	}

	if !IsDir(dummyDir) {
		t.Fatal("Expected path to be a directory")
	}
}

func TestEnsureFile(t *testing.T) {
	tmpDir := t.TempDir()
	dummyFile := filepath.Join(tmpDir, "dummy.txt")
	dummyDir := filepath.Join(tmpDir, "dummy_dir")

	// 1. Ensure a file that does not exist
	err := EnsureFile(dummyFile)
	if err != nil {
		t.Fatal(err)
	}

	if !IsFile(dummyFile) {
		t.Fatal("Expected path to be a file")
	}

	// 2. Ensure a file that already exists
	err = EnsureFile(dummyFile)
	if err != nil {
		t.Fatal(err)
	}

	// 3. Ensure a file when the path is a directory
	err = os.Mkdir(dummyDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = EnsureFile(dummyDir)
	if err == nil {
		t.Fatal("Expected an error when the path is a directory")
	}
}