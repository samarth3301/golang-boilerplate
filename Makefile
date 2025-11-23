.PHONY: build run test clean docker-build docker-run dev watch help

# Show help
help:
	@echo "Available commands:"
	@echo "  build         - Build the application"
	@echo "  run           - Run the application"
	@echo "  watch         - Live reload on file changes (like nodemon)"
	@echo "  test          - Run all tests"
	@echo "  test-unit     - Run unit tests"
	@echo "  test-integration - Run integration tests"
	@echo "  clean         - Clean build artifacts"
	@echo "  docker-build  - Build Docker image"
	@echo "  docker-run    - Run Docker container"
	@echo "  dev           - Run with docker-compose for development"
	@echo "  prod          - Run production setup"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  tidy          - Tidy modules"
	@echo "  help          - Show this help message"

# Build the application
build:
	go build -o bin/main ./cmd/main.go

# Run the application
run:
	go run ./cmd/main.go

# Watch for file changes and restart (like nodemon)
watch:
	PATH=$$PATH:/Users/samarthhhh/go/bin reflex -s -r '\.go$$' -- go run ./cmd/main.go

# Run tests
test:
	go test ./...

# Run integration tests
test-integration:
	go test -tags=integration ./tests/...

# Run unit tests
test-unit:
	go test ./main/... ./pkg/...

# Clean build artifacts
clean:
	rm -rf bin/

# Build Docker image
docker-build:
	docker build -f docker/Dockerfile -t golang-boilerplate .

# Run Docker container
docker-run:
	docker run -p 8080:8080 golang-boilerplate

# Run with docker-compose for development
dev:
	docker-compose up --build

# Run production setup
prod:
	cd docker && docker-compose -f docker-compose.prod.yml up --build

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Tidy modules
tidy:
	go mod tidy