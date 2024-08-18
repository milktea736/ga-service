#!/bin/bash
SCRIPT_DIR=$(dirname "$(realpath "$0")")
ROOT_DIR="$SCRIPT_DIR/.."

parse_dirs="$ROOT_DIR/cmd/server,$ROOT_DIR/api/handlers"
swag init --dir "$parse_dirs" --output "$ROOT_DIR/docs" --parseDependency
