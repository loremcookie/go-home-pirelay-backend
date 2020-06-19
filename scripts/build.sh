#!/bin/bash

#Check if bin folder exists create it if not
if [ ! -d "./bin" ]; then
  mkdir "./bin"
fi

@env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -o ./bin/api