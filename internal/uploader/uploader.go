package uploader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chscc/trtc-go/internal/api"
	"github.com/chscc/trtc-go/internal/config"
	"github.com/chscc/trtc-go/internal/models"
	"github.com/chscc/trtc-go/pkg/logger"
)

// Uploader handles file uploads to the TRTC API
type Uploader struct {
	client api.APIClient
	config *config.Config
	logger *logger.Logger
}

// New creates a new uploader
func New(config *config.Config, logger *logger.Logger) *Uploader {
	client := api.NewClient(config.APIEndpoint, config.IgnoreCertError, logger)
	return &Uploader{
		client: client,
		config: config,
		logger: logger,
	}
}

// NewWithClient creates a new uploader with a custom API client
func NewWithClient(client api.APIClient, config *config.Config, logger *logger.Logger) *Uploader {
	return &Uploader{
		client: client,
		config: config,
		logger: logger,
	}
}

// UploadFiles uploads files to the TRTC API
func (u *Uploader) UploadFiles(apiKey string, files []models.UploadFile) (*models.UploadResponse, error) {
	// Validate API key
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	// Validate files
	if len(files) == 0 {
		return nil, fmt.Errorf("at least one file is required")
	}

	// Check if files exist
	for _, file := range files {
		if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist: %s", file.FilePath)
		}
	}

	// Create upload request
	request := models.UploadRequest{
		APIKey: apiKey,
		Files:  files,
	}

	// Upload files
	return u.client.UploadFiles(request)
}

// UploadFilesFromPaths uploads files to the TRTC API from file paths
func (u *Uploader) UploadFilesFromPaths(apiKey string, coursesPath, equivalenciesPath, studentsPath, studentCoursesPath string) (*models.UploadResponse, error) {
	var files []models.UploadFile

	// Add courses file if provided
	if coursesPath != "" {
		absPath, err := filepath.Abs(coursesPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path for courses file: %w", err)
		}
		files = append(files, models.UploadFile{
			Type:     models.FileTypeCourses,
			FilePath: absPath,
		})
	}

	// Add equivalencies file if provided
	if equivalenciesPath != "" {
		absPath, err := filepath.Abs(equivalenciesPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path for equivalencies file: %w", err)
		}
		files = append(files, models.UploadFile{
			Type:     models.FileTypeEquivalencies,
			FilePath: absPath,
		})
	}

	// Add students file if provided
	if studentsPath != "" {
		absPath, err := filepath.Abs(studentsPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path for students file: %w", err)
		}
		files = append(files, models.UploadFile{
			Type:     models.FileTypeStudents,
			FilePath: absPath,
		})
	}

	// Add student courses file if provided
	if studentCoursesPath != "" {
		absPath, err := filepath.Abs(studentCoursesPath)
		if err != nil {
			return nil, fmt.Errorf("failed to get absolute path for student courses file: %w", err)
		}
		files = append(files, models.UploadFile{
			Type:     models.FileTypeStudentCourses,
			FilePath: absPath,
		})
	}

	// Upload files
	return u.UploadFiles(apiKey, files)
}
