package main

import (
	"fmt"
	"log"
	"time"

	"github.com/afeiship/cobrautils/shell"
)

func main() {
	// Example: Exec
	fmt.Println("--- Exec Example ---")
	output, err := shell.Exec("echo Hello from Exec")
	if err != nil {
		log.Fatalf("Exec failed: %v", err)
	}
	fmt.Printf("Exec Output: %s", output)

	// Example: Run
	fmt.Println("\n--- Run Example ---")
	fmt.Println("Running 'sleep 1'...")
	if err := shell.Run("sleep 1"); err != nil {
		log.Fatalf("Run failed: %v", err)
	}
	fmt.Println("Run finished.")

	// Example: RunWithSpinner (default config)
	fmt.Println("\n--- RunWithSpinner (Default) Example ---")
	fmt.Println("Running 'sleep 2' with default spinner...")
	if err := shell.RunWithSpinner("sleep 2", nil); err != nil {
		log.Fatalf("RunWithSpinner (default) failed: %v", err)
	}
	fmt.Println("RunWithSpinner (default) finished.")

	// Example: RunWithSpinner (custom config)
	fmt.Println("\n--- RunWithSpinner (Custom) Example ---")
	fmt.Println("Running 'sleep 2' with custom spinner...")
	customConfig := &shell.SpinnerConfig{
		CharSetIndex: 70, // A different character set
		Speed:        200 * time.Millisecond,
	}
	if err := shell.RunWithSpinner("sleep 2", customConfig); err != nil {
		log.Fatalf("RunWithSpinner (custom) failed: %v", err)
	}
	fmt.Println("RunWithSpinner (custom) finished.")
}
