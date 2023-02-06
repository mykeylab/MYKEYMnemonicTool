#!/bin/bash
rm -r mnemonicTool-linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mnemonicTool-linux main.go
