//go:build gles
// +build gles

package main

import (
	"os"
)

// This file is included when building with the gles tag.
// It provides a way to build the GUI application without OpenGL dependencies
// for cross-compilation purposes.

func init() {
	// Force software rendering for cross-platform compatibility
	os.Setenv("FYNE_RENDERER", "software")
}
