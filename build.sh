#!/bin/bash
mkdir -p dist/images
go build -o dist/worldapi cmd/worldapi/*.go
