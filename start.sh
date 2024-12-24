#!/bin/bash

echo "Starting GoWeb2 service..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    echo "Please install Go from https://go.dev/doc/install"
    exit 1
fi

# Generate code from OpenAPI spec
echo "Generating code from OpenAPI spec..."
oapi-codegen -config api/config.yaml api/openapi.yaml

# Run the service
echo "Starting the service..."
go mod tidy && go run main.go

# The service will be available at:
# - API: http://localhost:8080/api/v1
# - Swagger UI: http://localhost:8080/swagger/index.html
