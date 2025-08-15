#!/bin/bash

# Exit on error
set -e

# Get all packages in the module
packages=$(go list ./...)

# Loop over each package and compute coverage
for pkg in $packages; do
    # Run test with coverage
    output=$(go test -cover "$pkg" 2>&1)

    # Extract the coverage line
    if echo "$output" | grep -q "coverage:"; then
        coverage=$(echo "$output" | grep "coverage:" | sed -E 's/.*coverage: ([0-9.]+)% of statements/\1%/')
        echo "$pkg: $coverage"
    else
        echo "$pkg: no coverage info"
    fi
done