# Go Interfaces Code Kata: Building a Message Processing System

## Objective
Practice implementing and using Go interfaces by building a message processing system that can handle different types of messages (text, JSON, and binary) through a clean interface design.

## Problem Description
You are building a message processing system that needs to handle different types of messages. The system should be able to:
1. Validate incoming messages
2. Process messages based on their type
3. Store processed messages
4. Retrieve messages when needed

## Requirements

### 1. Core Interfaces
Create the following interfaces:
```go
type MessageValidator interface {
    Validate(data []byte) error
}

type MessageProcessor interface {
    Process(data []byte) ([]byte, error)
}

type MessageStorage interface {
    Store(data []byte) error
    Retrieve(id string) ([]byte, error)
}
```

### 2. Message Types
Implement three different message types:
- TextMessage: Plain text messages
- JSONMessage: JSON formatted messages
- BinaryMessage: Binary data messages

### 3. Processing Rules
- Text messages should be converted to uppercase
- JSON messages should be validated and pretty-printed
- Binary messages should be base64 encoded

## Tasks

### Level 1: Basic Implementation
1. Implement the `MessageValidator` interface for each message type
2. Create a basic `MessageProcessor` that can handle text messages
3. Implement a simple in-memory `MessageStorage`

### Level 2: Advanced Features
1. Add support for JSON and binary messages
2. Implement error handling and custom error types
3. Add message metadata (timestamp, type, size)

### Level 3: Interface Composition
1. Create a `MessageService` interface that composes the other interfaces
2. Implement a `MessagePipeline` that can process messages in sequence
3. Add support for message transformation chains

## Example Usage
```go
// Example of how the final implementation might be used
func main() {
    // Create components
    validator := NewMessageValidator()
    processor := NewMessageProcessor()
    storage := NewMessageStorage()
    
    // Create service
    service := NewMessageService(validator, processor, storage)
    
    // Process a text message
    textMsg := []byte("hello world")
    processed, err := service.ProcessMessage(textMsg)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Processed message: %s\n", processed)
}
```

## Success Criteria
1. All interfaces are properly implemented
2. Each message type is handled correctly
3. Error handling is comprehensive
4. Code is well-tested
5. Interfaces are small and focused
6. Implementation follows Go best practices

## Bonus Challenges
1. Add support for message encryption/decryption
2. Implement a message queue system
3. Add support for message validation rules
4. Create a message transformation pipeline
5. Implement message compression

## Tips
- Start with the smallest possible interfaces
- Use interface composition to build more complex functionality
- Write tests for each interface implementation
- Consider using the `io.Reader` and `io.Writer` interfaces where appropriate
- Remember to handle errors properly

## Learning Objectives
- Understanding Go interface design
- Practicing interface composition
- Implementing error handling
- Working with different data types
- Writing clean, maintainable code

## Time Estimate
- Level 1: 1-2 hours
- Level 2: 2-3 hours
- Level 3: 3-4 hours
- Bonus challenges: 4+ hours

## Resources
- Go documentation on interfaces
- `io` package documentation
- Go error handling patterns
- Interface composition examples 