#!/bin/bash

# Get packages
PKGS := $(shell go list ./... | grep -v /vendor)

# Run tests
go test $(PKGS)