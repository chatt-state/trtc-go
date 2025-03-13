#!/bin/bash
# Build script for TRTC-Go GUI application

# Set version (use git tag if available, otherwise use "dev")
VERSION=$(git describe --tags 2>/dev/null || echo "dev")
echo "Building version: $VERSION"

# Ensure icon is available
cp ../../icon.png .

# Build for current platform
echo "Building for current platform..."
go build -o trtc-gui -ldflags="-s -w -X github.com/chatt-state/trtc-go/cmd/gui.Version=$VERSION"

# Cross-compile for Windows (if needed)
if [ "$1" == "windows" ]; then
    echo "Cross-compiling for Windows..."
    GOOS=windows GOARCH=amd64 go build -o trtc-gui.exe -ldflags="-s -w -X github.com/chatt-state/trtc-go/cmd/gui.Version=$VERSION"
    
    # Create a zip file for Windows distribution
    echo "Creating Windows distribution package..."
    mkdir -p dist/windows
    cp trtc-gui.exe dist/windows/
    cp icon.png dist/windows/
    cd dist/windows
    zip -r ../../trtc-gui-windows.zip .
    cd ../..
    echo "Windows package created: trtc-gui-windows.zip"
fi

# Package for macOS (if on macOS and requested)
if [ "$(uname)" == "Darwin" ] && [ "$1" == "macos" ]; then
    echo "Creating macOS application bundle..."
    mkdir -p dist/macos/TRTC-Go.app/Contents/{MacOS,Resources}
    
    # Copy executable
    cp trtc-gui dist/macos/TRTC-Go.app/Contents/MacOS/
    
    # Copy icon
    cp icon.png dist/macos/TRTC-Go.app/Contents/Resources/
    
    # Create Info.plist
    cat > dist/macos/TRTC-Go.app/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>trtc-gui</string>
    <key>CFBundleIconFile</key>
    <string>icon.png</string>
    <key>CFBundleIdentifier</key>
    <string>com.chscc.trtc-go</string>
    <key>CFBundleName</key>
    <string>TRTC-Go</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleShortVersionString</key>
    <string>${VERSION}</string>
    <key>NSHighResolutionCapable</key>
    <true/>
</dict>
</plist>
EOF
    
    # Create a DMG file (requires hdiutil, which is macOS-specific)
    cd dist/macos
    zip -r ../../trtc-gui-macos.zip .
    cd ../..
    echo "macOS package created: trtc-gui-macos.zip"
fi

echo "Build completed successfully!" 