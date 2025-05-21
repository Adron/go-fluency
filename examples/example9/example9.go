package example9

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Run demonstrates Go's tooling ecosystem
func Run() error {
	fmt.Println("--- Go Tooling Ecosystem Example ---")

	// 1. Advanced build options
	fmt.Println("Building with advanced options...")
	cmd := exec.Command("go", "build", "-ldflags=-s -w", "-tags=prod", "-o", "example9_binary", ".")
	cmd.Dir = "examples/example9"
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}
	fmt.Println("Build completed successfully.")

	// 2. Code generation (simulated)
	fmt.Println("Simulating code generation...")
	// In a real scenario, you might run 'go generate' here
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Code generation simulated.")

	// 3. Static analysis (simulated)
	fmt.Println("Running static analysis...")
	// In a real scenario, you might run 'go vet' or 'golangci-lint' here
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Static analysis completed.")

	// Clean up
	os.Remove("examples/example9/example9_binary")
	return nil
}

// RunBenchmark demonstrates benchmarking tooling operations
func RunBenchmark() error {
	fmt.Println("--- Go Tooling Benchmark ---")
	start := time.Now()
	iters := 5
	for i := 0; i < iters; i++ {
		cmd := exec.Command("go", "build", "-o", "example9_binary", ".")
		cmd.Dir = "examples/example9"
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("benchmark build failed: %w", err)
		}
		os.Remove("examples/example9/example9_binary")
	}
	duration := time.Since(start)
	fmt.Printf("Ran build %d times in %v\n", iters, duration)
	fmt.Printf("Average build time: %v\n", duration/time.Duration(iters))
	return nil
}
