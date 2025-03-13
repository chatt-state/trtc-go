@echo off
REM Build script for TRTC-Go GUI application on Windows

REM Set version (use git tag if available, otherwise use "dev")
FOR /F "tokens=*" %%g IN ('git describe --tags 2^>nul') do (SET VERSION=%%g)
if "%VERSION%"=="" (SET VERSION=dev)
echo Building version: %VERSION%

REM Ensure icon is available
copy ..\..\icon.png .

REM Build for Windows
echo Building for Windows...
go build -o trtc-gui.exe -ldflags="-s -w -X github.com/chatt-state/trtc-go/cmd/gui.Version=%VERSION%"

REM Create distribution package
echo Creating distribution package...
if not exist dist\windows mkdir dist\windows
copy trtc-gui.exe dist\windows\
copy icon.png dist\windows\

echo Build completed successfully! 