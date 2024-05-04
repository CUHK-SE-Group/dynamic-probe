#!/bin/bash

# Description:
# This script installs the required packages and runs a script to get the vmlinux header.

# Exit immediately if a command exits with a non-zero status
set -e

# Update package list
echo "Updating package list..."
sudo apt update

# Install jq and wget if not already installed
echo "Installing required packages: jq, wget..."
sudo apt install -y jq wget curl llvm clang

# Change to the 'headers' directory (create it if it doesn't exist)
HEADERS_DIR="headers"
if [ ! -d "$HEADERS_DIR" ]; then
  echo "Creating directory: $HEADERS_DIR"
  mkdir "$HEADERS_DIR"
fi

echo "Navigating to directory: $HEADERS_DIR"
cd "$HEADERS_DIR"

# Execute the script to get the vmlinux header
SCRIPT_NAME="get_vmlinux.h.sh"
if [ ! -f "$SCRIPT_NAME" ]; then
  echo "Script '$SCRIPT_NAME' not found in the current directory."
  exit 1
fi

echo "Running script: $SCRIPT_NAME"
bash "$SCRIPT_NAME"

# Execute the script to get the vmlinux header
SCRIPT_NAME="update.sh"
if [ ! -f "$SCRIPT_NAME" ]; then
  echo "Script '$SCRIPT_NAME' not found in the current directory."
  exit 1
fi

echo "Running script: $SCRIPT_NAME"
bash "$SCRIPT_NAME"

echo "Script executed successfully."
