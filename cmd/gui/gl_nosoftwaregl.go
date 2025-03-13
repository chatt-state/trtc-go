//go:build gles && !windows && !darwin && !linux
// +build gles,!windows,!darwin,!linux

package main

import (
	"os"
)

// This file is included only when building with the gles tag and not on a specific platform.
// It provides a way to build the GUI application without OpenGL dependencies
// for cross-compilation purposes.

func init() {
	// Force software rendering for cross-platform compatibility
	os.Setenv("FYNE_RENDERER", "software")
}
