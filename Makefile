APPNAME := speedtest-tweet
VERSION := 1.0

# Information needed for test and linting reports generation
export REPORTS_DIR=./reports

.DEFAULT_GOAL := build

.PHONY: build
build:
	mkdir -p build
	GOOS=$(GOOS) GOARCH=$(GOARCH) APPNAME=$(APPNAME) ./scripts/build.sh

lint:
	./scripts/lint.sh

test:
	./scripts/test.sh


