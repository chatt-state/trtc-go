# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

### Changed

### Fixed

## [0.0.1] - 2024-03-13

### Added
- Initial release of TRTC-Go with CLI and GUI interfaces
- Support for uploading Courses, Equivalencies, Students, and Student Courses files
- Custom Tennessee theme for the GUI
- Native file dialogs using zenity
- Configuration management
- Detailed logging
- Cross-platform support (Windows, macOS, Linux)
- Build scripts for easy compilation
- GitHub Actions workflows for automated testing and releases
- Pre-commit hooks to ensure code quality and test validation

### Changed
- Updated GoReleaser configuration to use changelog as release notes
- Replaced native file dialog with zenity for improved cross-platform support
- Standardized release artifact naming convention for better consistency
- Updated build process to generate properly named artifacts for all platforms
- Improved source code archive naming to follow project conventions

### Fixed 
- Corrected format string in Logger.Info call to use proper formatting
- Fixed build tag format in native dialog implementation files
- Fixed inconsistent artifact naming in release workflow
- Resolved macOS build issues by using a dedicated macOS GitHub runner
- Fixed Windows build by correcting ldflags parameter format