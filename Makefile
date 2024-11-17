BINARY_NAME=txparser
GO=go

.PHONY: all
all: test build

.PHONY: build
build:
	@echo "Building binary..."
	$(GO) build -o $(BINARY_NAME) ./cmd/main.go

.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./...

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)

