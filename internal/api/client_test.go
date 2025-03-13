package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/chscc/trtc-go/internal/models"
	"github.com/chscc/trtc-go/pkg/logger"
)

func TestClient_UploadFiles(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			t.Errorf("Expected method POST, got %s", r.Method)
		}

		// Check if the request has the correct content type
		contentType := r.Header.Get("Content-Type")
		if contentType == "" || contentType[:20] != "multipart/form-data;" {
			t.Errorf("Expected content type to start with multipart/form-data, got %s", contentType)
		}

		// Check if the request has the API key
		err := r.ParseMultipartForm(10 << 20) // 10 MB
		if err != nil {
			t.Fatalf("Failed to parse multipart form: %v", err)
		}

		// The API key is sent as "apikey" not "apiKey"
		if r.FormValue("apikey") != "test-api-key" {
			t.Errorf("Expected API key to be test-api-key, got %s", r.FormValue("apikey"))
		}

		// Return a successful response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return a properly formatted JSON response
		response := map[string]interface{}{
			"success": true,
			"message": "Upload successful",
			"code":    200,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "api-client-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a log file path
	logFilePath := filepath.Join(tempDir, "test.log")

	// Create a logger
	log, err := logger.New(logFilePath, logger.LevelInfo)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer log.Close()

	// Create a temporary file for testing
	testFilePath := filepath.Join(tempDir, "test.csv")
	if err := os.WriteFile(testFilePath, []byte("test,data"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a client with the test server URL
	client := NewClient(server.URL, false, log)

	// Create a test file
	testFile := models.UploadFile{
		Type:     models.FileTypeCourses,
		FilePath: testFilePath,
	}

	// Create a request
	request := models.UploadRequest{
		APIKey: "test-api-key",
		Files:  []models.UploadFile{testFile},
	}

	// Test uploading a file
	response, err := client.UploadFiles(request)
	if err != nil {
		t.Fatalf("Failed to upload file: %v", err)
	}

	// Check the response
	if !response.Success {
		t.Errorf("Expected success to be true, got %t", response.Success)
	}
	// We're not checking the exact message content since it contains the full JSON response
	if response.Code != 200 {
		t.Errorf("Expected code to be 200, got %d", response.Code)
	}
}

func TestClient_UploadFilesError(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		// Return a properly formatted JSON response
		response := map[string]interface{}{
			"success": false,
			"message": "Upload failed",
			"code":    400,
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "api-client-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a log file path
	logFilePath := filepath.Join(tempDir, "test.log")

	// Create a logger
	log, err := logger.New(logFilePath, logger.LevelInfo)
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer log.Close()

	// Create a temporary file for testing
	testFilePath := filepath.Join(tempDir, "test.csv")
	if err := os.WriteFile(testFilePath, []byte("test,data"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a client with the test server URL
	client := NewClient(server.URL, false, log)

	// Create a test file
	testFile := models.UploadFile{
		Type:     models.FileTypeCourses,
		FilePath: testFilePath,
	}

	// Create a request
	request := models.UploadRequest{
		APIKey: "test-api-key",
		Files:  []models.UploadFile{testFile},
	}

	// Test uploading a file
	response, err := client.UploadFiles(request)
	if err != nil {
		t.Fatalf("Failed to upload file: %v", err)
	}

	// Check the response
	if response.Success {
		t.Errorf("Expected success to be false, got %t", response.Success)
	}
	// We're not checking the exact message content since it contains the full JSON response
	if response.Code != 400 {
		t.Errorf("Expected code to be 400, got %d", response.Code)
	}
}

func TestMockClient_UploadFiles(t *testing.T) {
	// Create a mock client
	mockClient := &MockClient{
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

	// Create a test file
	testFile := models.UploadFile{
		Type:     models.FileTypeCourses,
		FilePath: "test.csv",
	}

	// Create a request
	request := models.UploadRequest{
		APIKey: "test-api-key",
		Files:  []models.UploadFile{testFile},
	}

	// Test uploading a file
	response, err := mockClient.UploadFiles(request)
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
