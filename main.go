package main

import (
	"flag"
	"fmt"
	"os"

	"practice/examples/example1"
	"practice/examples/example2"
	"practice/examples/example3"
	"practice/examples/example4"
	"practice/examples/example5"
	"practice/examples/example6"
	"practice/examples/example7"
	"practice/examples/example8"
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
	example5Flag := flag.Bool("example5", false, "Run example 5 (Package Design)")
	example5IntegrationFlag := flag.Bool("example5-integration", false, "Run example 5 integration tests")
	example6Flag := flag.Bool("example6", false, "Run example 6 (Performance Optimization)")
	example6BenchmarkFlag := flag.Bool("example6-benchmark", false, "Run example 6 benchmarks")
	example7Flag := flag.Bool("example7", false, "Run example 7 (Dependency Management)")
	example7BenchmarkFlag := flag.Bool("example7-benchmark", false, "Run example 7 benchmarks")
	example8Flag := flag.Bool("example8", false, "Run example 8 (Standard Library)")
	example8BenchmarkFlag := flag.Bool("example8-benchmark", false, "Run example 8 benchmarks")
	flag.Parse()

	// Check if any example flag is set
	if !*example1Flag && !*example2Flag && !*example2PipelineFlag && !*example3Flag && !*example4Flag && !*example4BenchmarkFlag && !*example4IntegrationFlag && !*example5Flag && !*example5IntegrationFlag && !*example6Flag && !*example6BenchmarkFlag && !*example7Flag && !*example7BenchmarkFlag && !*example8Flag && !*example8BenchmarkFlag {
		fmt.Println("Please specify an example to run:")
		fmt.Println("  --example1              Run example 1 (Interface Design)")
		fmt.Println("  --example2              Run example 2 (Worker Pool)")
		fmt.Println("  --example2-pipeline     Run example 2 (Pipeline Pattern)")
		fmt.Println("  --example3              Run example 3 (Error Handling)")
		fmt.Println("  --example4              Run example 4 (Table-Driven Tests)")
		fmt.Println("  --example4-benchmark    Run example 4 (Benchmarks)")
		fmt.Println("  --example4-integration  Run example 4 (Integration Tests)")
		fmt.Println("  --example5              Run example 5 (Package Design)")
		fmt.Println("  --example5-integration  Run example 5 (Integration Tests)")
		fmt.Println("  --example6              Run example 6 (Performance Optimization)")
		fmt.Println("  --example6-benchmark    Run example 6 (Benchmarks)")
		fmt.Println("  --example7              Run example 7 (Dependency Management)")
		fmt.Println("  --example7-benchmark    Run example 7 (Benchmarks)")
		fmt.Println("  --example8              Run example 8 (Standard Library)")
		fmt.Println("  --example8-benchmark    Run example 8 (Benchmarks)")
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

	if *example5Flag {
		fmt.Println("Running package design example:")
		if err := example5.Run(); err != nil {
			fmt.Printf("Error running package design example: %v\n", err)
		}
	}

	if *example5IntegrationFlag {
		fmt.Println("Running package design integration tests:")
		if err := example5.RunIntegration(); err != nil {
			fmt.Printf("Error running package design integration tests: %v\n", err)
		}
	}

	if *example6Flag {
		fmt.Println("Running performance optimization example:")
		if err := example6.Run(); err != nil {
			fmt.Printf("Error running performance optimization example: %v\n", err)
		}
	}

	if *example6BenchmarkFlag {
		fmt.Println("Running performance benchmarks:")
		if err := example6.RunBenchmark(); err != nil {
			fmt.Printf("Error running performance benchmarks: %v\n", err)
		}
	}

	if *example7Flag {
		fmt.Println("Running dependency management example:")
		if err := example7.Run(); err != nil {
			fmt.Printf("Error running dependency management example: %v\n", err)
		}
	}

	if *example7BenchmarkFlag {
		fmt.Println("Running dependency management benchmarks:")
		if err := example7.RunBenchmark(); err != nil {
			fmt.Printf("Error running dependency management benchmarks: %v\n", err)
		}
	}

	if *example8Flag {
		fmt.Println("Running standard library example:")
		if err := example8.Run(); err != nil {
			fmt.Printf("Error running standard library example: %v\n", err)
		}
	}

	if *example8BenchmarkFlag {
		fmt.Println("Running standard library benchmarks:")
		if err := example8.RunBenchmark(); err != nil {
			fmt.Printf("Error running standard library benchmarks: %v\n", err)
		}
	}
}
