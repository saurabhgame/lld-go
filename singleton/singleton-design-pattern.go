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
