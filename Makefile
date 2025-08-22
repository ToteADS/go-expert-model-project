.PHONY: build test clean run

# Build the application
build:
	go build -o bin/projeto-modelo cmd/projeto-modelo/main.go

# Run the application
run:
	go run cmd/projeto-modelo/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Install dependencies
deps:
	go mod tidy
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Generate documentation
docs:
	godoc -http=:6060

# Create vendor directory
vendor:
	go mod vendor
