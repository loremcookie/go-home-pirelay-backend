#!/bin/bash

# Check if golangci linter is installed
if [[ ! -f $(go env GOPATH)/bin/golangci-lint ]]; then
  # IF golangci linter is not installed install it
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
fi

# Print version of linter
golangci-lint version

# Make the modules nice and clean
go mod tidy

# Lint withe the golangci linter
golangci-lint run cmd/... pkg/... internal/...