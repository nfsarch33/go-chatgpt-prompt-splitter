# Note that you may need to adjust your PowerShell execution policies to allow scripts to be run.
# You can do this by opening a PowerShell prompt as Administrator and running Set-ExecutionPolicy RemoteSigned,
# then confirming the change. This needs to be done only once per machine.

param(
    [switch]$build,
    [switch]$test,
    [switch]$lint,
    [switch]$run
)

# Function for building the application
function Build {
    Write-Host "Building application"
    go mod tidy
    go build -o main.exe -v ./cmd/go-chatgpt-prompt-splitter/main.go
}

# Function for testing the application
function Test {
    Write-Host "Testing application"
    go test -v ./...
}

# Function for linting the application
function Lint {
    Write-Host "Linting application"
    go mod tidy
    golangci-lint run ./...
}

# Function for running the application
function Run {
    Write-Host "Running application"
    ./main
}

# Run tasks based on parameters

# To build: .\build.ps1 -build
if ($build) { Build }

# To test: .\build.ps1 -test
if ($test) { Test }

# To lint: .\build.ps1 -lint
if ($lint) { Lint }

# To run: .\build.ps1 -run
if ($run) { Run }

# You can also combine these to run multiple tasks at once:
# To build and run: .\build.ps1 -build -run
# To test and lint: .\build.ps1 -test -lint
# To build, test, lint, and run: .\build.ps1 -build -test -lint -run
