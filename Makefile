# Adapted from https://www.thapaliya.com/en/writings/well-documented-makefiles/

THREADS ?= 4 #number of threads
CLIENTS ?= 50 #number of clients per thread
REQUESTS ?= 10000 #number of requests per client
DATA_SIZE ?= 32 #Object data size
KEY_PATTERN ?= R:R #Set:Get pattern
RATIO ?= 1:10 #Set:Get ratio
PORT ?= 7379 #Port for dicedb

# Default OS and architecture
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

VERSION=$(shell bash -c 'grep -oP "DiceDBVersion string = \"\K[^\"]+" config/config.go')

.PHONY: build test build-docker run test-one

.DEFAULT_GOAL := help

##@ Helpers
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
##@ Testing

# Changing the parallel package count to 1 due to a possible race condition which causes the tests to get stuck.
# TODO: Fix the tests to run in parallel, and remove the -p=1 flag.
test: ## run the integration tests
	go test -race -count=1 -p=1 github.com/donar-0/go-workspace/... 
	cd l/iteration/ && go test -bench=. 
	cd l/arraySlices/ && go test -bench=.
	cd l/konkruncy && go test -v -bench=.


testv: ## run the integration tests
	go test -race -count=1 -p=1 -v github.com/donar-0/go-workspace/... 
	(cd l/iteration && go test -v -bench=.)
	(cd l/arraySlices && go test -v -bench=.)
	(cd l/konkruncy && go test -v -bench=.)

test-one: ## run a single integration test function by name (e.g. make test-one TEST_FUNC=TestSetGet)
	go test -v -race -count=1 --run $(TEST_FUNC) ./l/...

unittest: ## run the unit tests
	go test -race -count=1 ./internal/...

unittest-one: ## run a single unit test function by name (e.g. make unittest-one TEST_FUNC=TestSetGet)
	go test -v -race -count=1 --run $(TEST_FUNC) ./l/...

##@ Development
run: ## run dicedb with the default configuration
	go run main.go

run-docker: ## run dicedb in a Docker container
	docker run -p 7379:7379 dicedb/dicedb:latest

format: ## format the code using go fmt
	golangci-lint run --no-config  --enable wsl_v5 --fix ./l/...
	golangci-lint run --no-config  --enable wsl_v5 --fix ./assertions/...

GOLANGCI_LINT_VERSION := 2.3.0

lint: check-golangci-lint ## run golangci-lint
	golangci-lint run ./l/...

check-golangci-lint:
	@if ! command -v golangci-lint > /dev/null || ! golangci-lint version | grep -q "$(GOLANGCI_LINT_VERSION)"; then \
		echo "Required golangci-lint version $(GOLANGCI_LINT_VERSION) not found."; \
		echo "Please install golangci-lint version $(GOLANGCI_LINT_VERSION) with the following command:"; \
		echo "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.63.4"; \
		exit 1; \
	fi

dev: format lint test ## Task we need to run during development

clean: ## clean the dicedb binary
	@echo "Cleaning build artifacts..."
	rm -f dicedb
	@echo "Clean complete."
