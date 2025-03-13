@echo off
REM Build script for TRTC-Go GUI application on Windows

REM Ensure icon is available
copy ..\..\icon.png .

REM Build for Windows
echo Building for Windows...
go build -o trtc-gui.exe

REM Create distribution package
echo Creating distribution package...
if not exist dist\windows mkdir dist\windows
copy trtc-gui.exe dist\windows\
copy icon.png dist\windows\

echo Build completed successfully! 