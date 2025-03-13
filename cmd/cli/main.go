package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chatt-state/trtc-go/internal/config"
	"github.com/chatt-state/trtc-go/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	// Version is the version of the application
	Version = "0.0.1"

	// Config is the application configuration
	Config *config.Config

	// Logger is the application logger
	Logger *logger.Logger

	// Command line flags
	logLevel int
)

func main() {
	// Create root command
	rootCmd := &cobra.Command{
		Use:     "trtc-go",
		Short:   "TRTC-Go is a tool for uploading files to the TRTC API",
		Version: Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Load configuration
			var err error
			Config, err = config.LoadConfig()
			if err != nil {
				return fmt.Errorf("failed to load configuration: %w", err)
			}

			// Create logger
			logFilePath := Config.LogFile
			if !filepath.IsAbs(logFilePath) {
				// If log file path is not absolute, make it relative to the current directory
				logFilePath, err = filepath.Abs(logFilePath)
				if err != nil {
					return fmt.Errorf("failed to get absolute path for log file: %w", err)
				}
			}

			Logger, err = logger.New(logFilePath, logLevel)
			if err != nil {
				return fmt.Errorf("failed to create logger: %w", err)
			}

			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			// Close logger
			if Logger != nil {
				Logger.Close()
			}
		},
	}

	// Add persistent flags
	rootCmd.PersistentFlags().IntVar(&logLevel, "log-level", logger.LevelInfo, "Log level (0=DEBUG, 1=INFO, 2=WARNING, 3=ERROR)")

	// Add commands
	rootCmd.AddCommand(newUploadCmd())
	rootCmd.AddCommand(newConfigCmd())

	// Execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
