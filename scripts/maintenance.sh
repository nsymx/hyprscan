#!/usr/bin/env bash

# Make sure to run this script from project's root.

function check_license() {
	echo ":: Validating LICENSE..."

	if ! command -v "go-licenses" >/dev/null 2>&1; then
		go install github.com/google/go-licenses/v2@latest
	fi

	go-licenses check ./... \
		--disallowed_types=forbidden,restricted \
		--ignore=golang.org/x/sys && echo ":: Done."
	echo ""
}

if [[ ! -f "./go.mod" ]]; then
	echo ":: Not a Go module directory, exiting..."
	exit 1
fi

if ! command -v "go" >/dev/null 2>&1; then
	echo ":: command not found: go"
	exit 1
fi

if [[ -f "./LICENSE" ]]; then
	check_license
fi

echo ":: Updating module dependencies..."

go get -u ./...
go mod tidy

echo ":: Done."
echo ""

echo ":: Running gofmt..."
gofmt -w -e .

echo ":: Done."
