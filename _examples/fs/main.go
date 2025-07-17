package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/afeiship/cobrautils/fs"
)

func main() {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "fs_example")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	fmt.Printf("Working in temporary directory: %s\n", tmpDir)

	// Example: EnsureDir
	dirPath := filepath.Join(tmpDir, "my_new_dir")
	fmt.Printf("Ensuring directory: %s\n", dirPath)
	if err := fs.EnsureDir(dirPath); err != nil {
		log.Fatalf("Failed to ensure directory: %v", err)
	}
	fmt.Printf("Directory exists: %t\n", fs.Exists(dirPath))

	// Example: EnsureFile
	filePath := filepath.Join(dirPath, "my_file.txt")
	fmt.Printf("Ensuring file: %s\n", filePath)
	if err := fs.EnsureFile(filePath); err != nil {
		log.Fatalf("Failed to ensure file: %v", err)
	}
	fmt.Printf("File exists: %t\n", fs.Exists(filePath))

	// Example: Copy
	copiedFilePath := filepath.Join(tmpDir, "my_copied_file.txt")
	fmt.Printf("Copying %s to %s\n", filePath, copiedFilePath)
	if err := fs.Copy(filePath, copiedFilePath); err != nil {
		log.Fatalf("Failed to copy file: %v", err)
	}
	fmt.Printf("Copied file exists: %t\n", fs.Exists(copiedFilePath))

	// Example: Remove
	fmt.Printf("Removing file: %s\n", copiedFilePath)
	if err := fs.Remove(copiedFilePath); err != nil {
		log.Fatalf("Failed to remove file: %v", err)
	}
	fmt.Printf("Copied file exists after removal: %t\n", fs.Exists(copiedFilePath))

	// Example: RemoveAll
	fmt.Printf("Removing directory and its contents: %s\n", dirPath)
	if err := fs.RemoveAll(dirPath); err != nil {
		log.Fatalf("Failed to remove all: %v", err)
	}
	fmt.Printf("Directory exists after removal: %t\n", fs.Exists(dirPath))
}
