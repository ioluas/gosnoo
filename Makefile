# gosnoo Makefile

BINARY_NAME := gosnoo
CMD_PATH := ./cmd/reddit-tui
BUILD_DIR := ./build

# Version info (can be overridden)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Linker flags
LDFLAGS_COMMON := -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)
LDFLAGS_DEBUG := $(LDFLAGS_COMMON)
LDFLAGS_RELEASE := $(LDFLAGS_COMMON) -s -w

# Build tags
BUILD_TAGS :=

# Platforms for cross-compilation
PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64 \
	windows/arm64

.PHONY: all build debug release clean test lint fmt vet run help deps multiplatform

# Default target
all: build

# Development build (same as debug)
build: debug

# Debug build - includes debug symbols, no optimization
debug:
	@echo "Building debug binary..."
	go build -tags "$(BUILD_TAGS)" -ldflags "$(LDFLAGS_DEBUG)" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Debug build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Release build - optimized, stripped
release:
	@echo "Building release binary..."
	CGO_ENABLED=0 go build -tags "$(BUILD_TAGS)" -trimpath -ldflags "$(LDFLAGS_RELEASE)" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Release build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Multiplatform builds
multiplatform:
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		GOOS=$$(echo $$platform | cut -d'/' -f1); \
		GOARCH=$$(echo $$platform | cut -d'/' -f2); \
		output=$(BUILD_DIR)/$(BINARY_NAME)-$${GOOS}-$${GOARCH}; \
		if [ "$${GOOS}" = "windows" ]; then output=$${output}.exe; fi; \
		echo "Building $${GOOS}/$${GOARCH}..."; \
		CGO_ENABLED=0 GOOS=$${GOOS} GOARCH=$${GOARCH} go build \
			-tags "$(BUILD_TAGS)" \
			-trimpath \
			-ldflags "$(LDFLAGS_RELEASE)" \
			-o $${output} $(CMD_PATH); \
	done
	@echo "Multiplatform build complete!"
	@ls -la $(BUILD_DIR)/

# Run the application
run: debug
	$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	go test -v -race ./...

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Lint code (requires golangci-lint)
lint:
	golangci-lint run ./...

# Format code
fmt:
	go fmt ./...
	gofumpt -l -w .

# Vet code
vet:
	go vet ./...

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Show help
help:
	@echo "reddit-tui Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make              Build debug binary (default)"
	@echo "  make build        Build debug binary"
	@echo "  make debug        Build with debug symbols"
	@echo "  make release      Build optimized release binary"
	@echo "  make multiplatform Build for all platforms (linux, darwin, windows)"
	@echo "  make run          Build and run the application"
	@echo "  make test         Run tests"
	@echo "  make test-coverage Run tests with coverage report"
	@echo "  make lint         Run linter (requires golangci-lint)"
	@echo "  make fmt          Format code"
	@echo "  make vet          Run go vet"
	@echo "  make deps         Download and tidy dependencies"
	@echo "  make clean        Remove build artifacts"
	@echo "  make help         Show this help"
