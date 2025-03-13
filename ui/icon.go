package ui

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
)

// AppIcon returns the application icon
func AppIcon() fyne.Resource {
	// Try to load the icon.png file from various locations
	iconPaths := []string{
		"icon.png",       // Current directory
		"../../icon.png", // Project root when running from cmd/gui
		"../icon.png",    // Project root when running from cmd
	}

	var iconData []byte
	var err error

	for _, path := range iconPaths {
		iconData, err = os.ReadFile(path)
		if err == nil {
			fmt.Printf("Successfully loaded icon from: %s\n", path)
			break
		}
	}

	if err != nil {
		fmt.Printf("Failed to load icon from any location: %v\n", err)
		// Return an empty resource if loading fails
		return fyne.NewStaticResource("icon", []byte{})
	}

	return fyne.NewStaticResource("icon.png", iconData)
}
