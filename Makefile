# Define variables
BINARY_NAME=server
GO_FILES=$(shell dir /s /b *.go)

# Build the executable
build:
	go build -o $(BINARY_NAME) cmd\main.go

# Run the server (build + run)
run: build
	./$(BINARY_NAME)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Install Go dependencies (if using Go modules)
install:
	go mod tidy

# Start the server (build + run)
start: build
	./$(BINARY_NAME)

# intall air (hot reload)
install-air:
	go install github.com/cosmtrek/air@latest

# Start the server with hot reload (install air first)
hot:
	air
