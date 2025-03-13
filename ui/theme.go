package ui

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// TennesseeTheme is a custom theme with Tennessee-inspired colors
type TennesseeTheme struct {
	fyne.Theme
}

// NewTennesseeTheme creates a new Tennessee-inspired theme
func NewTennesseeTheme() fyne.Theme {
	return &TennesseeTheme{Theme: theme.DefaultTheme()}
}

// Color returns the color for the specified name and theme
func (t *TennesseeTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Tennessee colors
	tennesseeBlue := color.NRGBA{R: 0, G: 40, B: 104, A: 255} // #002868
	tennesseeRed := color.NRGBA{R: 204, G: 1, B: 0, A: 255}   // #CC0100
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}      // #FFFFFF
	lighterBlue := color.NRGBA{R: 20, G: 60, B: 124, A: 255}  // #143C7C

	// Handle information dialog text color
	if name == theme.ColorNameForeground && strings.Contains(string(name), "information") {
		return white
	}

	// Handle success dialog text color
	if name == theme.ColorNameSuccess {
		return tennesseeBlue
	}

	// Handle text on success backgrounds
	if name == theme.ColorNameForegroundOnSuccess {
		return white
	}

	// Direct handling for specific color names
	if name == theme.ColorNameButton {
		return tennesseeBlue
	}

	if name == theme.ColorNameForegroundOnPrimary {
		return white
	}

	// Handle specific dialog button text
	if name == theme.ColorNameForeground {
		return tennesseeBlue
	}

	// Handle specific dialog button colors
	if string(name) == "confirmButton" {
		return tennesseeBlue
	}

	// Special handling for dialog buttons
	if string(name) == "buttonConfirm" || string(name) == "confirmButton" {
		return tennesseeBlue
	}

	// Special handling for text on buttons
	if string(name) == "buttonTextConfirm" || string(name) == "textConfirm" {
		return white
	}

	switch name {
	case theme.ColorNamePrimary:
		// Tennessee Blue for primary color
		return tennesseeBlue
	case theme.ColorNameForeground:
		// Tennessee Blue for text and icons
		return tennesseeBlue
	case theme.ColorNameButton:
		// Tennessee Blue for buttons
		return tennesseeBlue
	case theme.ColorNameBackground:
		// White for background
		return white
	case theme.ColorNameHover:
		// Lighter Blue for hover
		return lighterBlue
	case theme.ColorNameSelection:
		// Lighter Tennessee Blue for selection
		return lighterBlue
	case theme.ColorNameDisabled:
		// Light Gray for disabled elements
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255} // #C8C8C8
	case theme.ColorNamePlaceHolder:
		// Medium Gray for placeholders
		return color.NRGBA{R: 150, G: 150, B: 150, A: 255} // #969696
	case theme.ColorNameShadow:
		// Dark Gray for shadows
		return color.NRGBA{R: 100, G: 100, B: 100, A: 100} // #646464 with transparency
	case theme.ColorNameInputBackground:
		// White for input fields
		return white
	case theme.ColorNameHyperlink:
		// Tennessee Blue for hyperlinks
		return tennesseeBlue
	case theme.ColorNameInputBorder:
		// Tennessee Blue for input borders
		return tennesseeBlue
	case theme.ColorNameOverlayBackground:
		// White for dialog backgrounds
		return white
	case theme.ColorNameMenuBackground:
		// White for menu backgrounds
		return white
	case theme.ColorNameForegroundOnPrimary, theme.ColorNameForegroundOnError, theme.ColorNameForegroundOnSuccess, theme.ColorNameForegroundOnWarning:
		// White text on colored backgrounds
		return white
	case theme.ColorNamePressed:
		// Lighter blue for pressed state
		return lighterBlue
	case theme.ColorNameScrollBar:
		// Tennessee Blue for scrollbars
		return tennesseeBlue
	case theme.ColorNameError:
		// Tennessee Red for error
		return tennesseeRed
	case theme.ColorNameFocus:
		// Tennessee Blue for focus
		return tennesseeBlue
	default:
		// For any other color names, check if they might be related to button text
		if name == "buttonText" || name == "foregroundOnButton" {
			return white
		}

		// Handle text on hover and selection backgrounds
		if name == "foregroundOnHover" || name == "foregroundOnSelection" {
			return white
		}

		// Handle file dialog specific colors
		if strings.Contains(string(name), "foreground") &&
			(strings.Contains(string(name), "hover") ||
				strings.Contains(string(name), "selection") ||
				strings.Contains(string(name), "primary") ||
				strings.Contains(string(name), "button")) {
			return white
		}

		// Handle dialog button colors
		if strings.Contains(string(name), "button") ||
			strings.Contains(string(name), "dialog") ||
			strings.Contains(string(name), "confirm") ||
			strings.Contains(string(name), "open") {
			return tennesseeBlue
		}

		// Handle any color with "open" in the name
		if strings.Contains(string(name), "open") {
			return tennesseeBlue
		}

		// Handle confirm and cancel buttons
		if strings.Contains(string(name), "confirm") {
			return tennesseeBlue
		}

		if strings.Contains(string(name), "cancel") {
			return tennesseeRed
		}

		// Special handling for file dialog text
		if strings.Contains(string(name), "file") && strings.Contains(string(name), "foreground") {
			// For file dialog text that might be on blue backgrounds
			return white
		}

		// Handle any color with "action" in the name (often used for buttons)
		if strings.Contains(string(name), "action") {
			return tennesseeBlue
		}

		// Handle any color with "accent" in the name
		if strings.Contains(string(name), "accent") {
			return tennesseeBlue
		}

		return t.Theme.Color(name, variant)
	}
}

// Icon returns the icon for the specified name and theme
func (t *TennesseeTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// Use the default icons for now
	return t.Theme.Icon(name)
}

// Font returns the font for the specified text style and size
func (t *TennesseeTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.Theme.Font(style)
}

// Size returns the size for the specified element
func (t *TennesseeTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		// Slightly larger text for better readability
		return t.Theme.Size(name) * 1.1
	case theme.SizeNameHeadingText:
		// Larger heading text
		return t.Theme.Size(name) * 1.1
	case theme.SizeNameInputBorder:
		// Slightly thicker borders
		return 1.5
	default:
		return t.Theme.Size(name)
	}
}
