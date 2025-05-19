# Go Standard Library: AI's Common Oversights

## Mastering Go's Standard Library

Go's standard library is comprehensive and well-designed. While AI tools can suggest standard library usage, they often miss the deeper features and best practices.

### Common AI Standard Library Gaps

1. **Package Selection**: Not choosing the most appropriate package
2. **Feature Usage**: Missing advanced features
3. **Error Handling**: Not using standard library error patterns

### Staying Fluent

1. **Common Packages**:
   - `net/http`
   - `encoding/json`
   - `context`
   - `sync`
   - `time`

2. **Advanced Features**:
   ```go
   // AI might suggest:
   http.HandleFunc("/", handler)
   
   // Better:
   mux := http.NewServeMux()
   mux.Handle("/", http.HandlerFunc(handler))
   srv := &http.Server{
       Addr:    ":8080",
       Handler: mux,
   }
   ```

3. **Standard Patterns**:
   - Error wrapping
   - Context usage
   - HTTP middleware
   - JSON handling

### Key Takeaways

- Master common packages
- Understand advanced features
- Learn standard patterns
- Practice error handling
- Consider performance implications 