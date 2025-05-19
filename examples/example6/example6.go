package example6

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Common errors
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
)

// DataProcessor represents a data processing service
type DataProcessor interface {
	Process(ctx context.Context, data []string) ([]string, error)
	ProcessBatch(ctx context.Context, data []string, batchSize int) ([]string, error)
	ProcessConcurrent(ctx context.Context, data []string, workers int) ([]string, error)
}

// processor implements the DataProcessor interface
type processor struct {
	// Object pool for reuse
	pool sync.Pool
}

// NewProcessor creates a new data processor
func NewProcessor() DataProcessor {
	return &processor{
		pool: sync.Pool{
			New: func() interface{} {
				// Pre-allocate buffer for string processing
				return make([]byte, 0, 1024)
			},
		},
	}
}

// processItem processes a single item with memory optimization
func (p *processor) processItem(item string) string {
	// Get buffer from pool
	buf := p.pool.Get().([]byte)
	defer p.pool.Put(buf)

	// Reset buffer but keep capacity
	buf = buf[:0]

	// Process the item (example: convert to uppercase)
	for i := 0; i < len(item); i++ {
		if item[i] >= 'a' && item[i] <= 'z' {
			buf = append(buf, item[i]-32)
		} else {
			buf = append(buf, item[i])
		}
	}

	return string(buf)
}

// Process processes data with optimized memory allocation
func (p *processor) Process(ctx context.Context, data []string) ([]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%w: empty data", ErrInvalidInput)
	}

	// Pre-allocate result slice with exact size
	result := make([]string, len(data))

	// Process each item
	for i, item := range data {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			result[i] = p.processItem(item)
		}
	}

	return result, nil
}

// ProcessBatch processes data in batches for better memory management
func (p *processor) ProcessBatch(ctx context.Context, data []string, batchSize int) ([]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%w: empty data", ErrInvalidInput)
	}

	if batchSize <= 0 {
		batchSize = 100 // Default batch size
	}

	// Pre-allocate result slice
	result := make([]string, len(data))
	processed := 0

	// Process in batches
	for processed < len(data) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			end := processed + batchSize
			if end > len(data) {
				end = len(data)
			}

			// Process current batch
			for i := processed; i < end; i++ {
				result[i] = p.processItem(data[i])
			}

			processed = end
		}
	}

	return result, nil
}

// ProcessConcurrent processes data concurrently with worker pool
func (p *processor) ProcessConcurrent(ctx context.Context, data []string, workers int) ([]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("%w: empty data", ErrInvalidInput)
	}

	if workers <= 0 {
		workers = runtime.NumCPU() // Default to number of CPUs
	}

	// Create channels for job distribution
	jobs := make(chan int, len(data))
	results := make(chan struct {
		index int
		value string
	}, len(data))

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range jobs {
				select {
				case <-ctx.Done():
					return
				default:
					results <- struct {
						index int
						value string
					}{
						index: index,
						value: p.processItem(data[index]),
					}
				}
			}
		}()
	}

	// Send jobs
	go func() {
		defer close(jobs)
		for i := range data {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	result := make([]string, len(data))
	for r := range results {
		result[r.index] = r.value
	}

	return result, nil
}

// Run demonstrates the data processor with various operations
func Run() error {
	processor := NewProcessor()

	// Test data
	data := []string{
		"hello world",
		"go programming",
		"performance optimization",
		"concurrent processing",
		"memory management",
	}

	// Test cases
	testCases := []struct {
		name     string
		process  func() ([]string, error)
		expected int
	}{
		{
			name: "sequential processing",
			process: func() ([]string, error) {
				return processor.Process(context.Background(), data)
			},
			expected: len(data),
		},
		{
			name: "batch processing",
			process: func() ([]string, error) {
				return processor.ProcessBatch(context.Background(), data, 2)
			},
			expected: len(data),
		},
		{
			name: "concurrent processing",
			process: func() ([]string, error) {
				return processor.ProcessConcurrent(context.Background(), data, 2)
			},
			expected: len(data),
		},
	}

	// Run test cases
	for _, tc := range testCases {
		fmt.Printf("\nTesting: %s\n", tc.name)

		start := time.Now()
		result, err := tc.process()
		duration := time.Since(start)

		if err != nil {
			return fmt.Errorf("unexpected error: %w", err)
		}

		if len(result) != tc.expected {
			return fmt.Errorf("expected %d results, got %d", tc.expected, len(result))
		}

		fmt.Printf("Processed %d items in %v\n", len(result), duration)
		for i, item := range result {
			fmt.Printf("  %d: %s\n", i+1, item)
		}
	}

	return nil
}

// RunBenchmark demonstrates benchmarking different processing methods
func RunBenchmark() error {
	processor := NewProcessor()

	// Generate test data
	data := make([]string, 1000)
	for i := range data {
		data[i] = fmt.Sprintf("test data %d", i)
	}

	// Benchmark cases
	benchmarkCases := []struct {
		name    string
		process func() ([]string, error)
	}{
		{
			name: "sequential processing",
			process: func() ([]string, error) {
				return processor.Process(context.Background(), data)
			},
		},
		{
			name: "batch processing",
			process: func() ([]string, error) {
				return processor.ProcessBatch(context.Background(), data, 100)
			},
		},
		{
			name: "concurrent processing",
			process: func() ([]string, error) {
				return processor.ProcessConcurrent(context.Background(), data, runtime.NumCPU())
			},
		},
	}

	// Run benchmarks
	for _, bc := range benchmarkCases {
		fmt.Printf("\nBenchmarking: %s\n", bc.name)

		// Warm up
		for i := 0; i < 3; i++ {
			if _, err := bc.process(); err != nil {
				return fmt.Errorf("benchmark failed: %w", err)
			}
		}

		// Run benchmark
		start := time.Now()
		iterations := 10
		for i := 0; i < iterations; i++ {
			if _, err := bc.process(); err != nil {
				return fmt.Errorf("benchmark failed: %w", err)
			}
		}
		duration := time.Since(start)

		fmt.Printf("Processed %d items %d times in %v\n", len(data), iterations, duration)
		fmt.Printf("Average time per iteration: %v\n", duration/time.Duration(iterations))
		fmt.Printf("Average time per item: %v\n", duration/time.Duration(iterations*len(data)))
	}

	return nil
}
