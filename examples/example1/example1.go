package example1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Bad interface design (what AI might suggest)
// This interface is too broad and violates interface segregation
type DataProcessor interface {
	Process(data []byte) ([]byte, error)
	Validate(data []byte) error
	Store(data []byte) error
	Retrieve(id string) ([]byte, error)
	Delete(id string) error
	Update(id string, data []byte) error
}

// Good interface design (segregated interfaces)
// Each interface has a single responsibility
type Validator interface {
	Validate(data []byte) error
}

type Processor interface {
	Process(data []byte) ([]byte, error)
}

type Storage interface {
	Store(data []byte) error
	Retrieve(id string) ([]byte, error)
}

// Example of interface composition
type DataService interface {
	Validator
	Processor
	Storage
}

// Implementation of the segregated interfaces
type JSONValidator struct{}

func (v *JSONValidator) Validate(data []byte) error {
	var js json.RawMessage
	return json.Unmarshal(data, &js)
}

type DataProcessorImpl struct {
	validator Validator
	storage   Storage
}

func NewDataProcessor(validator Validator, storage Storage) *DataProcessorImpl {
	return &DataProcessorImpl{
		validator: validator,
		storage:   storage,
	}
}

func (p *DataProcessorImpl) Process(data []byte) ([]byte, error) {
	if err := p.validator.Validate(data); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Process the data (example: add timestamp)
	processed := struct {
		Data      json.RawMessage `json:"data"`
		Timestamp time.Time       `json:"timestamp"`
	}{
		Data:      data,
		Timestamp: time.Now(),
	}

	return json.Marshal(processed)
}

// Example of a focused interface for a specific use case
type StreamProcessor interface {
	ProcessStream(ctx context.Context, reader io.Reader) error
}

// InMemoryStorage is a simple implementation of the Storage interface
type InMemoryStorage struct {
	data map[string][]byte
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		data: make(map[string][]byte),
	}
}

func (s *InMemoryStorage) Store(data []byte) error {
	// Implementation not shown for brevity
	return nil
}

func (s *InMemoryStorage) Retrieve(id string) ([]byte, error) {
	// Implementation not shown for brevity
	return nil, nil
}

// Run executes the example
func Run() error {
	// Create components
	validator := &JSONValidator{}
	storage := &InMemoryStorage{}
	processor := NewDataProcessor(validator, storage)

	// Example usage
	data := []byte(`{"name": "test", "value": 123}`)

	// Process data
	processed, err := processor.Process(data)
	if err != nil {
		return fmt.Errorf("failed to process data: %w", err)
	}

	fmt.Printf("Processed data: %s\n", processed)
	return nil
}
