PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
PWD := $(shell pwd)

ifndef VERSION
	VERSION := $(shell git describe --tags --abbrev=0)
endif

COMMIT_HASH :=$(shell git rev-parse --short HEAD)
DEV_VERSION := dev-${COMMIT_HASH}

USERID := $(shell id -u $$USER)
GROUPID:= $(shell id -g $$USER)

all: install-generator install generate check

.PHONY: install
install:
	go mod tidy
	go mod vendor

# The generator versions are pinned so that "make generate" makes the
# committed gen/ code. Keep them the same as in scripts/generate.sh. Update
# them only together with a regeneration.
GOIMPORTS_VERSION := v0.53.0
XSD2GO_VERSION := v0.0.0-20261021193221-ef1732ddf011

install-generator:
	go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
	go install github.com/moov-io/xsd2go/cli/moovio_xsd2go@$(XSD2GO_VERSION)

generate:
	./scripts/generate.sh

# check-generate fails if "make generate" changes the committed gen/ code.
.PHONY: check-generate
check-generate: generate
	git diff --exit-code -- gen/
	@test -z "$$(git status --porcelain -- gen/)" || (git status --porcelain -- gen/ && exit 1)

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=40.0 ./lint-project.sh
endif
