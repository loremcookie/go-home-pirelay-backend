#!/bin/bash

# When binary folder exists delete it
if [ -d "./bin" ]; then
  rm -rf "./bin"
fi

# When database folder exists delete it
if [ -d "./db" ]; then
  rm -rf "./db"
fi