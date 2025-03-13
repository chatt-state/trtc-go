@echo off
REM Setup script for TRTC-Go development environment on Windows

echo Setting up TRTC-Go development environment...

REM Check if pre-commit is installed
where pre-commit >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo pre-commit is not installed. Please install it using one of the following methods:
    echo   pip install pre-commit
    echo   choco install pre-commit
    exit /b 1
)

REM Check if golangci-lint is installed
where golangci-lint >nul 2>&1
if %ERRORLEVEL% NEQ 0 (
    echo golangci-lint is not installed. It's recommended for development.
    echo You can install it using the following method:
    echo   choco install golangci-lint
    
    set /p CONTINUE=Continue without golangci-lint? (y/n) 
    if /i "%CONTINUE%" NEQ "y" exit /b 1
)

REM Install pre-commit hooks
echo Installing pre-commit hooks...
pre-commit install

echo Development environment setup complete!
echo You can now make changes to the code. The pre-commit hooks will run automatically when you commit changes. 