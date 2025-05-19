package main

import (
	"flag"
	"fmt"
	"os"

	"practice/examples/example1"
	"practice/examples/example2"
	"practice/examples/example3"
	"practice/examples/example4"
)

func main() {
	// Define command line flags
	example1Flag := flag.Bool("example1", false, "Run example 1 (Interface Design)")
	example2Flag := flag.Bool("example2", false, "Run example 2 (Concurrency Patterns)")
	example2PipelineFlag := flag.Bool("example2-pipeline", false, "Run example 2 pipeline pattern")
	example3Flag := flag.Bool("example3", false, "Run example 3 (Error Handling)")
	example4Flag := flag.Bool("example4", false, "Run example 4 (Testing Strategies)")
	example4BenchmarkFlag := flag.Bool("example4-benchmark", false, "Run example 4 benchmarks")
	example4IntegrationFlag := flag.Bool("example4-integration", false, "Run example 4 integration tests")
	flag.Parse()

	// Check if any example flag is set
	if !*example1Flag && !*example2Flag && !*example2PipelineFlag && !*example3Flag && !*example4Flag && !*example4BenchmarkFlag && !*example4IntegrationFlag {
		fmt.Println("Please specify an example to run:")
		fmt.Println("  --example1              Run example 1 (Interface Design)")
		fmt.Println("  --example2              Run example 2 (Worker Pool)")
		fmt.Println("  --example2-pipeline     Run example 2 (Pipeline Pattern)")
		fmt.Println("  --example3              Run example 3 (Error Handling)")
		fmt.Println("  --example4              Run example 4 (Table-Driven Tests)")
		fmt.Println("  --example4-benchmark    Run example 4 (Benchmarks)")
		fmt.Println("  --example4-integration  Run example 4 (Integration Tests)")
		os.Exit(1)
	}

	// Run the selected example
	if *example1Flag {
		if err := example1.Run(); err != nil {
			fmt.Printf("Error running example 1: %v\n", err)
		}
	}

	if *example2Flag {
		if err := example2.Run(); err != nil {
			fmt.Printf("Error running example 2: %v\n", err)
		}
	}

	if *example2PipelineFlag {
		if err := example2.RunPipeline(); err != nil {
			fmt.Printf("Error running example 2 pipeline: %v\n", err)
		}
	}

	if *example3Flag {
		fmt.Println("Running error handling example:")
		if err := example3.Run(); err != nil {
			fmt.Printf("Error running example 3: %v\n", err)
		}
		fmt.Println("\nRunning error wrapping example:")
		if err := example3.RunErrorWrapping(); err != nil {
			fmt.Printf("Error running error wrapping example: %v\n", err)
		}
	}

	if *example4Flag {
		fmt.Println("Running table-driven tests example:")
		if err := example4.Run(); err != nil {
			fmt.Printf("Error running table-driven tests: %v\n", err)
		}
	}

	if *example4BenchmarkFlag {
		fmt.Println("Running benchmarks example:")
		if err := example4.RunBenchmark(); err != nil {
			fmt.Printf("Error running benchmarks: %v\n", err)
		}
	}

	if *example4IntegrationFlag {
		fmt.Println("Running integration tests example:")
		if err := example4.RunIntegration(); err != nil {
			fmt.Printf("Error running integration tests: %v\n", err)
		}
	}
}
