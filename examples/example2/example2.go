package example2

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Worker represents a worker that processes jobs
type Worker struct {
	id      int
	ctx     context.Context
	wg      *sync.WaitGroup
	jobs    <-chan int
	results chan<- int
}

// NewWorker creates a new worker
func NewWorker(id int, ctx context.Context, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) *Worker {
	return &Worker{
		id:      id,
		ctx:     ctx,
		wg:      wg,
		jobs:    jobs,
		results: results,
	}
}

// Start begins processing jobs
func (w *Worker) Start() {
	go func() {
		defer w.wg.Done()
		for {
			select {
			case <-w.ctx.Done():
				fmt.Printf("Worker %d shutting down\n", w.id)
				return
			case job, ok := <-w.jobs:
				if !ok {
					return
				}
				// Simulate work
				time.Sleep(100 * time.Millisecond)
				result := job * 2
				w.results <- result
				fmt.Printf("Worker %d processed job %d, result: %d\n", w.id, job, result)
			}
		}
	}()
}

// Run demonstrates various concurrency patterns
func Run() error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create channels
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Create wait group
	var wg sync.WaitGroup

	// Start workers
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, ctx, &wg, jobs, results)
		worker.Start()
	}

	// Start result collector
	go func() {
		for result := range results {
			fmt.Printf("Received result: %d\n", result)
		}
	}()

	// Send jobs
	go func() {
		defer close(jobs)
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
				fmt.Printf("Sent job %d\n", i)
			}
		}
	}()

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Check if we timed out
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("processing timed out")
	}

	return nil
}

// RunPipeline demonstrates a pipeline pattern
func RunPipeline() error {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create channels for the pipeline
	numbers := make(chan int)
	squares := make(chan int)
	results := make(chan int)

	// Start the pipeline stages
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			select {
			case <-ctx.Done():
				return
			case numbers <- i:
				fmt.Printf("Generated: %d\n", i)
			}
		}
	}()

	go func() {
		defer close(squares)
		for n := range numbers {
			select {
			case <-ctx.Done():
				return
			case squares <- n * n:
				fmt.Printf("Squared: %d\n", n*n)
			}
		}
	}()

	go func() {
		defer close(results)
		for s := range squares {
			select {
			case <-ctx.Done():
				return
			case results <- s + 1:
				fmt.Printf("Final result: %d\n", s+1)
			}
		}
	}()

	// Collect results
	for r := range results {
		fmt.Printf("Pipeline output: %d\n", r)
	}

	return nil
}
