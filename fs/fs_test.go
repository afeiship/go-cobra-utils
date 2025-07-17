package fs

import (
	"os"
	"testing"
)

func TestExists(t *testing.T) {
	// Create a dummy file
	f, err := os.Create("dummy.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if !Exists("dummy.txt") {
		t.Fatal("Expected file to exist")
	}

	// Clean up
	os.Remove("dummy.txt")
}

func TestIsFile(t *testing.T) {
	// Create a dummy file
	f, err := os.Create("dummy.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if !IsFile("dummy.txt") {
		t.Fatal("Expected path to be a file")
	}

	// Clean up
	os.Remove("dummy.txt")
}

func TestIsDir(t *testing.T) {
	// Create a dummy directory
	err := os.Mkdir("dummy_dir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	if !IsDir("dummy_dir") {
		t.Fatal("Expected path to be a directory")
	}

	// Clean up
	os.Remove("dummy_dir")
}

func TestCopy(t *testing.T) {
	// Create a dummy file
	f, err := os.Create("dummy.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("hello world")
	f.Close()

	err = Copy("dummy.txt", "dummy_copy.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Check if the copied file exists
	if !Exists("dummy_copy.txt") {
		t.Fatal("Expected copied file to exist")
	}

	// Check if the content is the same
	content, err := os.ReadFile("dummy_copy.txt")
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "hello world" {
		t.Fatalf("Expected 'hello world', got '%s'", string(content))
	}

	// Clean up
	os.Remove("dummy.txt")
	os.Remove("dummy_copy.txt")
}

func TestEnsureDir(t *testing.T) {
	// Ensure a directory that does not exist
	err := EnsureDir("dummy_dir")
	if err != nil {
		t.Fatal(err)
	}

	if !IsDir("dummy_dir") {
		t.Fatal("Expected path to be a directory")
	}

	// Clean up
	os.Remove("dummy_dir")
}

func TestEnsureFile(t *testing.T) {
	// Ensure a file that does not exist
	err := EnsureFile("dummy.txt")
	if err != nil {
		t.Fatal(err)
	}

	if !IsFile("dummy.txt") {
		t.Fatal("Expected path to be a file")
	}

	// Clean up
	os.Remove("dummy.txt")
}
