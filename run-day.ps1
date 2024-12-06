param (
    [string]$subFolderName
)

# Get the current directory path
$currentDir = Get-Location
$puzzlesDir = Join-Path -Path $currentDir -ChildPath "puzzles"
# Combine the current directory with the subfolder name
$fullFolderPath = Join-Path -Path $puzzlesDir -ChildPath $subFolderName

# Check if the folder path is valid
if (-not (Test-Path $fullFolderPath)) {
    Write-Host "The folder path does not exist: $fullFolderPath"
    exit
}

# Set the directory to the provided folder path
Set-Location -Path $fullFolderPath

# Check if the main.go file exists in the folder
if (-Not (Test-Path "main.go")) {
    Write-Host "main.go file not found in the folder: $fullFolderPath"
    Set-Location -Path $currentDir
    exit
}

# Run the main.go file using the 'go run' command
Write-Host "Running main.go in folder: $fullFolderPath"
go run main.go

Set-Location -Path $currentDir