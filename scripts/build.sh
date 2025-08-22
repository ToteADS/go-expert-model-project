#!/bin/bash

# Build script for projeto-modelo

set -e

echo "Building projeto-modelo..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Build for current platform
go build -o bin/projeto-modelo cmd/projeto-modelo/main.go

echo "Build completed successfully!"
echo "Binary location: bin/projeto-modelo"
