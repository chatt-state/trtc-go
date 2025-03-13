@echo off
REM Setup script for TRTC-Go development environment on Windows

echo Setting up TRTC-Go development environment...

REM Check if pip is installed
where pip >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo pip is not installed. Please install Python and pip first.
    exit /b 1
)

REM Install pre-commit
echo Installing pre-commit...
pip install pre-commit

REM Install pre-commit hooks
echo Installing pre-commit hooks...
pre-commit install

echo Development environment setup complete!
echo You can now make changes to the code. The pre-commit hooks will run automatically when you commit changes. 