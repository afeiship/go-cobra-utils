package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/afeiship/cobrautils/archive"
	"github.com/afeiship/cobrautils/fs"
)

func main() {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "archive_example")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir) // Clean up

	fmt.Printf("Working in temporary directory: %s\n", tmpDir)

	// --- Tar.gz Example ---
	fmt.Println("\n--- Tar.gz Example ---")
	srcDirGz := filepath.Join(tmpDir, "src_gz")
	destArchiveGz := filepath.Join(tmpDir, "archive.tar.gz")
	decompressDirGz := filepath.Join(tmpDir, "dest_gz")

	// Create source directory and files
	if err := os.MkdirAll(srcDirGz, 0755); err != nil {
		log.Fatalf("Failed to create srcDirGz: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDirGz, "file1.txt"), []byte("hello tar.gz"), 0644); err != nil {
		log.Fatalf("Failed to write file1.txt: %v", err)
	}

	fmt.Printf("Compressing %s to %s\n", srcDirGz, destArchiveGz)
	if err := archive.CompressTarGz(srcDirGz, destArchiveGz); err != nil {
		log.Fatalf("Failed to compress tar.gz: %v", err)
	}
	fmt.Printf("Decompressing %s to %s\n", destArchiveGz, decompressDirGz)
	if err := archive.DecompressTarGz(destArchiveGz, decompressDirGz); err != nil {
		log.Fatalf("Failed to decompress tar.gz: %v", err)
	}

	// Verify
	if !fs.Exists(filepath.Join(decompressDirGz, "file1.txt")) {
		log.Fatalf("file1.txt not found after tar.gz decompression")
	}
	fmt.Println("Tar.gz example completed successfully.")

	// --- Tar.xz Example ---
	fmt.Println("\n--- Tar.xz Example ---")
	srcDirXz := filepath.Join(tmpDir, "src_xz")
	destArchiveXz := filepath.Join(tmpDir, "archive.tar.xz")
	decompressDirXz := filepath.Join(tmpDir, "dest_xz")

	// Create source directory and files
	if err := os.MkdirAll(srcDirXz, 0755); err != nil {
		log.Fatalf("Failed to create srcDirXz: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDirXz, "file1.txt"), []byte("hello tar.xz"), 0644); err != nil {
		log.Fatalf("Failed to write file1.txt: %v", err)
	}

	fmt.Printf("Compressing %s to %s\n", srcDirXz, destArchiveXz)
	if err := archive.CompressTarXz(srcDirXz, destArchiveXz); err != nil {
		log.Fatalf("Failed to compress tar.xz: %v", err)
	}
	fmt.Printf("Decompressing %s to %s\n", destArchiveXz, decompressDirXz)
	if err := archive.DecompressTarXz(destArchiveXz, decompressDirXz); err != nil {
		log.Fatalf("Failed to decompress tar.xz: %v", err)
	}

	// Verify
	if !fs.Exists(filepath.Join(decompressDirXz, "file1.txt")) {
		log.Fatalf("file1.txt not found after tar.xz decompression")
	}
	fmt.Println("Tar.xz example completed successfully.")
}