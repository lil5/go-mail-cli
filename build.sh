#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o bin/gomail-linux-amd64 main.go
GOOS=darwin GOARCH=amd64 go build -o bin/gomail-darwin-amd64 main.go
GOOS=windows GOARCH=amd64 go build -o bin/gomail-windows-amd64 main.go