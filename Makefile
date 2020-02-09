ENV ?= DEV

export REPORTS_DIR=./reports
# The binary to build (just the basename).
BIN := bulk-import

BASEIMAGE ?= alpine

PWD ?= $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

build: build/$(BIN)

build/$(BIN): build-dirs
	PKG=$(PKG)          \
	BIN=$(BIN)          \
	./scripts/build.sh

lint:
	./scripts/lint.sh

test:
	./scripts/test.sh

build-dirs:
	@mkdir -p build/
	@mkdir -p .go/src/$(PKG) .go/pkg .go/bin .go/std/

build: build-dirs
	@go build -o build/$(BIN) main.go


