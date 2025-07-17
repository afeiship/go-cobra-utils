package archive

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/afeiship/cobrautils/shell"
)

// CompressTarGz compresses a file or directory to a .tar.gz archive using shell commands.
func CompressTarGz(srcPath, destPath string) error {
	var cmd string
	fi, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		// If srcPath is a directory, archive its contents
		cmd = fmt.Sprintf("tar -czf %s -C %s .", destPath, srcPath)
	} else {
		// If srcPath is a file, archive the file itself
		srcDir := filepath.Dir(srcPath)
		srcName := filepath.Base(srcPath)
		cmd = fmt.Sprintf("tar -czf %s -C %s %s", destPath, srcDir, srcName)
	}
	return shell.Run(cmd)
}

// DecompressTarGz decompresses a .tar.gz archive to a directory using shell commands.
func DecompressTarGz(srcPath, destPath string) error {
	// Ensure destination directory exists
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return err
	}
	cmd := fmt.Sprintf("tar -xzf %s -C %s", srcPath, destPath)
	return shell.Run(cmd)
}

// CompressTarXz compresses a file or directory to a .tar.xz archive using shell commands.
func CompressTarXz(srcPath, destPath string) error {
	var cmd string
	fi, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		// If srcPath is a directory, archive its contents
		cmd = fmt.Sprintf("tar -cJf %s -C %s .", destPath, srcPath)
	} else {
		// If srcPath is a file, archive the file itself
		srcDir := filepath.Dir(srcPath)
		srcName := filepath.Base(srcPath)
		cmd = fmt.Sprintf("tar -cJf %s -C %s %s", destPath, srcDir, srcName)
	}
	return shell.Run(cmd)
}

// DecompressTarXz decompresses a .tar.xz archive to a directory using shell commands.
func DecompressTarXz(srcPath, destPath string) error {
	// Ensure destination directory exists
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return err
	}
	cmd := fmt.Sprintf("tar -xJf %s -C %s", srcPath, destPath)
	return shell.Run(cmd)
}