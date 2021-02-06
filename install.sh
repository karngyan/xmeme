#!/bin/bash

# dev installation

wget https://dl.google.com/go/go1.15.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.15.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin


# database installation

sudo apt update
sudo apt install -y postgresql

