package example10

import (
	"errors"
	"fmt"
	"time"
)

// Custom error for demonstration
var ErrInvalidInput = errors.New("invalid input")

// User represents a user in the system
type User struct {
	ID   int
	Name string
}

// ProcessUser demonstrates idiomatic error handling and naming conventions
func ProcessUser(id int, name string) (*User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid user ID: %w", ErrInvalidInput)
	}
	if name == "" {
		return nil, fmt.Errorf("empty user name: %w", ErrInvalidInput)
	}
	return &User{ID: id, Name: name}, nil
}

// Run demonstrates idiomatic Go practices
func Run() error {
	fmt.Println("--- Idiomatic Go Example ---")

	// 1. Error handling
	user, err := ProcessUser(1, "Gopher")
	if err != nil {
		return fmt.Errorf("failed to process user: %w", err)
	}
	fmt.Printf("Processed user: %+v\n", user)

	// 2. Naming conventions and code organization
	// This section demonstrates proper naming and organization
	fmt.Println("Demonstrating naming conventions and code organization...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Naming conventions and code organization demonstrated.")

	return nil
}

// RunBenchmark demonstrates benchmarking idiomatic Go practices
func RunBenchmark() error {
	fmt.Println("--- Idiomatic Go Benchmark ---")
	start := time.Now()
	iters := 1000
	for i := 0; i < iters; i++ {
		_, err := ProcessUser(1, "Gopher")
		if err != nil {
			return fmt.Errorf("benchmark failed: %w", err)
		}
	}
	duration := time.Since(start)
	fmt.Printf("Ran ProcessUser %d times in %v\n", iters, duration)
	fmt.Printf("Average time per run: %v\n", duration/time.Duration(iters))
	return nil
}
