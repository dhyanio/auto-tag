# Variables
BINARY_NAME := autotag
BINARY_DIR := bin
BINARY_PATH := $(BINARY_DIR)/$(BINARY_NAME)
GO_FILES := $(shell find . -type f -name '*.go')

# Targets
.PHONY: all build clean start test lint fmt

all: build

build: $(GO_FILES)
	@mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_PATH)

start: build
	$(BINARY_PATH) start -c $(or $(CONFIG), config.yaml)

test:
	@go test -v ./...

lint:
	@golangci-lint run ./...

fmt:
	@go fmt ./...

clean:
	@rm -rf $(BINARY_DIR)