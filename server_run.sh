#!/bin/bash

cd backend

export PATH=$PATH:/usr/local/go/bin

# run server
/usr/local/go/bin/go build main.go
chmod +x main
./main

