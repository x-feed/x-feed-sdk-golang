#!/usr/bin/env bash
#
# Generate all protobuf stubs.
# Run from repository root ./scripts/genproto.sh
#
set -e

CURRENT_DIR="$(pwd -P)"

docker run --rm -v $CURRENT_DIR/pkg:/defs namely/protoc-all:1.15 -i feed -l go -d /defs -o .
