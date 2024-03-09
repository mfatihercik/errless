# Makefile for the Go project

# Go parameters
GOCMD=$(GOROOT)/bin/go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=mybinary

# Linters
GOLANGCI_LINT_CMD=$(GOPATH)/bin/golangci-lint
GOLANGCI_LINT_EXISTS := $(shell which ${GOLANGCI_LINT_CMD})
GOLANGCI_LINT_VERSION="v1.51.2" # current latest version
GOLANGCI_LINT_CONFIG=".github/.golangci-linter.yml"


all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./... -tags=test

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./
	./$(BINARY_NAME)

deps:
	$(GOCMD) get -v ./...

dev-deps-golangci-lint: ## Install `golangci-lint`
	@echo "$(GREEN_COLOR)Installing $(YELLOW_COLOR)golangci-lint$(GREEN_COLOR) to $(RED_COLOR)$(GOPATH)/bin$(NO_COLOR)"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin ${GOLANGCI_LINT_VERSION}

### Linter
lint: ## Check the code standard (lint)
	@echo "$(GREEN_COLOR)Checking the code standard$(NO_COLOR)"
	$(GOLANGCI_LINT_CMD) -c $(GOLANGCI_LINT_CONFIG) run ./...
