.PHONY: help build run migrate dev clean

# Default target
help:
	@echo "Available commands:"
	@echo "  migrate  - Run database migrations"
	@echo "  run      - Run the API server"
	@echo "  dev      - Run in development mode with auto-reload"
	@echo "  build    - Build the application"
	@echo "  clean    - Clean build artifacts"
	@echo "  tidy     - Download and tidy dependencies"

# Install dependencies
tidy:
	go mod tidy

# Run database migrations
migrate:
	@echo "Running database migrations..."
	go run cmd/migrate/main.go

# Build the application
build:
	@echo "Building application..."
	go build -o bin/server main.go

# Run the API server
run: migrate
	@echo "Starting API server..."
	go run main.go

# Development mode with auto-reload (requires air)
dev:
	@echo "Running in development mode..."
	@if ! command -v air &> /dev/null; then \
		echo "Installing air for hot reload..."; \
		go install github.com/cosmtrek/air@latest; \
	fi
	air -c .air.toml

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f bin/server
	rm -f ecommerce.db

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest