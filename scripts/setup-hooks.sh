#!/bin/bash
set -e

echo "Setting up pre-commit hooks..."

# Check if pre-commit is installed
if ! command -v pre-commit &> /dev/null; then
    echo "pre-commit not found, installing..."
    if command -v brew &> /dev/null; then
        brew install pre-commit
    else
        echo "Homebrew not found. Please install pre-commit manually:"
        echo "  - macOS: brew install pre-commit"
        echo "  - Linux: pip install pre-commit"
        echo "  - Windows: pip install pre-commit"
        exit 1
    fi
fi

# Install the pre-commit hooks
pre-commit install

# Create a Git hook to ensure pre-commit runs on every commit
cat > .git/hooks/pre-commit << 'EOF'
#!/bin/bash
set -e

# Run pre-commit
pre-commit run --all-files
EOF

# Make the hook executable
chmod +x .git/hooks/pre-commit

echo "Pre-commit hooks installed successfully!"
echo "They will now run automatically on every commit." 