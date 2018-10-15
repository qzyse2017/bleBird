#!/bin/bash

OUTPUT_DIR="$(pwd)/dist"

mkdir -p $OUTPUT_DIR && \
    cd webb-app && \
    echo "Installing app depencies, wait please..." && \
    npm install && \
    echo "Building app..." && \
    npm run build && \
    cp -R dist "$OUTPUT_DIR/static"
    exit 0