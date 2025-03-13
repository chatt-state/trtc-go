package uploader

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/chatt-state/trtc-go/internal/api"
	"github.com/chatt-state/trtc-go/internal/config"
	"github.com/chatt-state/trtc-go/internal/models"
	"github.com/chatt-state/trtc-go/pkg/logger"
)

func setupTest(t *testing.T) (*logger.Logger, *config.Config, string) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "uploader-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Create a log file path
	logFilePath := filepath.Join(tempDir, "test.log")

	// Create a logger
	logger, err := logger.New(logFilePath, logger.LevelInfo)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	// Create a configuration
	config := &config.Config{
		APIKey:          "test-api-key",
		APIEndpoint:     "https://test-endpoint.com",
		LogFile:         logFilePath,
		IgnoreCertError: false,
	}

	return logger, config, tempDir
}

func TestUploadFiles(t *testing.T) {
	// Setup test
	logger, config, tempDir := setupTest(t)
	defer os.RemoveAll(tempDir)
	defer logger.Close()

	// Create a mock API client
	mockClient := &api.MockClient{
		UploadFilesFunc: func(request models.UploadRequest) (*models.UploadResponse, error) {
			// Check if the request is correct
			if request.APIKey != "test-api-key" {
				t.Errorf("Expected API key to be test-api-key, got %s", request.APIKey)
			}
			if len(request.Files) != 1 {
				t.Errorf("Expected 1 file, got %d", len(request.Files))
			}

			// Return a successful response
			return &models.UploadResponse{
				Success: true,
				Message: "Upload successful",
				Code:    200,
			}, nil
		},
	}

	// Create an uploader with the mock client
	uploader := NewWithClient(mockClient, config, logger)

	// Create a temporary file for testing
	testFilePath := filepath.Join(tempDir, "test.csv")
	if err := os.WriteFile(testFilePath, []byte("test,data"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a test file
	testFile := models.UploadFile{
		Type:     models.FileTypeCourses,
		FilePath: testFilePath,
	}

	// Test uploading a file
	response, err := uploader.UploadFiles("test-api-key", []models.UploadFile{testFile})
	if err != nil {
		t.Fatalf("Failed to upload file: %v", err)
	}

	// Check the response
	if !response.Success {
		t.Errorf("Expected success to be true, got %t", response.Success)
	}
	if response.Message != "Upload successful" {
		t.Errorf("Expected message to be 'Upload successful', got %s", response.Message)
	}
	if response.Code != 200 {
		t.Errorf("Expected code to be 200, got %d", response.Code)
	}
}

func TestUploadFilesError(t *testing.T) {
	// Setup test
	logger, config, tempDir := setupTest(t)
	defer os.RemoveAll(tempDir)
	defer logger.Close()

	// Create a mock API client that returns an error
	mockClient := &api.MockClient{
		UploadFilesFunc: func(request models.UploadRequest) (*models.UploadResponse, error) {
			return nil, errors.New("upload failed")
		},
	}

	// Create an uploader with the mock client
	uploader := NewWithClient(mockClient, config, logger)

	// Create a temporary file for testing
	testFilePath := filepath.Join(tempDir, "test.csv")
	if err := os.WriteFile(testFilePath, []byte("test,data"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a test file
	testFile := models.UploadFile{
		Type:     models.FileTypeCourses,
		FilePath: testFilePath,
	}

	// Test uploading a file
	_, err := uploader.UploadFiles("test-api-key", []models.UploadFile{testFile})
	if err == nil {
		t.Fatalf("Expected an error, got nil")
	}
	if err.Error() != "upload failed" {
		t.Errorf("Expected error message to be 'upload failed', got %s", err.Error())
	}
}

func TestUploadFilesFromPaths(t *testing.T) {
	// Setup test
	logger, config, tempDir := setupTest(t)
	defer os.RemoveAll(tempDir)
	defer logger.Close()

	// Create a mock API client
	mockClient := &api.MockClient{
		UploadFilesFunc: func(request models.UploadRequest) (*models.UploadResponse, error) {
			// Check if the request is correct
			if request.APIKey != "test-api-key" {
				t.Errorf("Expected API key to be test-api-key, got %s", request.APIKey)
			}
			if len(request.Files) != 2 {
				t.Errorf("Expected 2 files, got %d", len(request.Files))
			}

			// Return a successful response
			return &models.UploadResponse{
				Success: true,
				Message: "Upload successful",
				Code:    200,
			}, nil
		},
	}

	// Create an uploader with the mock client
	uploader := NewWithClient(mockClient, config, logger)

	// Create temporary files for testing
	coursesFilePath := filepath.Join(tempDir, "courses.csv")
	if err := os.WriteFile(coursesFilePath, []byte("test,data"), 0644); err != nil {
		t.Fatalf("Failed to create courses file: %v", err)
	}

	equivalenciesFilePath := filepath.Join(tempDir, "equivalencies.csv")
	if err := os.WriteFile(equivalenciesFilePath, []byte("test,data"), 0644); err != nil {
		t.Fatalf("Failed to create equivalencies file: %v", err)
	}

	// Test uploading files from paths
	response, err := uploader.UploadFilesFromPaths("test-api-key", coursesFilePath, equivalenciesFilePath, "", "")
	if err != nil {
		t.Fatalf("Failed to upload files from paths: %v", err)
	}

	// Check the response
	if !response.Success {
		t.Errorf("Expected success to be true, got %t", response.Success)
	}
	if response.Message != "Upload successful" {
		t.Errorf("Expected message to be 'Upload successful', got %s", response.Message)
	}
	if response.Code != 200 {
		t.Errorf("Expected code to be 200, got %d", response.Code)
	}
}

func TestUploadFilesValidation(t *testing.T) {
	// Setup test
	logger, config, tempDir := setupTest(t)
	defer os.RemoveAll(tempDir)
	defer logger.Close()

	// Create an uploader
	uploader := New(config, logger)

	// Test with empty API key
	_, err := uploader.UploadFiles("", []models.UploadFile{})
	if err == nil {
		t.Fatalf("Expected an error for empty API key, got nil")
	}
	if err.Error() != "API key is required" {
		t.Errorf("Expected error message to be 'API key is required', got %s", err.Error())
	}

	// Test with empty files
	_, err = uploader.UploadFiles("test-api-key", []models.UploadFile{})
	if err == nil {
		t.Fatalf("Expected an error for empty files, got nil")
	}
	if err.Error() != "at least one file is required" {
		t.Errorf("Expected error message to be 'at least one file is required', got %s", err.Error())
	}

	// Test with non-existent file
	_, err = uploader.UploadFiles("test-api-key", []models.UploadFile{
		{
			Type:     models.FileTypeCourses,
			FilePath: filepath.Join(tempDir, "non-existent.csv"),
		},
	})
	if err == nil {
		t.Fatalf("Expected an error for non-existent file, got nil")
	}
	if err.Error() != "file does not exist: "+filepath.Join(tempDir, "non-existent.csv") {
		t.Errorf("Expected error message to be 'file does not exist: %s', got %s", filepath.Join(tempDir, "non-existent.csv"), err.Error())
	}
}
