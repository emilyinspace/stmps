# Variable declaration
MAIN_PACKAGE_PATH := ./
BINARY_NAME := stmps


## help: Print help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## tidy: Format code and tidy module deps
.PHONY: tidy
tidy: 
	go fmt ./...
	go mod tidy -v


## build: Build the application
.PHONY: build
build: tidy
	go build -ldflags "-s -w -X main.clientCommitHash=`git rev-parse --short HEAD`" -o=./${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: Build the application and run it without arguments
.PHONY: run
run: build
	./${BINARY_NAME}

## install: Install binary locally
.PHONY: install
install: build
	mkdir -pv ${HOME}/.local/bin
	cp -v ./${BINARY_NAME} ${HOME}/.local/bin

