package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	// Test default configuration
	config := DefaultConfig()

	// Check default values
	if config.APIKey != "" {
		t.Errorf("Default API key should be empty, got %s", config.APIKey)
	}
	if config.APIEndpoint != "https://rts.tnreversetransfer.org/api/Upload" {
		t.Errorf("Default API endpoint should be https://rts.tnreversetransfer.org/api/Upload, got %s", config.APIEndpoint)
	}
	if config.LogFile != "log.txt" {
		t.Errorf("Default log file should be log.txt, got %s", config.LogFile)
	}
	if config.IgnoreCertError != false {
		t.Errorf("Default ignore cert error should be false, got %t", config.IgnoreCertError)
	}
}

func TestSaveAndLoadConfig(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "config-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a mock getConfigDir function for testing
	origGetConfigDir := getConfigDir
	getConfigDir = func() (string, error) {
		return tempDir, nil
	}
	// Restore the original function after the test
	defer func() {
		getConfigDir = origGetConfigDir
	}()

	// Create a test configuration
	testConfig := &Config{
		APIKey:          "test-api-key",
		APIEndpoint:     "https://test-endpoint.com",
		LogFile:         "test-log.txt",
		IgnoreCertError: true,
	}

	// Save the configuration
	err = SaveConfig(testConfig)
	if err != nil {
		t.Fatalf("Failed to save configuration: %v", err)
	}

	// Check if the config file was created
	configPath := filepath.Join(tempDir, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("Config file was not created: %v", err)
	}

	// Load the configuration
	loadedConfig, err := LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	// Check if the loaded configuration matches the saved configuration
	if loadedConfig.APIKey != testConfig.APIKey {
		t.Errorf("Loaded API key does not match: expected %s, got %s", testConfig.APIKey, loadedConfig.APIKey)
	}
	if loadedConfig.APIEndpoint != testConfig.APIEndpoint {
		t.Errorf("Loaded API endpoint does not match: expected %s, got %s", testConfig.APIEndpoint, loadedConfig.APIEndpoint)
	}
	if loadedConfig.LogFile != testConfig.LogFile {
		t.Errorf("Loaded log file does not match: expected %s, got %s", testConfig.LogFile, loadedConfig.LogFile)
	}
	if loadedConfig.IgnoreCertError != testConfig.IgnoreCertError {
		t.Errorf("Loaded ignore cert error does not match: expected %t, got %t", testConfig.IgnoreCertError, loadedConfig.IgnoreCertError)
	}
}
