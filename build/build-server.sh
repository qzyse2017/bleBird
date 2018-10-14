#!/bin/bash

GO_WORKSPACE="$GOPATH/src/github.com/qzyse2017/bleBird"
OUTPUT_DIR="$(pwd)/dist"

mkdir -p $OUTPUT_DIR && \
    cd $GO_WORKSPACE/server && \
    echo "Building server..." && \
    go get ./... && \
    go build && \
    cp server "$OUTPUT_DIR/server" && \
    exit 0