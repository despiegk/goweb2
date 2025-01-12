#!/bin/bash

# Remove generated API code
rm -f api/generated/api.gen.go
rm -f api/api.gen.go

# Remove any Go build files
rm -f *.exe
rm -rf bin/
rm -rf dist/

# Remove Go test cache
go clean -testcache

# Remove any temp files
find . -type f -name "*.tmp" -delete
find . -type f -name "*~" -delete
find . -type f -name "*.swp" -delete

echo "Cleanup complete"
