package example7

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DependencyManager demonstrates various dependency management operations
type DependencyManager struct {
	projectPath string
	logger      *logrus.Logger
}

// NewDependencyManager creates a new dependency manager instance
func NewDependencyManager(projectPath string) *DependencyManager {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &DependencyManager{
		projectPath: projectPath,
		logger:      logger,
	}
}

// InitializeModule initializes a new Go module
func (dm *DependencyManager) InitializeModule(moduleName string) error {
	dm.logger.Info("Initializing new Go module")

	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = dm.projectPath
	return cmd.Run()
}

// AddDependency adds a new dependency with specific version
func (dm *DependencyManager) AddDependency(dependency string, version string) error {
	dm.logger.Infof("Adding dependency: %s@%s", dependency, version)

	cmd := exec.Command("go", "get", fmt.Sprintf("%s@%s", dependency, version))
	cmd.Dir = dm.projectPath
	return cmd.Run()
}

// UpdateDependencies updates all dependencies to their latest versions
func (dm *DependencyManager) UpdateDependencies() error {
	dm.logger.Info("Updating all dependencies")

	cmd := exec.Command("go", "get", "-u", "./...")
	cmd.Dir = dm.projectPath
	return cmd.Run()
}

// TidyDependencies removes unused dependencies
func (dm *DependencyManager) TidyDependencies() error {
	dm.logger.Info("Tidying dependencies")

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dm.projectPath
	return cmd.Run()
}

// VendorDependencies creates a vendor directory
func (dm *DependencyManager) VendorDependencies() error {
	dm.logger.Info("Creating vendor directory")

	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = dm.projectPath
	return cmd.Run()
}

// ListDependencies lists all dependencies
func (dm *DependencyManager) ListDependencies() (string, error) {
	dm.logger.Info("Listing all dependencies")

	cmd := exec.Command("go", "list", "-m", "all")
	cmd.Dir = dm.projectPath
	output, err := cmd.Output()
	return string(output), err
}

// WhyDependency explains why a dependency is needed
func (dm *DependencyManager) WhyDependency(dependency string) (string, error) {
	dm.logger.Infof("Explaining dependency: %s", dependency)

	cmd := exec.Command("go", "mod", "why", dependency)
	cmd.Dir = dm.projectPath
	output, err := cmd.Output()
	return string(output), err
}

// Run demonstrates dependency management operations
func Run() error {
	// Create a temporary directory for our example
	tempDir, err := os.MkdirTemp("", "go-deps-example-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Initialize dependency manager
	dm := NewDependencyManager(tempDir)

	// Initialize a new module
	moduleName := fmt.Sprintf("example.com/%s", uuid.New().String())
	if err := dm.InitializeModule(moduleName); err != nil {
		return fmt.Errorf("failed to initialize module: %w", err)
	}

	// Add some dependencies
	dependencies := []struct {
		name    string
		version string
	}{
		{"github.com/sirupsen/logrus", "v1.9.3"},
		{"github.com/google/uuid", "v1.4.0"},
	}

	for _, dep := range dependencies {
		if err := dm.AddDependency(dep.name, dep.version); err != nil {
			return fmt.Errorf("failed to add dependency %s: %w", dep.name, err)
		}
	}

	// Tidy dependencies
	if err := dm.TidyDependencies(); err != nil {
		return fmt.Errorf("failed to tidy dependencies: %w", err)
	}

	// List all dependencies
	deps, err := dm.ListDependencies()
	if err != nil {
		return fmt.Errorf("failed to list dependencies: %w", err)
	}
	fmt.Println("Current dependencies:")
	fmt.Println(deps)

	// Explain why we need logrus
	why, err := dm.WhyDependency("github.com/sirupsen/logrus")
	if err != nil {
		return fmt.Errorf("failed to explain dependency: %w", err)
	}
	fmt.Println("\nWhy we need logrus:")
	fmt.Println(why)

	// Create vendor directory
	if err := dm.VendorDependencies(); err != nil {
		return fmt.Errorf("failed to vendor dependencies: %w", err)
	}

	// Verify vendor directory
	vendorPath := filepath.Join(tempDir, "vendor")
	if _, err := os.Stat(vendorPath); err != nil {
		return fmt.Errorf("vendor directory not created: %w", err)
	}

	fmt.Println("\nDependency management operations completed successfully!")
	return nil
}

// RunBenchmark demonstrates dependency management performance
func RunBenchmark() error {
	// Create a temporary directory for benchmarking
	tempDir, err := os.MkdirTemp("", "go-deps-benchmark-*")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	dm := NewDependencyManager(tempDir)

	// Initialize module
	moduleName := fmt.Sprintf("example.com/%s", uuid.New().String())
	if err := dm.InitializeModule(moduleName); err != nil {
		return fmt.Errorf("failed to initialize module: %w", err)
	}

	// Benchmark operations
	operations := []struct {
		name string
		fn   func() error
	}{
		{"AddDependency", func() error {
			return dm.AddDependency("github.com/sirupsen/logrus", "v1.9.3")
		}},
		{"TidyDependencies", dm.TidyDependencies},
		{"VendorDependencies", dm.VendorDependencies},
		{"ListDependencies", func() error {
			_, err := dm.ListDependencies()
			return err
		}},
	}

	iterations := 5
	for _, op := range operations {
		fmt.Printf("\nBenchmarking: %s\n", op.name)

		// Warm up
		for i := 0; i < 2; i++ {
			if err := op.fn(); err != nil {
				return fmt.Errorf("benchmark failed: %w", err)
			}
		}

		// Run benchmark
		start := time.Now()
		for i := 0; i < iterations; i++ {
			if err := op.fn(); err != nil {
				return fmt.Errorf("benchmark failed: %w", err)
			}
		}
		duration := time.Since(start)

		fmt.Printf("Average time per operation: %v\n", duration/time.Duration(iterations))
	}

	return nil
}
