package main

import (
	"fmt"
	"os"

	"github.com/chatt-state/trtc-go/internal/uploader"
	"github.com/spf13/cobra"
)

// Command line flags for upload command
var (
	apiKey             string
	coursesPath        string
	equivalenciesPath  string
	studentsPath       string
	studentCoursesPath string
)

// newUploadCmd creates a new upload command
func newUploadCmd() *cobra.Command {
	uploadCmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload files to the TRTC API",
		Long: `Upload files to the Tennessee Reverse Transfer Consortium (TRTC) API.
You can upload courses, equivalencies, students, and student courses files.
At least one file must be specified.`,
		Example: `  # Upload a courses file
  trtc-go upload -apikey="your-api-key" -courses="path/to/courses.csv"

  # Upload multiple file types
  trtc-go upload -apikey="your-api-key" -courses="path/to/courses.csv" -equivalencies="path/to/equivalencies.csv"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpload()
		},
	}

	// Add flags
	uploadCmd.Flags().StringVar(&apiKey, "apikey", "", "API key for authentication (required)")
	uploadCmd.Flags().StringVar(&coursesPath, "courses", "", "Path to courses file")
	uploadCmd.Flags().StringVar(&equivalenciesPath, "equivalencies", "", "Path to equivalencies file")
	uploadCmd.Flags().StringVar(&studentsPath, "students", "", "Path to students file")
	uploadCmd.Flags().StringVar(&studentCoursesPath, "studentcourses", "", "Path to student courses file")

	// Mark required flags
	if err := uploadCmd.MarkFlagRequired("apikey"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking apikey flag as required: %v\n", err)
	}

	return uploadCmd
}

// runUpload runs the upload command
func runUpload() error {
	// Check if at least one file is specified
	if coursesPath == "" && equivalenciesPath == "" && studentsPath == "" && studentCoursesPath == "" {
		return fmt.Errorf("at least one file must be specified")
	}

	// Check if files exist
	filesToCheck := map[string]string{
		"courses":        coursesPath,
		"equivalencies":  equivalenciesPath,
		"students":       studentsPath,
		"studentcourses": studentCoursesPath,
	}

	for name, path := range filesToCheck {
		if path != "" {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				return fmt.Errorf("%s file does not exist: %s", name, path)
			}
		}
	}

	// Create uploader
	u := uploader.New(Config, Logger)

	// Upload files
	Logger.Info("Uploading files to %s", Config.APIEndpoint)
	response, err := u.UploadFilesFromPaths(apiKey, coursesPath, equivalenciesPath, studentsPath, studentCoursesPath)
	if err != nil {
		return fmt.Errorf("failed to upload files: %w", err)
	}

	// Print response
	if response.Success {
		fmt.Println("Upload successful!")
		fmt.Println(response.Message)
		return nil
	} else {
		fmt.Printf("Upload failed with status code %d\n", response.Code)
		fmt.Println(response.Message)
		return fmt.Errorf("upload failed")
	}
}
