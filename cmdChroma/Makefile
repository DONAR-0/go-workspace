# Variables 
BINARY_NAME=chroma
MAIN_PATH=./cli
BUILD_DIR=./dist
GOARCH=amd64
GOOS=linux

#Default target
.PHONY: all
all: clean build

#Build the CLI binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for specific OS/Architecture
.PHONY: build-linux
build-linux:
	@echo "Building $(BINARY_NAME) for Linux"
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	@echo "✅ Linux build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64"
	
.PHONY: build-darwin
build-darwin:
	@echo "Building $(BINARY_NAME) for macOS"
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	@echo "✅ macOS build complete: $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64"

.PHONY: build-windows
build-windows:
	@echo "Building $(BINARY_NAME) for Windows"
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	@echo "✅ Windows build complete: $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe"

.PHONY: build-all
build-all: build-darwin build-windows build-linux

.PHONY: deps
deps:
	@echo "Installing dependencies"
	go mod download
	go mod tidy
	@echo "✅ Dependencies installed"

.PHONY: test
test:
	@echo "Running tests"
	go test -v ./...
	@echo "✅ Tests complete"
 
# Run tests with coverage 
test-coverage:
	@echo "Running tests"
	go test -v  -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Tests complete"

# Format Code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted"

# Lint Code
.PHONY: lint
lint:
	@echo "Linting code...."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
		go vet ./...; \
	fi 
	@echo "✅ Linting complete"
 
# Install binary to system
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME) to system"
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin
	@echo "✅ $(BINARY_NAME) installed to /usr/local/bin"

#Uninstall binary from system
.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from system..."
	sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "✅ $(BINARY_NAME) uninstalled"

# Run the cli.
.PHONY: Run
run: build
	@echo "Running in developement mode..."
	go run $(MAIN_PATH)

# Show Help
.PHONY: help
help:
	@echo "Available Targets"
	@echo "	all 		- Clean and build the CLI"
	@echo "	build 		- Build the CLI binary"
	@echo "	build-linux	- Build the CLI Linux binary"
