#!/bin/bash
SCRIPT_DIR=$(dirname "$(realpath "$0")")
ROOT_DIR="$SCRIPT_DIR/.."

bash "$ROOT_DIR/scripts/build-docs.sh"
go run "$ROOT_DIR/cmd/server/main.go"