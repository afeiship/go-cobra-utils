package archive

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompressDecompressTarGz(t *testing.T) {
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "src")
	destArchive := filepath.Join(tmpDir, "archive.tar.gz")
	decompressDir := filepath.Join(tmpDir, "dest")

	// Create source directory and files
	err := os.MkdirAll(srcDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir(filepath.Join(srcDir, "subdir"), 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(srcDir, "subdir", "file2.txt"), []byte("world"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Compress
	err = CompressTarGz(srcDir, destArchive)
	if err != nil {
		t.Fatal(err)
	}

	// Decompress
	err = DecompressTarGz(destArchive, decompressDir)
	if err != nil {
		t.Fatal(err)
	}

	// Verify decompressed files
	_, err = os.Stat(filepath.Join(decompressDir, "file1.txt"))
	if os.IsNotExist(err) {
		t.Fatal("file1.txt not found after decompression")
	}

	_, err = os.Stat(filepath.Join(decompressDir, "subdir", "file2.txt"))
	if os.IsNotExist(err) {
		t.Fatal("file2.txt not found after decompression")
	}
}

func TestCompressDecompressTarXz(t *testing.T) {
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "src_xz")
	destArchive := filepath.Join(tmpDir, "archive.tar.xz")
	decompressDir := filepath.Join(tmpDir, "dest_xz")

	// Create source directory and files
	err := os.MkdirAll(srcDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(srcDir, "file1.txt"), []byte("hello xz"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Compress
	err = CompressTarXz(srcDir, destArchive)
	if err != nil {
		t.Fatal(err)
	}

	// Decompress
	err = DecompressTarXz(destArchive, decompressDir)
	if err != nil {
		t.Fatal(err)
	}

	// Verify decompressed files
	content, err := os.ReadFile(filepath.Join(decompressDir, "file1.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "hello xz" {
		t.Fatalf("Expected \"hello xz\", got \"%s\"", string(content))
	}
}
