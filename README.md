# TRTC-Go

TRTC-Go is a modern, user-friendly tool for uploading files to the Tennessee Reverse Transfer (TRTC) API. It provides both a command-line interface (CLI) and a graphical user interface (GUI) for ease of use.

## Table of Contents
- [Installation Guide](#installation-guide)
  - [Windows Installation](#windows-installation)
  - [macOS Installation](#macos-installation)
  - [Linux Installation](#linux-installation)
- [Features](#features)
- [Building from Source](#building-from-source)
- [Usage](#usage)
  - [GUI Usage](#gui-usage)
  - [CLI Usage](#cli-usage)
- [Configuration](#configuration)
- [Development Setup](#development-setup)
- [License](#license)

## Installation Guide

### Windows Installation

1. Download the latest Windows release from the [Releases](https://github.com/chatt-state/trtc-go/releases) page.
   - Look for files named `trtc-go_Windows_x86_64.zip` (for 64-bit Windows) or `trtc-go_Windows_i386.zip` (for 32-bit Windows)
2. Extract the ZIP file to a location of your choice.
3. Double-click on `trtc-go-gui.exe` to run the application.
4. For first-time setup, you may need to configure your API key in the application settings.

### macOS Installation

1. Download the latest macOS release from the [Releases](https://github.com/chatt-state/trtc-go/releases) page.
   - Look for files named `trtc-go_Darwin_x86_64.dmg` or `trtc-go_Darwin_arm64.dmg` (for Apple Silicon Macs)
   - Note: "Darwin" is the name of the macOS operating system core
2. Open Terminal and navigate to your Downloads folder:
   ```bash
   cd ~/Downloads
   ```
3. Remove the quarantine attribute to allow the app to open:
   ```bash
   xattr -c trtc-go-gui.app
   ```
4. Move the application to your Applications folder:
   ```bash
   mv trtc-go-gui.app /Applications/
   ```
5. Open the app from your Applications folder.
6. For first-time setup, you may need to configure your API key in the application settings.

### Linux Installation

1. Download the latest Linux release from the [Releases](https://github.com/chatt-state/trtc-go/releases) page.
   - Look for files named `trtc-go_Linux_x86_64.tar.gz` (for 64-bit Linux) or `trtc-go_Linux_i386.tar.gz` (for 32-bit Linux)
   - For ARM-based systems like Raspberry Pi, look for `trtc-go_Linux_armv6.tar.gz` or `trtc-go_Linux_arm64.tar.gz`
2. Extract the archive file:
   ```bash
   tar -xzf trtc-go_Linux_x86_64.tar.gz
   ```
3. Make the file executable:
   ```bash
   chmod +x trtc-go-gui
   ```
4. Run the application:
   ```bash
   ./trtc-go-gui
   ```

## Features

- Upload Courses, Equivalencies, Students, and Student Courses files
- Support for both CSV and Excel file formats
- Simple, intuitive graphical interface
- Powerful command-line interface for automation
- Detailed logging and error reporting
- Cross-platform support (Windows, macOS, Linux)

## Building from Source

```bash
# Clone the repository
git clone https://github.com/chatt-state/trtc-go.git
cd trtc-go

# Build the CLI
go build -o trtc-go ./cmd/cli

# Build the GUI (macOS with full features)
go build -o trtc-go-gui ./cmd/gui

# Build the GUI (Windows/Linux or cross-platform)
go build -tags=gles -o trtc-go-gui ./cmd/gui
```

### Cross-Platform Compatibility

This project uses [Fyne](https://fyne.io/) for its GUI, which has specific requirements for cross-compilation due to its CGO dependencies. For cross-platform compatibility, we use the `gles` build tag which enables software rendering.

For Windows and Linux builds, we use the `-tags=gles` flag to ensure compatibility:

```bash
# Build for Windows/Linux with software rendering
go build -tags=gles -o trtc-go-gui ./cmd/gui
```

### Building with Build Scripts

For convenience, build scripts are provided for both macOS/Linux and Windows:

**macOS/Linux:**
```bash
# Build for current platform
cd cmd/gui
./build.sh

# Build for macOS with app bundle
./build.sh macos

# Cross-compile for Windows
./build.sh windows
```

**Windows:**
```batch
cd cmd\gui
build.bat
```

### Releasing with GoReleaser

This project uses [GoReleaser](https://goreleaser.com/) to automate the release process. To create a new release:

1. Tag the commit you want to release:
   ```bash
   git tag -a v0.1.0 -m "First release"
   git push origin v0.1.0
   ```

2. GitHub Actions will automatically build and release the binaries for all platforms.

## Usage

### GUI Usage

The graphical interface provides an intuitive way to upload files to the TRTC system.

1. Launch the application by double-clicking the executable (Windows) or opening the app (macOS).
2. Configure your API key if this is your first time using the application.
3. Use the file selection buttons to choose your data files for upload.
4. Click the "Upload" button to begin the upload process.
5. View the logs panel for detailed information about the upload process.

### CLI Usage

```bash
# Upload a courses file
trtc-go upload -apikey="your-api-key" -courses="path/to/courses.csv"

# Upload multiple file types
trtc-go upload -apikey="your-api-key" -courses="path/to/courses.csv" -equivalencies="path/to/equivalencies.csv"

# Configure settings
trtc-go config set -endpoint="https://api.example.com"

# Get help
trtc-go help
```

## Configuration

TRTC-Go stores its configuration in a file located at:

- Windows: `%APPDATA%\trtc-go\config.yaml`
- macOS: `$HOME/Library/Application Support/trtc-go/config.yaml`
- Linux: `$HOME/.config/trtc-go/config.yaml`

You can edit this file directly or use the configuration commands in the CLI.

### API Endpoint

The default API endpoint is set to `https://rts.tnreversetransfer.org/api/Upload`. If you need to change it, you can use the following command:

```bash
trtc-go config set --endpoint="https://your-api-endpoint.com"
```

## Development Setup

This project uses pre-commit hooks to ensure code quality and that tests pass before commits. To set up the pre-commit hooks:

#### Using Setup Scripts

We provide setup scripts to make it easy to install the pre-commit hooks:

**macOS/Linux:**
```bash
./scripts/setup-dev.sh
```

**Windows:**
```batch
scripts\setup-dev.bat
```

#### Manual Setup

1. Install pre-commit:
   ```bash
   # Using pip
   pip install pre-commit
   
   # Or using Homebrew on macOS
   brew install pre-commit
   ```

2. Install the git hooks:
   ```bash
   pre-commit install
   ```

3. (Optional) Install golangci-lint:
   ```bash
   # macOS
   brew install golangci-lint
   
   # Windows
   choco install golangci-lint
   
   # Linux
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
   ```

Now, every time you commit, the pre-commit hooks will run to ensure tests pass and code quality is maintained.

## License

This project is licensed under the MIT License - see the LICENSE file for details.