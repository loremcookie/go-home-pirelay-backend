#!/bin/bash

# Check if script is run with sudo privileges
if [[ $EUID -ne 0 ]]; then
  # When script isn't run as root exit
   echo "This script must be run as root"
   exit 1
fi

# Run api binary
sudo ./bin/api
