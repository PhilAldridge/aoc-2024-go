# Check for argument
if [ -z "$1" ]; then
  echo "Usage: $0 <folder-to-run>"
  exit 1
fi

go run ./puzzles/$1/main.go