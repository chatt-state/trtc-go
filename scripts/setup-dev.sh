#!/bin/bash
# Setup script for TRTC-Go development environment

echo "Setting up TRTC-Go development environment..."

# Check if pre-commit is installed
if ! command -v pre-commit &> /dev/null; then
    echo "pre-commit is not installed. Please install it using one of the following methods:"
    echo "  pip install pre-commit"
    echo "  brew install pre-commit (macOS)"
    exit 1
fi

# Check if golangci-lint is installed
if ! command -v golangci-lint &> /dev/null; then
    echo "golangci-lint is not installed. It's recommended for development."
    echo "You can install it using one of the following methods:"
    echo "  brew install golangci-lint (macOS)"
    echo "  choco install golangci-lint (Windows)"
    echo "  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin (Linux)"
    
    read -p "Continue without golangci-lint? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Install pre-commit hooks
echo "Installing pre-commit hooks..."
pre-commit install

echo "Development environment setup complete!"
echo "You can now make changes to the code. The pre-commit hooks will run automatically when you commit changes." 