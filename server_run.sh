#!/bin/bash

cd backend

# db setup
sudo -u postgres psql -c "CREATE ROLE xmeme LOGIN ENCRYPTED PASSWORD 'xmeme';"
sudo -u postgres psql -c "CREATE DATABASE xmeme --owner xmeme"


# run server
go build main.go
chmod +x main
./main.go

