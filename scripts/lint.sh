#!/bin/bash

# Make the modules nice and clean
go mod tidy

# Lint withe the gometalinter linter
gometalinter ./... --vendor