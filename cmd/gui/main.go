package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"

	"github.com/chatt-state/trtc-go/internal/config"
	"github.com/chatt-state/trtc-go/pkg/logger"
	"github.com/chatt-state/trtc-go/ui"
)

var (
	// Version is the version of the application
	Version = "0.1.0"

	// Config is the application configuration
	Config *config.Config

	// Logger is the application logger
	Logger *logger.Logger
)

func main() {
	// Force the use of native system dialogs
	os.Setenv("FYNE_DIALOG_DISABLE", "0")

	// Create application with a unique ID
	a := app.NewWithID("com.chscc.trtc-go")
	a.SetIcon(ui.AppIcon())

	// Set the custom Tennessee theme
	a.Settings().SetTheme(ui.NewTennesseeTheme())

	w := a.NewWindow("TRTC File Uploader")
	w.Resize(fyne.NewSize(600, 400))

	// Load configuration
	var err error
	Config, err = config.LoadConfig()
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to load configuration: %w", err), w)
		return
	}

	// Create logger
	logFilePath := Config.LogFile
	if !filepath.IsAbs(logFilePath) {
		// If log file path is not absolute, make it relative to the current directory
		logFilePath, err = filepath.Abs(logFilePath)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to get absolute path for log file: %w", err), w)
			return
		}
	}

	Logger, err = logger.New(logFilePath, logger.LevelInfo)
	if err != nil {
		dialog.ShowError(fmt.Errorf("failed to create logger: %w", err), w)
		return
	}
	defer Logger.Close()

	// Create main content
	content := container.NewVBox(
		widget.NewLabelWithStyle("TRTC File Uploader", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Upload files to the Tennessee Reverse Transfer Consortium (TRTC) API"),
		container.NewHBox(
			widget.NewLabel("Version:"),
			widget.NewLabel(Version),
		),
		container.NewHBox(
			widget.NewLabel("API Endpoint:"),
			widget.NewLabel(Config.APIEndpoint),
		),
		widget.NewSeparator(),
		createMainContent(w),
	)

	// Set window content
	w.SetContent(content)

	// Show and run
	w.ShowAndRun()
}

// createMainContent creates the main content of the application
func createMainContent(w fyne.Window) fyne.CanvasObject {
	// Create file selection widgets
	coursesCheck := widget.NewCheck("Courses", nil)
	coursesPath := widget.NewEntry()
	coursesPath.Disable()
	coursesButton := widget.NewButtonWithIcon("Browse...", theme.DocumentIcon(), func() {
		selectFile(w, coursesPath, coursesCheck)
	})
	coursesButton.Importance = widget.HighImportance

	equivalenciesCheck := widget.NewCheck("Equivalencies", nil)
	equivalenciesPath := widget.NewEntry()
	equivalenciesPath.Disable()
	equivalenciesButton := widget.NewButtonWithIcon("Browse...", theme.DocumentIcon(), func() {
		selectFile(w, equivalenciesPath, equivalenciesCheck)
	})
	equivalenciesButton.Importance = widget.HighImportance

	studentsCheck := widget.NewCheck("Students", nil)
	studentsPath := widget.NewEntry()
	studentsPath.Disable()
	studentsButton := widget.NewButtonWithIcon("Browse...", theme.DocumentIcon(), func() {
		selectFile(w, studentsPath, studentsCheck)
	})
	studentsButton.Importance = widget.HighImportance

	studentCoursesCheck := widget.NewCheck("Student Courses", nil)
	studentCoursesPath := widget.NewEntry()
	studentCoursesPath.Disable()
	studentCoursesButton := widget.NewButtonWithIcon("Browse...", theme.DocumentIcon(), func() {
		selectFile(w, studentCoursesPath, studentCoursesCheck)
	})
	studentCoursesButton.Importance = widget.HighImportance

	// Create API key entry
	apiKeyLabel := widget.NewLabel("API Key:")
	apiKeyEntry := widget.NewPasswordEntry()
	apiKeyEntry.SetPlaceHolder("Enter your API key")

	// Create status label
	statusLabel := widget.NewLabelWithStyle("Ready", fyne.TextAlignCenter, fyne.TextStyle{})

	// Create buttons
	settingsButton := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {
		showSettingsDialog(w)
	})
	settingsButton.Importance = widget.HighImportance

	uploadButton := widget.NewButtonWithIcon("Upload Files", theme.UploadIcon(), func() {
		// Validate input
		if apiKeyEntry.Text == "" {
			dialog.ShowError(fmt.Errorf("API key is required"), w)
			return
		}

		// Check if at least one file is selected
		if !coursesCheck.Checked && !equivalenciesCheck.Checked && !studentsCheck.Checked && !studentCoursesCheck.Checked {
			dialog.ShowError(fmt.Errorf("at least one file must be selected"), w)
			return
		}

		// Perform upload
		go performUpload(
			w,
			statusLabel,
			apiKeyEntry.Text,
			coursesCheck.Checked, coursesPath.Text,
			equivalenciesCheck.Checked, equivalenciesPath.Text,
			studentsCheck.Checked, studentsPath.Text,
			studentCoursesCheck.Checked, studentCoursesPath.Text,
		)
	})
	uploadButton.Importance = widget.HighImportance

	// Create layout with more padding and spacing for a traditional desktop look
	fileSelectionContainer := container.NewVBox(
		container.NewGridWithColumns(3,
			coursesCheck,
			coursesPath,
			coursesButton,
		),
		container.NewGridWithColumns(3,
			equivalenciesCheck,
			equivalenciesPath,
			equivalenciesButton,
		),
		container.NewGridWithColumns(3,
			studentsCheck,
			studentsPath,
			studentsButton,
		),
		container.NewGridWithColumns(3,
			studentCoursesCheck,
			studentCoursesPath,
			studentCoursesButton,
		),
	)

	apiKeyContainer := container.NewVBox(
		container.NewGridWithColumns(2,
			apiKeyLabel,
			apiKeyEntry,
		),
	)

	buttonContainer := container.NewHBox(
		settingsButton,
		widget.NewSeparator(),
		uploadButton,
	)

	// Add padding around each section
	return container.NewVBox(
		container.NewPadded(fileSelectionContainer),
		widget.NewSeparator(),
		container.NewPadded(apiKeyContainer),
		widget.NewSeparator(),
		container.NewPadded(buttonContainer),
		statusLabel,
	)
}

// selectFile shows a file dialog to select a file
func selectFile(w fyne.Window, entry *widget.Entry, check *widget.Check) {
	// Use zenity with proper configuration
	path, err := zenity.SelectFile(
		zenity.Title("Select File"),
		zenity.FileFilters{
			{Name: "Spreadsheet Files", Patterns: []string{"*.csv", "*.xlsx", "*.xls"}},
			{Name: "All Files", Patterns: []string{"*"}},
		},
		zenity.Modal(),
	)

	// Check if the error is a cancellation error
	if err == zenity.ErrCanceled {
		// User canceled the dialog, do nothing
		Logger.Info("Zenity file dialog was cancelled by user")
		return
	}

	// Handle other errors with zenity
	if err != nil {
		Logger.Info("Zenity file dialog failed with error: %v", err)

		// Fall back to Fyne's dialog only on actual errors
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if uri == nil {
				return
			}
			entry.SetText(uri.URI().Path())
			check.SetChecked(true)
		}, w)
		return
	}

	// Handle successful file selection
	if path != "" {
		entry.SetText(path)
		check.SetChecked(true)
	}
}

// showSettingsDialog shows the settings dialog
func showSettingsDialog(w fyne.Window) {
	// Create form items
	endpointEntry := widget.NewEntry()
	endpointEntry.SetText(Config.APIEndpoint)

	logFileEntry := widget.NewEntry()
	logFileEntry.SetText(Config.LogFile)

	ignoreCertErrorCheck := widget.NewCheck("Ignore Certificate Errors", nil)
	ignoreCertErrorCheck.SetChecked(Config.IgnoreCertError)

	// Create form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "API Endpoint", Widget: endpointEntry},
			{Text: "Log File", Widget: logFileEntry},
			{Text: "", Widget: ignoreCertErrorCheck},
		},
		OnSubmit: func() {
			// Update configuration
			Config.APIEndpoint = endpointEntry.Text
			Config.LogFile = logFileEntry.Text
			Config.IgnoreCertError = ignoreCertErrorCheck.Checked

			// Save configuration
			if err := config.SaveConfig(Config); err != nil {
				dialog.ShowError(fmt.Errorf("failed to save configuration: %w", err), w)
				return
			}

			dialog.ShowInformation("Settings", "Settings saved successfully", w)
		},
		OnCancel: func() {
			// Do nothing
		},
	}

	// Create dialog
	d := dialog.NewCustom("Settings", "Cancel", form, w)

	// Customize dialog buttons
	saveButton := widget.NewButton("Save", func() {
		form.OnSubmit()
		d.Hide()
	})
	saveButton.Importance = widget.HighImportance

	cancelButton := widget.NewButton("Cancel", func() {
		d.Hide()
	})
	cancelButton.Importance = widget.DangerImportance

	d.SetButtons([]fyne.CanvasObject{
		cancelButton,
		saveButton,
	})

	d.Show()
}

// performUpload performs the upload operation
func performUpload(
	w fyne.Window,
	statusLabel *widget.Label,
	apiKey string,
	coursesChecked bool, coursesPath string,
	equivalenciesChecked bool, equivalenciesPath string,
	studentsChecked bool, studentsPath string,
	studentCoursesChecked bool, studentCoursesPath string,
) {
	// Update status
	statusLabel.SetText("Uploading...")

	// Create uploader
	u := ui.NewUploader(Config, Logger)

	// Prepare file paths
	var coursesFile, equivalenciesFile, studentsFile, studentCoursesFile string
	if coursesChecked {
		coursesFile = coursesPath
	}
	if equivalenciesChecked {
		equivalenciesFile = equivalenciesPath
	}
	if studentsChecked {
		studentsFile = studentsPath
	}
	if studentCoursesChecked {
		studentCoursesFile = studentCoursesPath
	}

	// Upload files
	response, err := u.UploadFiles(apiKey, coursesFile, equivalenciesFile, studentsFile, studentCoursesFile)
	if err != nil {
		statusLabel.SetText("Error: " + err.Error())
		dialog.ShowError(err, w)
		return
	}

	// Update status
	if response.Success {
		statusLabel.SetText("Upload successful!")
		dialog.ShowInformation("Success", "Files uploaded successfully", w)
	} else {
		statusLabel.SetText(fmt.Sprintf("Upload failed with status code %d", response.Code))
		dialog.ShowError(fmt.Errorf("upload failed: %s", response.Message), w)
	}
}
