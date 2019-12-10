WORKDIR = $(shell pwd)

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o build/worldapi cmd/worldapi/*.go

.PHONY: run
## run: run the application
run:
	@WORLDAPI_DATA_PATH=$(WORKDIR) \
	WORLDAPI_SAVE_TARGET=filesystem \
	WORLDAPI_SAVE_DIRECTORY=$(WORKDIR)/build \
	WORLDAPI_WEB_DOMAIN=ironarachne.test \
	build/worldapi

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning..."
	@rm -rf build/

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

.PHONY: help
## help: Prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'