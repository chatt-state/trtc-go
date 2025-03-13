package api

import (
	"github.com/chatt-state/trtc-go/internal/models"
)

// MockClient is a mock implementation of the API client for testing
type MockClient struct {
	UploadFilesFunc func(request models.UploadRequest) (*models.UploadResponse, error)
}

// UploadFiles calls the mock implementation
func (m *MockClient) UploadFiles(request models.UploadRequest) (*models.UploadResponse, error) {
	return m.UploadFilesFunc(request)
}
