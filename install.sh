#!/bin/env bash

if [ `whoami` != root ]; then
    echo Please run this script as root or using sudo
    exit
fi

install_path=/usr/bin/rundeck-parser
current_version_link=https://github.com/graytonio/rundeck-parser/releases/download/v0.2/rundeck-parser-v0.2-linux-amd64.tar.gz
tmp_download_location=/tmp/rundeck-parser.tar.gz

echo "Checking for existing install"

if [ -f "$install_path" ]; then
    echo "Found existing installation"
    echo "Removing existing installation"
    sudo rm $install_path
fi

echo "Downloading latest version"
rm /tmp/rundeck-parser.tar.gz
curl -LJ $current_version_link -o $tmp_download_location

echo "Installing rundeck-parser"
sudo tar -xvzf $tmp_download_location -C /usr/bin
echo "rundeck-parser installed"