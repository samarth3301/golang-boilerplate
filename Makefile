.PHONY: build run test clean docker-build docker-run dev

# Build the application
build:
	go build -o bin/main ./cmd/main.go

# Run the application
run:
	go run ./cmd/main.go

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