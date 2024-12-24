#!/bin/bash

echo "Installing GoWeb2 dependencies..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    echo "Please install Go from https://go.dev/doc/install"
    exit 1
fi

# Initialize go module if not already initialized
if [ ! -f "go.mod" ]; then
    go mod init github.com/despiegk/goweb2
fi

# Install required packages
go get -u github.com/labstack/echo/v4
go get -u github.com/deepmap/oapi-codegen/v2/pkg/middleware
go get -u github.com/getkin/kin-openapi/openapi3

# Install oapi-codegen CLI
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

# Create necessary directories
mkdir -p api/v1
mkdir -p internal/model
mkdir -p internal/service
mkdir -p docs
mkdir -p pkg

echo "Installation completed successfully!"
echo "You can now run ./start.sh to start the service"
