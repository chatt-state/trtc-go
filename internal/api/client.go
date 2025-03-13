package api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/chscc/trtc-go/internal/models"
	"github.com/chscc/trtc-go/pkg/logger"
)

// APIClient is an interface for the API client
type APIClient interface {
	UploadFiles(request models.UploadRequest) (*models.UploadResponse, error)
}

// Client represents an API client for the TRTC API
type Client struct {
	endpoint        string
	httpClient      *http.Client
	logger          *logger.Logger
	ignoreCertError bool
}

// NewClient creates a new API client
func NewClient(endpoint string, ignoreCertError bool, logger *logger.Logger) *Client {
	// Create HTTP client with custom transport if needed
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: ignoreCertError,
		},
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Minute, // 10 minutes timeout, same as original
	}

	return &Client{
		endpoint:        endpoint,
		httpClient:      httpClient,
		logger:          logger,
		ignoreCertError: ignoreCertError,
	}
}

// UploadFiles uploads files to the TRTC API
func (c *Client) UploadFiles(request models.UploadRequest) (*models.UploadResponse, error) {
	c.logger.Info("Uploading files to %s", c.endpoint)

	// Create multipart form
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Add API key
	if err := w.WriteField("apikey", request.APIKey); err != nil {
		return nil, fmt.Errorf("failed to write API key: %w", err)
	}

	// Add files
	for _, file := range request.Files {
		c.logger.Info("Adding file: %s (type: %s)", file.FilePath, file.Type.String())

		// Open file
		f, err := os.Open(file.FilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s: %w", file.FilePath, err)
		}
		defer f.Close()

		// Create form file
		fw, err := w.CreateFormFile(file.Type.String(), filepath.Base(file.FilePath))
		if err != nil {
			return nil, fmt.Errorf("failed to create form file: %w", err)
		}

		// Copy file content to form
		if _, err = io.Copy(fw, f); err != nil {
			return nil, fmt.Errorf("failed to copy file content: %w", err)
		}
	}

	// Close multipart writer
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", c.endpoint, &b)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	// Send request
	c.logger.Info("Sending request to %s", c.endpoint)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Create response
	response := &models.UploadResponse{
		Success: resp.StatusCode == http.StatusOK,
		Message: string(body),
		Code:    resp.StatusCode,
	}

	c.logger.Info("Response: %d %s", resp.StatusCode, resp.Status)
	if !response.Success {
		c.logger.Error("Upload failed: %s", response.Message)
	} else {
		c.logger.Info("Upload successful")
	}

	return response, nil
}
