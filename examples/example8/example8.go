package example8

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Custom error for demonstration
var ErrBadRequest = errors.New("bad request")

// User is a sample struct for JSON handling
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// handlerWithContext demonstrates context usage and error wrapping
func handlerWithContext(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := User{ID: 1, Name: "Gopher"}

	// Simulate work with context timeout
	select {
	case <-time.After(50 * time.Millisecond):
		// continue
	case <-ctx.Done():
		http.Error(w, fmt.Errorf("context cancelled: %w", ctx.Err()).Error(), http.StatusRequestTimeout)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Errorf("json encode error: %w", err).Error(), http.StatusInternalServerError)
	}
}

// middleware demonstrates HTTP middleware pattern
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// concurrentCounter demonstrates sync usage
func concurrentCounter(n int) int {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return count
}

// Run demonstrates advanced standard library usage
func Run() error {
	fmt.Println("--- Go Standard Library Example ---")

	// 1. Advanced net/http server setup
	mux := http.NewServeMux()
	mux.Handle("/user", middleware(http.HandlerFunc(handlerWithContext)))

	srv := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	// Start server in a goroutine
	done := make(chan struct{})
	go func() {
		fmt.Println("Starting HTTP server on :8081 (GET /user)...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
		close(done)
	}()

	// Give the server a moment to start
	time.Sleep(100 * time.Millisecond)

	// 2. Make a request with context and handle JSON
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/user", nil)
	if err != nil {
		return fmt.Errorf("request creation failed: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}
	fmt.Printf("Received user: %+v\n", user)

	// 3. Demonstrate sync usage
	count := concurrentCounter(100)
	fmt.Printf("Concurrent counter result: %d\n", count)

	// 4. Graceful shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}
	<-done
	fmt.Println("Server gracefully stopped.")

	return nil
}

// RunBenchmark demonstrates benchmarking standard library features
func RunBenchmark() error {
	fmt.Println("--- Go Standard Library Benchmark ---")
	start := time.Now()
	iters := 1000
	for i := 0; i < iters; i++ {
		_ = concurrentCounter(100)
	}
	duration := time.Since(start)
	fmt.Printf("Ran concurrentCounter 1000 times in %v\n", duration)
	fmt.Printf("Average per run: %v\n", duration/time.Duration(iters))
	return nil
}
