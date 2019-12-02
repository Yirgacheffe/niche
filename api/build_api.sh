#!/bin/sh

# build go source code
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o niche-api
