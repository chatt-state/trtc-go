package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logger levels
const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
)

// Logger represents a logger instance
type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	level         int
	file          *os.File
}

// New creates a new logger instance
func New(logFilePath string, level int) (*Logger, error) {
	// Create log directory if it doesn't exist
	logDir := filepath.Dir(logFilePath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Open log file
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	// Create multi-writer for console and file
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Create logger instance
	logger := &Logger{
		debugLogger:   log.New(multiWriter, "DEBUG: ", log.Ldate|log.Ltime),
		infoLogger:    log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime),
		warningLogger: log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime),
		errorLogger:   log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime),
		level:         level,
		file:          file,
	}

	// Log startup message
	logger.Info("Logger initialized at %s", time.Now().Format(time.RFC3339))

	return logger, nil
}

// Close closes the logger
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level <= LevelDebug {
		l.debugLogger.Printf(format, v...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level <= LevelInfo {
		l.infoLogger.Printf(format, v...)
	}
}

// Warning logs a warning message
func (l *Logger) Warning(format string, v ...interface{}) {
	if l.level <= LevelWarning {
		l.warningLogger.Printf(format, v...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	if l.level <= LevelError {
		l.errorLogger.Printf(format, v...)
	}
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level int) {
	l.level = level
}
