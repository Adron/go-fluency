package example5

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Common errors
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
)

// Task represents a task in the system
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TaskService defines the interface for task operations
type TaskService interface {
	Create(ctx context.Context, task *Task) error
	Get(ctx context.Context, id string) (*Task, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*Task, error)
}

// taskService implements the TaskService interface
type taskService struct {
	tasks map[string]*Task
}

// NewTaskService creates a new task service
func NewTaskService() TaskService {
	return &taskService{
		tasks: make(map[string]*Task),
	}
}

// validateTask validates a task
func (s *taskService) validateTask(task *Task) error {
	if task.Title == "" {
		return fmt.Errorf("%w: title cannot be empty", ErrInvalidInput)
	}

	if task.Status == "" {
		return fmt.Errorf("%w: status cannot be empty", ErrInvalidInput)
	}

	return nil
}

// Create creates a new task
func (s *taskService) Create(ctx context.Context, task *Task) error {
	if err := s.validateTask(task); err != nil {
		return err
	}

	if _, exists := s.tasks[task.ID]; exists {
		return fmt.Errorf("task with ID %s already exists", task.ID)
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now
	s.tasks[task.ID] = task
	return nil
}

// Get retrieves a task by ID
func (s *taskService) Get(ctx context.Context, id string) (*Task, error) {
	task, exists := s.tasks[id]
	if !exists {
		return nil, fmt.Errorf("%w: task with ID %s", ErrNotFound, id)
	}
	return task, nil
}

// Update updates an existing task
func (s *taskService) Update(ctx context.Context, task *Task) error {
	if err := s.validateTask(task); err != nil {
		return err
	}

	if _, exists := s.tasks[task.ID]; !exists {
		return fmt.Errorf("%w: task with ID %s", ErrNotFound, task.ID)
	}

	task.UpdatedAt = time.Now()
	s.tasks[task.ID] = task
	return nil
}

// Delete deletes a task by ID
func (s *taskService) Delete(ctx context.Context, id string) error {
	if _, exists := s.tasks[id]; !exists {
		return fmt.Errorf("%w: task with ID %s", ErrNotFound, id)
	}

	delete(s.tasks, id)
	return nil
}

// List returns all tasks
func (s *taskService) List(ctx context.Context) ([]*Task, error) {
	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Run demonstrates the task service with various operations
func Run() error {
	service := NewTaskService()

	// Test cases for task creation
	testCases := []struct {
		name    string
		task    *Task
		wantErr bool
	}{
		{
			name: "valid task",
			task: &Task{
				ID:          "1",
				Title:       "Complete project",
				Description: "Finish the Go package design example",
				Status:      "pending",
			},
			wantErr: false,
		},
		{
			name: "invalid title",
			task: &Task{
				ID:          "2",
				Title:       "",
				Description: "Invalid task",
				Status:      "pending",
			},
			wantErr: true,
		},
		{
			name: "invalid status",
			task: &Task{
				ID:          "3",
				Title:       "Another task",
				Description: "With invalid status",
				Status:      "",
			},
			wantErr: true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		fmt.Printf("\nTesting: %s\n", tc.name)

		// Create task
		err := service.Create(context.Background(), tc.task)
		if err != nil {
			if !tc.wantErr {
				return fmt.Errorf("unexpected error: %w", err)
			}
			fmt.Printf("Expected error: %v\n", err)
			continue
		}

		// Try to get the task
		task, err := service.Get(context.Background(), tc.task.ID)
		if err != nil {
			return fmt.Errorf("failed to get task: %w", err)
		}

		// Print task details
		fmt.Printf("Created task: %+v\n", task)
	}

	return nil
}

// RunIntegration demonstrates integration testing scenarios
func RunIntegration() error {
	service := NewTaskService()

	// Create a task
	task := &Task{
		ID:          "1",
		Title:       "Integration test task",
		Description: "Testing task operations",
		Status:      "pending",
	}

	fmt.Println("\nRunning integration test scenarios:")

	// Test scenario 1: Create and retrieve
	fmt.Println("\nScenario 1: Create and retrieve task")
	if err := service.Create(context.Background(), task); err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	retrieved, err := service.Get(context.Background(), task.ID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}
	fmt.Printf("Retrieved task: %+v\n", retrieved)

	// Test scenario 2: Update task
	fmt.Println("\nScenario 2: Update task")
	task.Status = "completed"
	if err := service.Update(context.Background(), task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	updated, err := service.Get(context.Background(), task.ID)
	if err != nil {
		return fmt.Errorf("failed to get updated task: %w", err)
	}
	fmt.Printf("Updated task: %+v\n", updated)

	// Test scenario 3: List tasks
	fmt.Println("\nScenario 3: List tasks")
	tasks, err := service.List(context.Background())
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}
	fmt.Printf("Total tasks: %d\n", len(tasks))

	// Test scenario 4: Delete task
	fmt.Println("\nScenario 4: Delete task")
	if err := service.Delete(context.Background(), task.ID); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	_, err = service.Get(context.Background(), task.ID)
	if err == nil {
		return fmt.Errorf("task should be deleted")
	}
	fmt.Printf("Task deleted successfully, error on retrieval: %v\n", err)

	return nil
}
