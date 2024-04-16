PATH  := $(PWD)/bin:$(PATH)
SHELL := env PATH=$(PATH) /bin/sh

# Terraform
version  ?= "1.7.1"
os       ?= $(shell uname|tr A-Z a-z)
ifeq ($(shell uname -m),x86_64)
  arch   ?= "amd64"
endif
ifeq ($(shell uname -m),i686)
  arch   ?= "386"
endif
ifeq ($(shell uname -m),aarch64)
  arch   ?= "arm"
endif
ifeq ($(shell uname -m),arm64)
  arch   ?= "arm64"
endif

pwd       := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
terraform := $(shell command -v $(pwd)/bin/terraform 2> /dev/null)

ifndef terraform
  install ?= "true"
endif

.PHONY: install
install:
ifeq ($(install),"true")
	@curl --create-dirs --output-dir ./bin -O -L https://releases.hashicorp.com/terraform/$(version)/terraform_$(version)_$(os)_$(arch).zip
	@unzip -d $(pwd)/bin $(pwd)/bin/terraform_$(version)_$(os)_$(arch).zip && rm ./bin/terraform_$(version)_$(os)_$(arch).zip
endif
	@terraform --version

.PHONY: test clean all 

all:
	echo "Building..."

clean:
	@rm -rf $(pwd)/bin

test:
	echo "Testing..."
