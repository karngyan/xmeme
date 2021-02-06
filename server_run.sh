#!/bin/bash

cd backend

export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go

# run server
$GOROOT/bin/go build main.go
chmod +x main
./main

