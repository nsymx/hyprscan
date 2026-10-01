#!/usr/bin/env bash

bin_dir="./bin"

if ! command -v "go" >/dev/null 2>&1; then
	echo ":: command not found: go"
	exit 1
fi

if [[ -d "$bin_dir" ]]; then
	echo ":: Removing $bin_dir..."
	rm -rf "$bin_dir"
fi
