# TRTC-Go

TRTC-Go is a modern, user-friendly tool for uploading files to the Tennessee Reverse Transfer Consortium (TRTC) API. It provides both a command-line interface (CLI) and a graphical user interface (GUI) for ease of use.

## Features

- Upload Courses, Equivalencies, Students, and Student Courses files
- Support for both CSV and Excel file formats
- Simple, intuitive graphical interface
- Powerful command-line interface for automation
- Detailed logging and error reporting
- Cross-platform support

## Installation

### Binary Installation

Download the latest release for your platform from the [Releases](https://github.com/chscc/trtc-go/releases) page.

### Building from Source

```bash
# Clone the repository
git clone https://github.com/chscc/trtc-go.git
cd trtc-go

# Build the CLI
go build -o trtc-go ./cmd/cli

# Build the GUI
go build -o trtc-go-gui ./cmd/gui
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

### GUI

Simply run the `trtc-go-gui` executable. The application will guide you through the process of uploading files.

### CLI

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

## License

This project is licensed under the MIT License - see the LICENSE file for details. 