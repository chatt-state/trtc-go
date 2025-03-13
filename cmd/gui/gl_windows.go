//go:build windows
// +build windows

package main

import (
	"os"
)

func init() {
	// On Windows, we need to set this environment variable to avoid OpenGL issues
	os.Setenv("FYNE_RENDERER", "software")
}
