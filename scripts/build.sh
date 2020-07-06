#!/bin/bash

#Check if bin folder exists create it if not
if [ ! -d "./bin" ]; then
  mkdir "./bin"
fi

export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=arm && export GOARM=5
go build ./cmd/api/api.go
mv ./api ./bin/.
