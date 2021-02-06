#!/bin/bash

cd backend

# run server
go build main.go
chmod +x main
./main.go

