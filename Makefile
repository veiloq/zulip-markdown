# Include .env file if it exists
-include .env

.PHONY: all test lint clean fmt test-coverage version build install deps install-tools help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=zlmd

# Build parameters
BUILD_DIR=build
VERSION=$(shell cat VERSION 2>/dev/null || echo "0.1.0")

# TARGET: all
#
# DESCRIPTION:
#   Main target to run tests ✅
#
# USAGE EXAMPLES:
#   - make all
#
# EXPLANATION:
#   This is the default target that gets executed when running make without arguments
all: test

# TARGET: build
#
# DESCRIPTION:
#   Compiles the library to verify it builds correctly ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make build
#
# EXPLANATION:
#   Compiles the library to validate the code builds without errors
build:
	$(GOBUILD) ./...

# TARGET: build-binary
#
# DESCRIPTION:
#   Compiles the CLI binary for the current platform ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make build-binary
#
build-binary:
	@echo "==> Building binary for current platform"
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

# TARGET: build-all
#
# DESCRIPTION:
#   Builds binaries for all supported platforms ✅
#
# PREREQUISITES:
#   - Go toolchain with cross-compilation support
#
# USAGE EXAMPLES:
#   - make build-all
#
build-all: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64

# TARGET: build-linux-amd64
#
# DESCRIPTION:
#   Builds binary for Linux AMD64 platform ✅
#
# PREREQUISITES:
#   - Go toolchain with cross-compilation support
#
# USAGE EXAMPLES:
#   - make build-linux-amd64
#
build-linux-amd64:
	@echo "==> Building for Linux (amd64)"
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/$(BINARY_NAME)

# TARGET: build-linux-arm64
#
# DESCRIPTION:
#   Builds binary for Linux ARM64 platform ✅
#
# PREREQUISITES:
#   - Go toolchain with cross-compilation support
#
# USAGE EXAMPLES:
#   - make build-linux-arm64
#
build-linux-arm64:
	@echo "==> Building for Linux (arm64)"
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 ./cmd/$(BINARY_NAME)

# TARGET: build-darwin-amd64
#
# DESCRIPTION:
#   Builds binary for macOS AMD64 platform ✅
#
# PREREQUISITES:
#   - Go toolchain with cross-compilation support
#
# USAGE EXAMPLES:
#   - make build-darwin-amd64
#
build-darwin-amd64:
	@echo "==> Building for macOS (amd64)"
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/$(BINARY_NAME)

# TARGET: build-darwin-arm64
#
# DESCRIPTION:
#   Builds binary for macOS ARM64 platform ✅
#
# PREREQUISITES:
#   - Go toolchain with cross-compilation support
#
# USAGE EXAMPLES:
#   - make build-darwin-arm64
#
build-darwin-arm64:
	@echo "==> Building for macOS (arm64)"
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 ./cmd/$(BINARY_NAME)

# TARGET: install
#
# DESCRIPTION:
#   Installs the library package ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make install
#
# EXPLANATION:
#   Installs the library package for use by other projects
install:
	$(GOCMD) install ./...

# TARGET: fmt
#
# DESCRIPTION:
#   Formats Go code according to standard style ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make fmt
#
fmt:
	@echo "==> Formatting code"
	$(GOCMD) fmt ./...

# TARGET: lint
#
# DESCRIPTION:
#   Runs linters to check code quality ✅
#
# PREREQUISITES:
#   - golangci-lint tool (install with go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
#
# USAGE EXAMPLES:
#   - make lint
#
lint:
	@echo "==> Running linters"
	golangci-lint run ./...

# TARGET: test
#
# DESCRIPTION:
#   Runs all tests in the project ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make test
#
test:
	@echo "==> Running tests"
	$(GOTEST) -v ./...

# TARGET: test-coverage
#
# DESCRIPTION:
#   Runs tests with coverage reporting ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make test-coverage
#
test-coverage:
	@echo "==> Running tests with coverage"
	$(GOTEST) -v -race -coverprofile=coverage.out -covermode=atomic ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# TARGET: clean
#
# DESCRIPTION:
#   Removes build artifacts and temporary files ✅
#
# USAGE EXAMPLES:
#   - make clean
#
clean:
	@echo "==> Cleaning up"
	$(GOCLEAN)
	rm -f coverage.out coverage.html
	rm -rf $(BUILD_DIR)

# TARGET: version
#
# DESCRIPTION:
#   Shows the current version of the project ✅
#
# USAGE EXAMPLES:
#   - make version
#
version:
	@echo "ZLMD version: $(VERSION)"

# TARGET: deps
#
# DESCRIPTION:
#   Downloads and tidies dependencies ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make deps
#
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# TARGET: install-tools
#
# DESCRIPTION:
#   Installs development tools required for the project ✅
#
# PREREQUISITES:
#   - Go toolchain
#
# USAGE EXAMPLES:
#   - make install-tools
#
install-tools:
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# TARGET: help
#
# DESCRIPTION:
#   Displays help information about available targets ℹ️
#
# USAGE EXAMPLES:
#   - make help
#   - make
#
help:
	@echo "Solution: veiloq/zulip-markdown"
	@echo ""
	@echo "Description: A Markdown processor for Zulip"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@echo "  all - Run tests (default)"
	@echo "  build - Compile the library to verify it builds correctly"
	@echo "  install - Install the library package"
	@echo "  fmt - Format Go code according to standard style"
	@echo "  lint - Run linters to check code quality"
	@echo "  test - Run all tests in the project"
	@echo "  test-coverage - Run tests with coverage reporting"
	@echo "  clean - Remove build artifacts and temporary files"
	@echo "  version - Show current version"
	@echo "  deps - Download and tidy dependencies"
	@echo "  install-tools - Install development tools"
	@echo "  help - Display this help information"
	@echo ""
	@echo "Status Indicators:"
	@echo "  ✅ - Success/OK      - Operation completed successfully"
	@echo "  ❌ - Error/Failure   - Operation failed or critical issue detected"
	@echo "  ⚠️ - Warning         - Potential issue that requires attention" 