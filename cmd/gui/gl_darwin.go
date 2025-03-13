//go:build darwin
// +build darwin

package main

import (
	"os"
)

func init() {
	// On macOS, we can use the native renderer
	// This is just a placeholder for macOS-specific initialization
	os.Setenv("FYNE_SCALE", "1.0") // Ensure proper scaling on macOS
}
