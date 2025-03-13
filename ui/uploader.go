package ui

import (
	"github.com/chscc/trtc-go/internal/config"
	"github.com/chscc/trtc-go/internal/models"
	"github.com/chscc/trtc-go/internal/uploader"
	"github.com/chscc/trtc-go/pkg/logger"
)

// Uploader is a wrapper around the uploader package for the UI
type Uploader struct {
	uploader *uploader.Uploader
}

// NewUploader creates a new uploader for the UI
func NewUploader(config *config.Config, logger *logger.Logger) *Uploader {
	return &Uploader{
		uploader: uploader.New(config, logger),
	}
}

// UploadFiles uploads files to the TRTC API
func (u *Uploader) UploadFiles(apiKey, coursesPath, equivalenciesPath, studentsPath, studentCoursesPath string) (*models.UploadResponse, error) {
	return u.uploader.UploadFilesFromPaths(apiKey, coursesPath, equivalenciesPath, studentsPath, studentCoursesPath)
}
