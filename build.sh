#!/bin/bash
mkdir -p dist/images
VERSION=`cat VERSION`
echo "Building world version v${VERSION}..."
go build -ldflags="-X 'github.com/ironarachne/world/cmd.Version=${VERSION}'" -o dist/world main.go
