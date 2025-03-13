package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	APIKey          string `mapstructure:"api_key"`
	APIEndpoint     string `mapstructure:"api_endpoint"`
	LogFile         string `mapstructure:"log_file"`
	IgnoreCertError bool   `mapstructure:"ignore_cert_error"`
}

// DefaultConfig returns a configuration with default values
func DefaultConfig() *Config {
	return &Config{
		APIKey:          "",
		APIEndpoint:     "https://rts.tnreversetransfer.org/api/Upload",
		LogFile:         "log.txt",
		IgnoreCertError: false,
	}
}

// getConfigDirFunc is a function type for getting the config directory
type getConfigDirFunc func() (string, error)

// getConfigDir returns the directory where the config file should be stored
var getConfigDir getConfigDirFunc = func() (string, error) {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("APPDATA environment variable not set")
		}
		configDir = filepath.Join(appData, "trtc-go")
	case "darwin":
		home := os.Getenv("HOME")
		if home == "" {
			return "", fmt.Errorf("HOME environment variable not set")
		}
		configDir = filepath.Join(home, "Library", "Application Support", "trtc-go")
	default: // linux and others
		home := os.Getenv("HOME")
		if home == "" {
			return "", fmt.Errorf("HOME environment variable not set")
		}
		configDir = filepath.Join(home, ".config", "trtc-go")
	}

	return configDir, nil
}

// LoadConfig loads the configuration from the config file
func LoadConfig() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, "config.yaml")

	// If config file doesn't exist, create it with default values
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}

		config := DefaultConfig()
		if err := SaveConfig(config); err != nil {
			return nil, fmt.Errorf("failed to save default config: %w", err)
		}
		return config, nil
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the configuration to the config file
func SaveConfig(config *Config) error {
	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	configPath := filepath.Join(configDir, "config.yaml")

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	viper.Set("api_key", config.APIKey)
	viper.Set("api_endpoint", config.APIEndpoint)
	viper.Set("log_file", config.LogFile)
	viper.Set("ignore_cert_error", config.IgnoreCertError)

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
