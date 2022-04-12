GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=tasker

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

all: help

## Build:
build: ## Build project and put the output binary in bin/
	$(GOCMD) build -v -o bin/$(BINARY_NAME)

build-linux: ## Build project for linux and put the output binary in bin/
	$(GOCMD) mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOCMD) build -v -o bin/$(BINARY_NAME)

## Run:
run: ## Build and run project
	$(GOCMD) build -v -o bin/$(BINARY_NAME) 
	bin/./$(BINARY_NAME)

## Clean:
clean: ## Remove build related file
	rm bin/tasker

## Vendor:
vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

## Test:
test: ## Run the tests of the project
	$(GOCMD) test -v ./...

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)