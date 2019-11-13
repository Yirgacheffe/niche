#!/bin/sh

# Build go source code
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o greetings
