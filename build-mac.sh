#!/bin/bash
rm -r mnemonicTool-mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mnemonicTool-mac main.go
