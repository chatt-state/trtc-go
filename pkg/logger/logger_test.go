package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "logger-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a log file path
	logFilePath := filepath.Join(tempDir, "test.log")

	// Test creating a new logger
	logger, err := New(logFilePath, LevelInfo)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Check if the log file was created
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		t.Fatalf("Log file was not created: %v", err)
	}

	// Test logging at different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")

	// Read the log file
	content, err := os.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Convert content to string
	logContent := string(content)

	// Debug message should not be in the log file (level is INFO)
	if strings.Contains(logContent, "This is a debug message") {
		t.Errorf("Debug message should not be in the log file")
	}

	// Info, warning, and error messages should be in the log file
	if !strings.Contains(logContent, "This is an info message") {
		t.Errorf("Info message should be in the log file")
	}
	if !strings.Contains(logContent, "This is a warning message") {
		t.Errorf("Warning message should be in the log file")
	}
	if !strings.Contains(logContent, "This is an error message") {
		t.Errorf("Error message should be in the log file")
	}
}

func TestSetLevel(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "logger-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a log file path
	logFilePath := filepath.Join(tempDir, "test.log")

	// Test creating a new logger
	logger, err := New(logFilePath, LevelInfo)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Close()

	// Test setting the level
	logger.SetLevel(LevelWarning)

	// Test logging at different levels
	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warning("This is a warning message")
	logger.Error("This is an error message")

	// Read the log file
	content, err := os.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Convert content to string
	logContent := string(content)

	// Debug and info messages should not be in the log file (level is WARNING)
	if strings.Contains(logContent, "This is a debug message") {
		t.Errorf("Debug message should not be in the log file")
	}
	if strings.Contains(logContent, "This is an info message") {
		t.Errorf("Info message should not be in the log file")
	}

	// Warning and error messages should be in the log file
	if !strings.Contains(logContent, "This is a warning message") {
		t.Errorf("Warning message should be in the log file")
	}
	if !strings.Contains(logContent, "This is an error message") {
		t.Errorf("Error message should be in the log file")
	}
}
