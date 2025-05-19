package example4

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Common errors
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

// UserService handles user operations
type UserService struct {
	users map[int]*User
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

// ValidateUser validates a user
func (s *UserService) ValidateUser(user *User) error {
	if user.Name == "" {
		return fmt.Errorf("%w: name cannot be empty", ErrInvalidInput)
	}

	if user.Age < 0 || user.Age > 150 {
		return fmt.Errorf("%w: age must be between 0 and 150", ErrInvalidInput)
	}

	return nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *User) error {
	// Validate the user
	if err := s.ValidateUser(user); err != nil {
		return err
	}

	// Check if user already exists
	if _, exists := s.users[user.ID]; exists {
		return fmt.Errorf("user with ID %d already exists", user.ID)
	}

	// Create the user
	user.CreatedAt = time.Now()
	s.users[user.ID] = user
	return nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id int) (*User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("%w: user with ID %d", ErrNotFound, id)
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *User) error {
	// Validate the user
	if err := s.ValidateUser(user); err != nil {
		return err
	}

	// Check if user exists
	if _, exists := s.users[user.ID]; !exists {
		return fmt.Errorf("%w: user with ID %d", ErrNotFound, user.ID)
	}

	// Update the user
	s.users[user.ID] = user
	return nil
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int) error {
	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("%w: user with ID %d", ErrNotFound, id)
	}

	delete(s.users, id)
	return nil
}

// Run demonstrates the user service with various operations
func Run() error {
	service := NewUserService()

	// Test cases for user creation
	testCases := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		{
			name: "valid user",
			user: &User{
				ID:   1,
				Name: "John Doe",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			user: &User{
				ID:   2,
				Name: "",
				Age:  30,
			},
			wantErr: true,
		},
		{
			name: "invalid age",
			user: &User{
				ID:   3,
				Name: "Jane Doe",
				Age:  -1,
			},
			wantErr: true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		fmt.Printf("\nTesting: %s\n", tc.name)

		// Create user
		err := service.CreateUser(tc.user)
		if err != nil {
			if !tc.wantErr {
				return fmt.Errorf("unexpected error: %w", err)
			}
			fmt.Printf("Expected error: %v\n", err)
			continue
		}

		// Try to get the user
		user, err := service.GetUser(tc.user.ID)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		// Print user details
		userJSON, _ := json.MarshalIndent(user, "", "  ")
		fmt.Printf("Created user: %s\n", userJSON)
	}

	return nil
}

// RunBenchmark demonstrates benchmarking the user service
func RunBenchmark() error {
	service := NewUserService()

	// Create a test user
	user := &User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	// Benchmark user creation
	fmt.Println("\nBenchmarking user creation:")
	start := time.Now()
	for i := 0; i < 1000; i++ {
		user.ID = i
		if err := service.CreateUser(user); err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
	}
	duration := time.Since(start)
	fmt.Printf("Created 1000 users in %v\n", duration)

	// Benchmark user retrieval
	fmt.Println("\nBenchmarking user retrieval:")
	start = time.Now()
	for i := 0; i < 1000; i++ {
		if _, err := service.GetUser(i); err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}
	}
	duration = time.Since(start)
	fmt.Printf("Retrieved 1000 users in %v\n", duration)

	return nil
}

// RunIntegration demonstrates integration testing scenarios
func RunIntegration() error {
	service := NewUserService()

	// Create a user
	user := &User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	fmt.Println("\nRunning integration test scenarios:")

	// Test scenario 1: Create and retrieve
	fmt.Println("\nScenario 1: Create and retrieve user")
	if err := service.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	retrieved, err := service.GetUser(user.ID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	fmt.Printf("Retrieved user: %+v\n", retrieved)

	// Test scenario 2: Update user
	fmt.Println("\nScenario 2: Update user")
	user.Name = "John Updated"
	if err := service.UpdateUser(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	updated, err := service.GetUser(user.ID)
	if err != nil {
		return fmt.Errorf("failed to get updated user: %w", err)
	}
	fmt.Printf("Updated user: %+v\n", updated)

	// Test scenario 3: Delete user
	fmt.Println("\nScenario 3: Delete user")
	if err := service.DeleteUser(user.ID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	_, err = service.GetUser(user.ID)
	if err == nil {
		return fmt.Errorf("user should not exist after deletion")
	}
	fmt.Printf("User deleted successfully, error on retrieval: %v\n", err)

	return nil
}
