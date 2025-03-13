#!/bin/bash
# Setup script for TRTC-Go development environment

echo "Setting up TRTC-Go development environment..."

# Check if pip is installed
if ! command -v pip &> /dev/null; then
    echo "pip is not installed. Please install Python and pip first."
    exit 1
fi

# Install pre-commit
echo "Installing pre-commit..."
pip install pre-commit

# Install pre-commit hooks
echo "Installing pre-commit hooks..."
pre-commit install

echo "Development environment setup complete!"
echo "You can now make changes to the code. The pre-commit hooks will run automatically when you commit changes." 