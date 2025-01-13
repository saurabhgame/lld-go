# Singleton Design Pattern


*The Singleton Design Pattern ensures that a class has only one instance and provides
a global point of access to it. It's commonly used when exactly one object is needed
to coordinate actions across a system.*


### Key Characteristics
1. Single Instance: Ensures that only one instance of the class is created.
2. Global Access Point: Provides a way to access the instance globally.
3. Lazy Initialization (Optional): Creates the instance only when it is needed.

### When to Use Singleton
1. Managing shared resources like database connections, logging, or configuration settings.
2. Enforcing a single point of control in the system, e.g., managing application state.


```go 
package main

import (
	"fmt"
	"sync"
)

// Logger is a singleton class for logging messages
type Logger struct {
	logs []string
}

var (
	instance *Logger
	once     sync.Once // Ensures the singleton is created only once
)

// GetInstance provides a global access point to the singleton instance
func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{logs: []string{}}
	})
	return instance
}

// LogMessage adds a log message to the logger
func (l *Logger) LogMessage(message string) {
	l.logs = append(l.logs, message)
	fmt.Println("Log:", message)
}

// ShowLogs prints all the logged messages
func (l *Logger) ShowLogs() {
	fmt.Println("All Logs:")
	for _, log := range l.logs {
		fmt.Println(log)
	}
}

func main() {
	// Get the singleton instance
	logger1 := GetInstance()
	logger1.LogMessage("First message")
	logger1.LogMessage("Second message")

	// Get the same instance
	logger2 := GetInstance()
	logger2.LogMessage("Third message")

	// Verify that logger1 and logger2 are the same instance
	if logger1 == logger2 {
		fmt.Println("Logger1 and Logger2 are the same instance.")
	}

	// Show all logged messages
	logger2.ShowLogs()
}

```
```go
Log: Starting the application
Log: Performing some task
Log: Task completed
Logger is a singleton instance.
All Logs:
- Starting the application
- Performing some task
- Task completed
```

## Without Singleton: Logging Example

### What’s the Problem?
1. Multiple Instances: logger1 and logger2 are separate instances, so their logs are stored independently.
2. Inconsistent State: Logs are not centralized, leading to potential confusion and inconsistency.


### With the Singleton pattern:

1. Single Instance: Only one logger instance exists, ensuring all logs are stored in the same place.
2. Global Access: The logger can be accessed consistently across the entire application.


```go
package main

import "fmt"

// Logger represents a simple logger
type Logger struct {
	messages []string
}

// Log adds a message to the logger
func (l *Logger) Log(message string) {
	l.messages = append(l.messages, message)
	fmt.Println("Log:", message)
}

// ShowLogs displays all logged messages
func (l *Logger) ShowLogs() {
	fmt.Println("All Logs:")
	for _, msg := range l.messages {
		fmt.Println("-", msg)
	}
}

func main() {
	// Create multiple logger instances
	logger1 := &Logger{}
	logger2 := &Logger{}

	// Log messages with logger1
	logger1.Log("Message from logger1")

	// Log messages with logger2
	logger2.Log("Message from logger2")

	// Show logs
	fmt.Println("\nLogger1 Logs:")
	logger1.ShowLogs()

	fmt.Println("\nLogger2 Logs:")
	logger2.ShowLogs()
}
```

```go
Log: Message from logger1
Log: Message from logger2

Logger1 Logs:
- Message from logger1

Logger2 Logs:
- Message from logger2
```


