# Error Handling in Go: Best Practices and Patterns

Error handling is a fundamental aspect of Go programming. Unlike many other languages that use exceptions, Go takes a more explicit approach to error handling. This post explores best practices and patterns for handling errors in Go, with a focus on creating robust and maintainable code.

## The Error Interface

In Go, errors are values that implement the `error` interface:

```go
type error interface {
    Error() string
}
```

This simple interface allows for flexible error handling while maintaining type safety.

## Common Error Handling Patterns

### 1. Error Wrapping

Go 1.13 introduced error wrapping with `fmt.Errorf` and the `%w` verb:

```go
if err != nil {
    return fmt.Errorf("failed to process data: %w", err)
}
```

This preserves the original error while adding context.

### 2. Custom Error Types

For more complex error handling, create custom error types:

```go
type ValidationError struct {
    Field string
    Value interface{}
    Err   error
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field %s: %v", e.Field, e.Err)
}

func (e *ValidationError) Unwrap() error {
    return e.Err
}
```

### 3. Error Checking

Use `errors.Is` and `errors.As` for error type checking:

```go
if errors.Is(err, ErrNotFound) {
    // Handle not found error
}

var validationErr *ValidationError
if errors.As(err, &validationErr) {
    // Handle validation error
}
```

## Best Practices

1. **Always Check Errors**: Never ignore errors returned by functions.
2. **Add Context**: Wrap errors with additional context using `fmt.Errorf`.
3. **Use Custom Error Types**: Create specific error types for different error cases.
4. **Handle Errors at the Appropriate Level**: Don't handle errors too early or too late.
5. **Document Error Conditions**: Clearly document what errors your functions can return.

## Example Implementation

Let's look at a practical example that demonstrates these patterns. The example shows a user service with validation and processing errors:

```go
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

// ProcessingError represents a processing error
type ProcessingError struct {
    Operation string
    Err       error
    Timestamp time.Time
}

// UserService handles user operations
type UserService struct {
    users map[int]*User
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *User) error {
    // Validate the user
    if err := s.ValidateUser(user); err != nil {
        return &ProcessingError{
            Operation: "create_user",
            Err:      err,
            Timestamp: time.Now(),
        }
    }
    // ... rest of the implementation
}
```

## Running the Example

To run the error handling example:

```bash
go run main.go 3
```

This will demonstrate:
1. Basic error handling with custom error types
2. Error wrapping and unwrapping
3. Error type checking with `errors.Is` and `errors.As`
4. Proper error context and propagation

## Conclusion

Effective error handling is crucial for building reliable Go applications. By following these patterns and best practices, you can create more maintainable and robust code. Remember to:

- Always check errors
- Add appropriate context
- Use custom error types when needed
- Handle errors at the right level
- Document error conditions

The example code demonstrates these concepts in a practical way, showing how to implement proper error handling in a real-world scenario.

## Further Reading

- [Go Error Handling](https://go.dev/blog/error-handling-and-go)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Error Handling in Go](https://go.dev/doc/effective_go#errors)

### AI-Suggested vs. Proper Code Examples

1. **Basic Error Handling**:
```go
// AI-Suggested (Problematic):
func processUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    }
    if user.Name == "" {
        return errors.New("name is empty")
    }
    return nil
}

// Proper Implementation:
func processUser(user *User) error {
    if user == nil {
        return fmt.Errorf("invalid user: %w", ErrInvalidInput)
    }
    if user.Name == "" {
        return &ValidationError{
            Field: "name",
            Value: user.Name,
            Err:   ErrInvalidInput,
        }
    }
    return nil
}
```

2. **Error Wrapping**:
```go
// AI-Suggested (Problematic):
func fetchUser(id int) (*User, error) {
    user, err := db.GetUser(id)
    if err != nil {
        return nil, err
    }
    return user, nil
}

// Proper Implementation:
func fetchUser(id int) (*User, error) {
    user, err := db.GetUser(id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("user %d: %w", id, ErrNotFound)
        }
        return nil, fmt.Errorf("failed to fetch user %d: %w", id, err)
    }
    return user, nil
}
```

### ⚠️ Warnings About AI Tool Over-Reliance

1. **Error Context Loss**:
   - AI tools often strip error context
   - They may not preserve error chains
   - They might miss important error types
   - They often don't consider error wrapping

2. **Common Pitfalls**:
   - Generic error messages without context
   - Missing error type information
   - Improper error propagation
   - Ignored errors in defer statements

3. **Real-World Issues**:
   - Debugging difficulties due to lost error context
   - Production issues from swallowed errors
   - Inconsistent error handling patterns
   - Missing error type information in logs

### Maintaining Fluency Through Practice

1. **Daily Exercises**:
   - Write error handling without AI assistance
   - Review and refactor error handling in existing code
   - Practice different error patterns
   - Implement proper error wrapping

2. **Code Review Checklist**:
   - Verify error context preservation
   - Check for proper error wrapping
   - Ensure consistent error patterns
   - Validate error type usage
   - Review error documentation

3. **Learning Resources**:
   - Go Error Handling (https://go.dev/blog/error-handling-and-go)
   - Working with Errors in Go 1.13 (https://go.dev/blog/go1.13-errors)
   - Error Handling in Go (https://go.dev/doc/effective_go#errors)
   - Go Error Handling Patterns (talks)

### Real-World AI-Generated Code Issues

1. **Case Study: Lost Error Context**:
```go
// AI-Generated Code (Problematic):
func processData(data []byte) error {
    if err := validate(data); err != nil {
        return err  // Lost context!
    }
    return nil
}

// Fixed Version:
func processData(data []byte) error {
    if err := validate(data); err != nil {
        return fmt.Errorf("failed to validate data: %w", err)
    }
    return nil
}
```

2. **Case Study: Improper Error Type**:
```go
// AI-Generated Code (Problematic):
func getUser(id int) (*User, error) {
    if id <= 0 {
        return nil, errors.New("invalid id")  // Generic error!
    }
    // ...
}

// Fixed Version:
func getUser(id int) (*User, error) {
    if id <= 0 {
        return nil, &ValidationError{
            Field: "id",
            Value: id,
            Err:   ErrInvalidInput,
        }
    }
    // ...
}
``` 