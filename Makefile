.PHONY: all build test clean examples install-deps help

all: build

help: ## Display this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install-deps: ## Install ThorVG dependencies
	@echo "Installing ThorVG C library..."
	@echo "Please follow the installation instructions in README.md"

build: ## Build the Go package
	@echo "Building thorvg Go bindings..."
	go build -v ./...

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

examples: ## Build examples
	@echo "Building examples..."
	@for example in examples/*.go; do \
		if ! grep -q "+build ignore" "$$example"; then \
			echo "Building $$example..."; \
			go build -o /tmp/$$(basename $$example .go) $$example; \
		fi \
	done

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f output.png svg_output.png
	@rm -f *.test
	@rm -f *.out
	@go clean

fmt: ## Format Go code
	@echo "Formatting code..."
	go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	go vet ./...

doc: ## Generate and view documentation
	@echo "Opening documentation..."
	godoc -http=:6060 &
	@echo "Documentation available at http://localhost:6060/pkg/github.com/millken/thorvg/"
	@echo "Press Ctrl+C to stop the server"

check: fmt vet ## Run formatting and vetting checks
	@echo "All checks passed!"
