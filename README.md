# The Go (Golang) Programming Language Handbook

Go (commonly referred to as Golang) is a robust, high-performance, and modern programming language created by Google engineers. This handbook covers all essential knowledge about Go programming, from basic concepts to advanced techniques.

## Table of Contents

- [Getting Started](docs/1.getting-started/1.0_getting-started.md)

  - [Installation](docs/1.getting-started/1.1_installation.md)
  - [First Program](docs/1.getting-started/1.2_first-program.md)
  - [Go Workspace](docs/1.getting-started/1.3_workspace.md)
  - [Go Modules](docs/1.getting-started/1.4_modules.md)
  - [Go Tools](docs/1.getting-started/1.5_tools.md)

- [Language Basics](docs/2.language-basics/2.0_basics.md)

  - [Variables](docs/2.language-basics/2.1_variables.md)
  - [Data Types](docs/2.language-basics/2.2_data_types.md)
  - [Constants](docs/2.language-basics/2.3_constants.md)
  - [Operators](docs/2.language-basics/2.4_operators.md)
  - [Control Flow](docs/2.language-basics/2.5_control_flow.md)
  - [Functions](docs/2.language-basics/2.6_functions.md)
  - [Packages](docs/2.language-basics/2.7_packages.md)
  - [Error Handling](docs/2.language-basics/2.8_error_handling.md)
  - [Type Conversions](docs/2.language-basics/2.9_type_conversions.md)
  - [Defer and Panic](docs/2.language-basics/2.10_defer_panic.md)

- [Data Structures](docs/3.data-structures/3.0_data_structures.md)

  - [Arrays](docs/3.data-structures/3.1_arrays.md)
  - [Slices](docs/3.data-structures/3.2_slice.md)
  - [Maps](docs/3.data-structures/3.3_maps.md)
  - [Structs](docs/3.data-structures/3.4_structs.md)
  - [Pointers](docs/3.data-structures/3.5_pointers.md)
  - [Interfaces](docs/3.data-structures/3.6_interfaces.md)
  - [Type Embedding](docs/3.data-structures/3.7_type_embedding.md)
  - [Custom Types](docs/3.data-structures/3.8_custom_types.md)

- [Concurrency](docs/4.concurrency/4.0_concurrency.md)

  - [Goroutines](docs/4.concurrency/4.1_goroutines.md)
  - [Channels](docs/4.concurrency/4.2_channels.md)
  - [Select](docs/4.concurrency/4.3_select.md)
  - [Mutexes](docs/4.concurrency/4.4_mutexes.md)
  - [Context](docs/4.concurrency/4.5_context.md)
  - [Race Conditions](docs/4.concurrency/4.6_race_conditions.md)
  - [Worker Pools](docs/4.concurrency/4.7_worker_pools.md)
  - [Atomic Operations](docs/4.concurrency/4.8_atomic.md)
  - [Timeouts](docs/4.concurrency/4.9_timeouts.md)
  - [Rate Limiting](docs/4.concurrency/4.10_rate_limiting.md)
  - [Synchronization](docs/4.concurrency/4.11_synchronization.md)

- [Standard Library](docs/5.standard-library/5.0_standard_library.md)

  - [I/O Operations](docs/5.standard-library/5.1_io.md)
  - [File Handling](docs/5.standard-library/5.2_files.md)
  - [Time and Date](docs/5.standard-library/5.3_time.md)
  - [JSON Processing](docs/5.standard-library/5.4_json.md)
  - [Regular Expressions](docs/5.standard-library/5.5_regex.md)
  - HTTP Client
  - HTTP Server
  - Template Engine

- Testing

  - Unit Testing
  - Table Tests
  - Benchmarking
  - Mocking
  - Test Coverage
  - HTTP Testing
  - Integration Tests

- Web Development

  - Web Servers
  - Routing
  - Middleware
  - Sessions
  - Authentication
  - REST APIs
  - GraphQL
  - WebSockets

- Database Access

  - SQL Basics
  - Connection Pooling
  - Migrations
  - ORMs
  - NoSQL
  - Transactions
  - Query Building

- Best Practices

  - Project Structure
  - Error Management
  - Logging
  - Configuration
  - Performance
  - Security
  - Code Style

## The Go Advantage

Go merges the best features from various programming languages:

- **Clarity**: Clean syntax and minimal language features
- **Speed**: Compiled language with garbage collection
- **Concurrency**: Built-in support via goroutines and channels
- **Reliability**: Robust typing system with built-in memory protection
- **Efficiency**: Fast compilation and comprehensive standard library
- **Modern**: Built for today's microservices, cloud-native applications, and distributed systems

## Language Overview

```go
package main

import (
    "fmt"
    "sync"
)

// Demonstration of parallel processing
func main() {
    var wg sync.WaitGroup
    messages := make(chan string, 3)

    // Execute parallel  tasks
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, messages, &wg)
    }

    // Close channel when done
    go func() {
        wg.Wait()
        close(messages)
    }()

    // Collect results
    for msg := range messages {
        fmt.Println(msg)
    }
}

func worker(id int, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    ch <- fmt.Sprintf("Worker %d completed task", id)
}
```

## [Getting Started](docs/1.getting-started/1.0_getting-started.md)

Begin your Go journey with our comprehensive getting started guide:

### 1. [Installation](docs/1.getting-started/1.1_installation.md):

- System setup
- Environment configuration
- Verification steps

### 2. [First Steps](docs/1.getting-started/1.2_first-program.md)

- Basic syntax
- Running programs
- Development workflow

### 3. [Workspace Setup](docs/1.getting-started/1.3_workspace.md)

- Project organization
- Code structure
- Best practices

### 4. [Go Modules](docs/1.getting-started/1.4_modules.md)

- Dependency management
- Version control
- Package distribution

### 5. [Development Tools](docs/1.getting-started/1.5_tools.md)

- Command-line utilities
- IDE integration
- Debugging tools

## [Core Programming Concepts](docs/2.language-basics/2.0_basics.md)

Build mastery in Go's essential elements:

### 1. [Variables and Data Types](docs/2.language-basics/2.1_variables.md)

```go
var language string = "Golang"
version := 1.23 // Automatic type detection
const MaxUsers = 1000
```

### 2. Flow Control

```go
if gender == "male" {
    fmt.Println("You're a man!")
} else if gender == "female" {
    fmt.Println("You're a woman!")
} else {
    fmt.Println("Other!")
}

for index, value := range items {
    fmt.Printf("Index: %d, Value: %v\n", index, value)
}
```

### 3. Functions

```go
func multiply(a, b int) int {
    return a * b
}

// Function with multiple return values
func processData(input string) (result string, err error) {
    if len(input) == 0 {
        return "", fmt.Errorf("empty input provided")
    }
    return strings.ToUpper(input), nil
}
```

## Data Structures

Master Go's fundamental data structures:

### 1. Arrays and Slices

```go
// Array
var scores [3]int = [3]int{95, 87, 92}

// Slice
items := []string{"apple", "banana", "cherry"}
items = append(items, "date", "elderberry")
```

### 2. Maps

```go
inventory := map[string]int{
    "laptops": 15,
    "phones":  32,
    "tablets": 8,
}

// Check for existence
if count, exists := inventory["laptops"]; exists {
    fmt.Printf("Laptops in stock: %d\n", count)
}
```

3. Structs

```go
type Employee struct {
    ID       int
    FullName string
    Email    string
    Salary   float64
}

// Method definition
func (e Employee) GetDetails() string {
    return fmt.Sprintf("%s (%s)", e.FullName, e.Email)
}

emp := Employee{
    ID:       101,
    FullName: "Jane Smith",
    Email:    "jane@company.com",
    Salary:   75000.0,
}
```

## Concurrency

Explore Go's exceptional concurrency capabilities:

### 1. Goroutines

```go
func processTask(taskID int) {
    fmt.Printf("Processing task %d\n", taskID)
    time.Sleep(time.Millisecond * 500)
    fmt.Printf("Task %d completed\n", taskID)
}

// Launch parallel execution
for i := 1; i <= 5; i++ {
    go processTask(i)
}
```

### 2. Channels

```go
results := make(chan string, 10)

go func() {
    for i := 1; i <= 5; i++ {
        results <- fmt.Sprintf("Result %d", i)
    }
    close(results)
}()

// Consume results
for result := range results {
    fmt.Println("Received:", result)
}
```

### 3. Select

```go
timeout := time.After(time.Second * 2)
ticker := time.NewTicker(time.Millisecond * 500)

for {
    select {
    case <-ticker.C:
        fmt.Println("Tick")
    case <-timeout:
        fmt.Println("Operation timed out")
        return
    }
}
```

## Standard Library

Discover Go's rich standard library:

### 1. IO Operations

```go
content := []byte("Hello, Go programming!")
err := ioutil.WriteFile("greeting.txt", content, 0644)
if err != nil {
    log.Printf("Error writing file: %v", err)
}

data, err := ioutil.ReadFile("greeting.txt")
if err == nil {
    fmt.Println(string(data))
}
```

### 2. HTTP Server

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Web!")
})
http.ListenAndServe(":8080", nil)
```

### 3. JSON Processing

```go
type User struct {
    Name string `json:"name"`
    Age int     `json:"age"`
}

json.Marshal(user)
```

## Testing

Learn about Go's testing capabilities:

### 1. Unit Testing

```go
func TestAdd(t *testing.T) {
    if got := add(2, 3); got != 5 {
        t.Errorf("add(2, 3) = %v; want 5", got)
    }
}
```

### 2. Benchmarking

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(2, 3)
    }
}
```

## Web Development

Build web applications with Go:

1. Web Servers
2. REST APIs
3. Middleware
4. Authentication

## Database Access

Work with databases in Go:

1. SQL Basics
2. ORMs
3. Migrations
4. Transactions

## Best Practices

Follow Go best practices:

1. Project Structure
2. Error Handling
3. Performance
4. Security

## Learning Path

### 1. Beginner

- [Installation and setup](docs/1.getting-started/1.1_installation.md)
- [Basic syntax and types](docs/2.language-basics/2.1_variables.md)
- Control structures
- Functions and packages

### 2. Intermediate

- Interface and methods
- Error handling
- Testing
- Concurrency basics

### 3. Advanced

- Advanced concurrency patterns
- Memory management
- Performance optimization
- Systems programming

## Next Steps

Start with the [Getting Started](docs/1.getting-started/1.0.getting-started.md) guide

## Additional Resources

- [Go Documentation](https://go.dev/doc)
- [Go by Example](https://gobyexample.com)
- [Go Tour](https://go.dev/tour/welcome/1)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Playground](https://go.dev/play)
- [WebReference - Guide to Go](https://webreference.com/go)
