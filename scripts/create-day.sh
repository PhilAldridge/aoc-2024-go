#!/bin/bash

# ----- CONFIG -----
SOURCE_DIR="./template"   # change this to your source folder
# -------------------

# Check for argument
if [ -z "$1" ]; then
  echo "Usage: $0 <new-folder-name>"
  exit 1
fi

TARGET_DIR="./puzzles/$1"

# Create the folder if it doesn't exist
if [ -d "$TARGET_DIR" ]; then
  echo "Folder '$TARGET_DIR' already exists."
else
  mkdir -p "$TARGET_DIR"
  echo "Created folder: $TARGET_DIR"
fi

# Copy files
if [ ! -d "$SOURCE_DIR" ]; then
  echo "Error: SOURCE_DIR '$SOURCE_DIR' does not exist."
  exit 1
fi

cp -r "$SOURCE_DIR"/* "$TARGET_DIR"/
echo "Copied files from $SOURCE_DIR to $TARGET_DIR"
