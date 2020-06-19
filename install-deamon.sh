#!/bin/bash

# Check if script is run with sudo privileges
if [[ $EUID -ne 0 ]]; then
  # When script isn't run as root exit
   echo "This script must be run as root"
   exit 1
fi

# Add user for backend to run on
sudo useradd go-home-pirelay-backend -s /sbin/nologin -M

# Copy systemd service file to systemd config folder to be used
cp -R ./init/api/go-home-api.service /lib/systemd/system/.

# Give config file permissions
sudo chmod 755 /lib/systemd/system/go-home-api.service

# Make directory's to store config in
mkdir /etc/go-home-pirelay-backend
mkdir /etc/go-home-pirelay-backend/api

# Copy config file to ect/ directory
cp ./config/api/API_CONFIG.env /etc/go-home-pirelay-backend/api/.

# Enable service
sudo systemctl enable go-home-api.service

# Start service
sudo systemctl start go-home-api
