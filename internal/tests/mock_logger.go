package tests

import (
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"sync"
)

type MockLogger struct {
	mu      sync.Mutex
	entries []LogEntry
	fields  map[string]interface{}
}

// NewMockLogger returns a new mock logger
func NewMockLogger() *MockLogger {
	return &MockLogger{
		entries: []LogEntry{},
		fields:  map[string]interface{}{},
	}
}

type LogEntry struct {
	Level   string
	Format  string
	Args    []interface{}
	Message string
}

func (m *MockLogger) log(level string, format string, v ...interface{}) runtime.Logger {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.entries = append(m.entries, LogEntry{
		Level:   level,
		Format:  format,
		Args:    v,
		Message: fmt.Sprintf(format, v...),
	})

	return m
}

func (m *MockLogger) Debug(format string, v ...interface{}) {
	m.log("DEBUG", format, v...)
}

func (m *MockLogger) Info(format string, v ...interface{}) {
	m.log("INFO", format, v...)
}

func (m *MockLogger) Warn(format string, v ...interface{}) {
	m.log("WARN", format, v...)
}

func (m *MockLogger) Error(format string, v ...interface{}) {
	m.log("ERROR", format, v...)
}

func (m *MockLogger) WithField(key string, v interface{}) runtime.Logger {
	return m.WithFields(map[string]interface{}{key: v})
}

func (m *MockLogger) WithFields(fields map[string]interface{}) runtime.Logger {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m
}

func (m *MockLogger) Fields() map[string]interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()

	cp := make(map[string]interface{}, len(m.fields))
	for k, v := range m.fields {
		cp[k] = v
	}
	return cp
}

// Entries returns all captured log entries
func (m *MockLogger) Entries() []LogEntry {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]LogEntry(nil), m.entries...)
}

func (m *MockLogger) HasError(message string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, entry := range m.entries {
		if entry.Level == "ERROR" && entry.Message == message {
			return true
		}
	}

	return false
}
