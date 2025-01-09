# Variables
APP_NAME = task-cli
GO_FILES = $(wildcard **/*.go) main.go

# Default target
all: build

# Build the executable
build: $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME)
	@echo "$(APP_NAME) built successfully!"

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@./$(APP_NAME) $(ARGS)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(APP_NAME)
	@echo "Clean up complete!"

# Format Go code
format:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Code formatted successfully!"

# Lint Go code
lint:
	@echo "Linting code..."
	@go vet ./...
	@echo "Linting complete!"

# Test the application
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete!"

# Help menu
help:
	@echo "Makefile for Task Tracker CLI"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the application"
	@echo "  run ARGS='<args>' Run the application with arguments"
	@echo "  clean        Remove build artifacts"
	@echo "  format       Format Go code"
	@echo "  lint         Lint Go code"
	@echo "  test         Run tests"
	@echo "  help         Show this help menu"
