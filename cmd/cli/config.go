package main

import (
	"fmt"

	"github.com/chatt-state/trtc-go/internal/config"
	"github.com/spf13/cobra"
)

// Command line flags for config command
var (
	endpoint        string
	logFile         string
	ignoreCertError bool
)

// newConfigCmd creates a new config command
func newConfigCmd() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the application",
		Long:  `Configure the application settings such as API endpoint, log file, and certificate validation.`,
	}

	// Add subcommands
	configCmd.AddCommand(newConfigGetCmd())
	configCmd.AddCommand(newConfigSetCmd())

	return configCmd
}

// newConfigGetCmd creates a new config get command
func newConfigGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "Get configuration values",
		Long:  `Get the current configuration values.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfigGet()
		},
	}

	return getCmd
}

// newConfigSetCmd creates a new config set command
func newConfigSetCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set configuration values",
		Long:  `Set configuration values such as API endpoint, log file, and certificate validation.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runConfigSet(cmd)
		},
	}

	// Add flags
	setCmd.Flags().StringVar(&endpoint, "endpoint", "", "API endpoint URL")
	setCmd.Flags().StringVar(&logFile, "logfile", "", "Log file path")
	setCmd.Flags().BoolVar(&ignoreCertError, "ignore-cert-error", false, "Ignore certificate errors")

	return setCmd
}

// runConfigGet runs the config get command
func runConfigGet() error {
	fmt.Println("Current Configuration:")
	fmt.Printf("API Endpoint: %s\n", Config.APIEndpoint)
	fmt.Printf("Log File: %s\n", Config.LogFile)
	fmt.Printf("Ignore Certificate Errors: %t\n", Config.IgnoreCertError)
	return nil
}

// runConfigSet runs the config set command
func runConfigSet(cmd *cobra.Command) error {
	// Check if any flags were set
	flagsSet := false

	// Update endpoint if provided
	if endpoint != "" {
		Config.APIEndpoint = endpoint
		flagsSet = true
		Logger.Info("API endpoint set to %s", endpoint)
	}

	// Update log file if provided
	if logFile != "" {
		Config.LogFile = logFile
		flagsSet = true
		Logger.Info("Log file set to %s", logFile)
	}

	// Update ignore cert error if provided
	if cmd.Flags().Changed("ignore-cert-error") {
		Config.IgnoreCertError = ignoreCertError
		flagsSet = true
		Logger.Info("Ignore certificate errors set to %t", ignoreCertError)
	}

	// If no flags were set, print current configuration
	if !flagsSet {
		fmt.Println("No configuration values were provided. Current configuration:")
		return runConfigGet()
	}

	// Save configuration
	if err := config.SaveConfig(Config); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	fmt.Println("Configuration saved successfully.")
	return nil
}
