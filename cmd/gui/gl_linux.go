//go:build linux
// +build linux

package main

import (
	"os"
)

func init() {
	// On Linux with CGO enabled, we can use hardware acceleration
	// But we'll still set software rendering as a fallback for systems without proper GL support
	os.Setenv("FYNE_RENDERER", "gl")
}
