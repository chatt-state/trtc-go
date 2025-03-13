//go:build nosoftwaregl
// +build nosoftwaregl

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// This file is included only when building with the nosoftwaregl tag.
// It provides a way to build the GUI application without OpenGL dependencies
// for cross-compilation purposes.

func init() {
	// Set the environment variable to disable hardware acceleration
	// This allows the application to run without OpenGL
	fyne.SetCurrentApp(fyne.CurrentApp())
	fyne.CurrentApp().Settings().SetTheme(theme.DefaultTheme())
}
