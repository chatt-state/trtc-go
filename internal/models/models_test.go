package models

import (
	"testing"
)

func TestFileTypeString(t *testing.T) {
	// Test FileType.String() method
	testCases := []struct {
		fileType FileType
		expected string
	}{
		{FileTypeCourses, "courses"},
		{FileTypeEquivalencies, "equivalencies"},
		{FileTypeStudents, "students"},
		{FileTypeStudentCourses, "studentcourses"},
		{FileType(99), "unknown"}, // Unknown file type
	}

	for _, tc := range testCases {
		result := tc.fileType.String()
		if result != tc.expected {
			t.Errorf("FileType.String() for %d returned %s, expected %s", tc.fileType, result, tc.expected)
		}
	}
}

func TestUploadFile(t *testing.T) {
	// Test UploadFile struct
	file := UploadFile{
		Type:     FileTypeCourses,
		FilePath: "/path/to/file.csv",
	}

	if file.Type != FileTypeCourses {
		t.Errorf("UploadFile.Type should be FileTypeCourses, got %d", file.Type)
	}
	if file.FilePath != "/path/to/file.csv" {
		t.Errorf("UploadFile.FilePath should be /path/to/file.csv, got %s", file.FilePath)
	}
}

func TestUploadRequest(t *testing.T) {
	// Test UploadRequest struct
	files := []UploadFile{
		{
			Type:     FileTypeCourses,
			FilePath: "/path/to/courses.csv",
		},
		{
			Type:     FileTypeEquivalencies,
			FilePath: "/path/to/equivalencies.csv",
		},
	}

	request := UploadRequest{
		APIKey: "test-api-key",
		Files:  files,
	}

	if request.APIKey != "test-api-key" {
		t.Errorf("UploadRequest.APIKey should be test-api-key, got %s", request.APIKey)
	}
	if len(request.Files) != 2 {
		t.Errorf("UploadRequest.Files should have 2 files, got %d", len(request.Files))
	}
	if request.Files[0].Type != FileTypeCourses {
		t.Errorf("UploadRequest.Files[0].Type should be FileTypeCourses, got %d", request.Files[0].Type)
	}
	if request.Files[1].Type != FileTypeEquivalencies {
		t.Errorf("UploadRequest.Files[1].Type should be FileTypeEquivalencies, got %d", request.Files[1].Type)
	}
}

func TestUploadResponse(t *testing.T) {
	// Test UploadResponse struct
	response := UploadResponse{
		Success: true,
		Message: "Upload successful",
		Code:    200,
	}

	if !response.Success {
		t.Errorf("UploadResponse.Success should be true, got %t", response.Success)
	}
	if response.Message != "Upload successful" {
		t.Errorf("UploadResponse.Message should be 'Upload successful', got %s", response.Message)
	}
	if response.Code != 200 {
		t.Errorf("UploadResponse.Code should be 200, got %d", response.Code)
	}
}
