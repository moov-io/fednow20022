#!/bin/bash
set -e

files=($(find ./xsd -name "*.xsd" | sort -u))
for file in "${files[@]}"
do
    moovio_xsd2go convert "$file" github.com/moov-io/fednow20022 gen
done

# Use our custom Time types
go run ./scripts/fixup_imports.go

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
for file in "${files[@]}"
do
    gofmt -w $file
    goimports -w $file
done
