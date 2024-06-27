IMAGE_NAME=parser
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: docs
docs: ## Run go generate to generate API reference documentation.
	go generate ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test -v -race ./... -coverprofile cover.out

##@ Build

.PHONY: build
build: fmt vet ## Build manager binary.
	go build -o bin/${IMAGE_NAME} application/${IMAGE_NAME}.go

.PHONY: docker-build
docker-build: ## Build docker image with the manager.
	docker build -t ${IMAGE_NAME} .

.PHONY: compose-build
compose-build: ## Build an application with docker-compose.
	docker-compose build

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMAGE_NAME}

##@ Run

.PHONY: compose-run
compose-run: compose-build ## Run docker-compose with application stack.
	docker-compose up -d

.PHONY: compose-down
compose-down: ## Shutdown an application.
	docker-compose down

# go-get-tool will 'go install' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ; \
GOBIN=$(PROJECT_DIR)/bin go install $(2) ; \
}
endef
