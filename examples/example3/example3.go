package example3

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Common errors
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
)

// ValidationError represents a validation error
type ValidationError struct {
	Field string
	Value interface{}
	Err   error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field %s with value %v: %v", e.Field, e.Value, e.Err)
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}

// ProcessingError represents a processing error
type ProcessingError struct {
	Operation string
	Err       error
	Timestamp time.Time
}

func (e *ProcessingError) Error() string {
	return fmt.Sprintf("processing error during %s at %v: %v", e.Operation, e.Timestamp, e.Err)
}

func (e *ProcessingError) Unwrap() error {
	return e.Err
}

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
		return &ValidationError{
			Field: "name",
			Value: user.Name,
			Err:   fmt.Errorf("name cannot be empty"),
		}
	}

	if user.Age < 0 || user.Age > 150 {
		return &ValidationError{
			Field: "age",
			Value: user.Age,
			Err:   fmt.Errorf("age must be between 0 and 150"),
		}
	}

	return nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *User) error {
	// Validate the user
	if err := s.ValidateUser(user); err != nil {
		return &ProcessingError{
			Operation: "create_user",
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	// Check if user already exists
	if _, exists := s.users[user.ID]; exists {
		return &ProcessingError{
			Operation: "create_user",
			Err:       fmt.Errorf("user with ID %d already exists", user.ID),
			Timestamp: time.Now(),
		}
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
		return nil, &ProcessingError{
			Operation: "get_user",
			Err:       fmt.Errorf("%w: user with ID %d", ErrNotFound, id),
			Timestamp: time.Now(),
		}
	}
	return user, nil
}

// Run demonstrates error handling patterns
func Run() error {
	service := NewUserService()

	// Test cases
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

			// Handle different error types
			var validationErr *ValidationError
			var processingErr *ProcessingError

			switch {
			case errors.As(err, &validationErr):
				fmt.Printf("Validation error: %v\n", validationErr)
			case errors.As(err, &processingErr):
				fmt.Printf("Processing error: %v\n", processingErr)
			default:
				fmt.Printf("Unknown error: %v\n", err)
			}
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

// RunErrorWrapping demonstrates error wrapping patterns
func RunErrorWrapping() error {
	// Example of error wrapping
	err := processData("invalid")
	if err != nil {
		// Unwrap the error chain
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Unwrapped: %v\n", errors.Unwrap(err))

		// Check for specific error types
		if errors.Is(err, ErrInvalidInput) {
			fmt.Println("Error is ErrInvalidInput")
		}

		var validationErr *ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("Error is ValidationError: %v\n", validationErr)
		}
	}

	return nil
}

func processData(input string) error {
	// Try to parse the input
	num, err := strconv.Atoi(input)
	if err != nil {
		return &ValidationError{
			Field: "input",
			Value: input,
			Err:   fmt.Errorf("%w: %v", ErrInvalidInput, err),
		}
	}

	if num < 0 {
		return &ValidationError{
			Field: "input",
			Value: num,
			Err:   fmt.Errorf("%w: number must be positive", ErrInvalidInput),
		}
	}

	return nil
}
