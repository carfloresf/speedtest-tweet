#!/bin/sh

set -o errexit
set -o nounset

if [ -z "${BIN}" ]; then
    echo "BIN must be set"
    exit 1
fi

go build -o build/${BIN} main.go
echo "Successfully built"
