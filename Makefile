BIN           := $(PWD)/_bin
CACHE         := $(PWD)/_cache
GOPATH        := $(CACHE)/go
PATH          := $(BIN):$(PATH)
SHELL         := env PATH=$(PATH) GOPATH=$(GOPATH) /bin/sh
PROVIDER_NAME := terraform-provider-github

os       ?= $(shell uname|tr A-Z a-z)
ifeq ($(shell uname -m),x86_64)
  arch   ?= amd64
endif
ifeq ($(shell uname -m),i686)
  arch   ?= 386
endif
ifeq ($(shell uname -m),aarch64)
  arch   ?= arm
endif
ifeq ($(shell uname -m),arm64)
  arch   ?= arm64
endif

.PHONY: all
all: format lint test build install

.PHONY: tools
tools: $(BIN)/go $(BIN)/golangci-lint

# Setup Go
go_version      := 1.22.1
go_package_name := go$(go_version).$(os)-$(arch)
go_package_url  := https://go.dev/dl/$(go_package_name).tar.gz
go_install_path := $(BIN)/go-$(go_version)-$(os)-$(arch)

$(BIN)/go:
	@mkdir -p $(BIN)
	@mkdir -p $(GOPATH)
	@echo "Downloading Go $(go_version) to $(go_install_path)..."
	@curl --silent --show-error --fail --create-dirs --output-dir $(BIN) -O -L $(go_package_url)
	@tar -C $(BIN) -xzf $(BIN)/$(go_package_name).tar.gz && rm $(BIN)/$(go_package_name).tar.gz
	@mv $(BIN)/go $(go_install_path)
	@ln -s $(go_install_path)/bin/go $(BIN)/go

# Setup golangci
golangci_version      := 1.57.1
golangci_package_name := golangci-lint-$(golangci_version)-$(os)-$(arch)
golangci_package_url  := https://github.com/golangci/golangci-lint/releases/download/v$(golangci_version)/$(golangci_package_name).tar.gz
golangci_install_path := $(BIN)/$(golangci_package_name)

$(BIN)/golangci-lint:
	@mkdir -p $(BIN)
	@echo "Downloading golangci-lint $(golangci_version) to $(BIN)/golangci-lint-$(golangci_version)..."
	@curl --silent --show-error --fail --create-dirs --output-dir $(BIN) -O -L $(golangci_package_url)
	@tar -C $(BIN) -xzf $(BIN)/$(golangci_package_name).tar.gz && rm $(BIN)/$(golangci_package_name).tar.gz
	@ln -s $(golangci_install_path)/golangci-lint $(BIN)/golangci-lint

.PHONY: update
update: $(BIN)/go
	@echo "Updating dependencies..."
	@go get -u
	@go mod tidy

.PHONY: build
build: $(BIN)/$(PROVIDER_NAME)

$(BIN)/$(PROVIDER_NAME): update
	@echo "Building..."
	@go build -o $(BIN)

.PHONY: install
install: $(CACHE)/bin/$(PROVIDER_NAME)

$(CACHE)/bin/$(PROVIDER_NAME): update
	@echo "Installing..."
	@go install ./...

.PHONY: format
format: $(BIN)/go
	@echo "Formatting..."
	@go fmt ./...

.PHONY: lint
lint: tools
	@echo "Linting..."
	@golangci-lint run ./...

.PHONY: test
test: $(BIN)/go
	@echo "Testing..."

.PHONY: clean
clean:
	@echo "Removing the $(CACHE) directory..."
	@go clean -modcache
	@rm -rf $(CACHE)
	@echo "Removing the $(BIN) directory..."
	@rm -rf $(BIN)
