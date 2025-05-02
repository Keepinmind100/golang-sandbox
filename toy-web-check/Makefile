# Makefile for web-check CLI tool

APP_NAME = web-check
BIN_DIR = $(HOME)/go/bin

.PHONY: all build install clean test format run

all: build

## Build the app
build:
	go build -o $(APP_NAME) .

## Install to $(BIN_DIR)
install: build
	mv $(APP_NAME) $(BIN_DIR)/$(APP_NAME)
	@echo "✅ Installed to $(BIN_DIR)/$(APP_NAME)"

## Remove the binary
clean:
	rm -f $(APP_NAME)

## Run formatting
format:
	go fmt ./...

## Run tests
test:
	go test ./...

## Run the app with example args
run:
	go run main.go check https://google.com https://github.com
