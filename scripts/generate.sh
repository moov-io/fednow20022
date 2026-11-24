#!/bin/bash
set -e

# The generator versions are pinned so that this script makes the committed
# gen/ code. Keep them the same as in the makefile. Update them only together
# with a regeneration.
XSD2GO_VERSION="${XSD2GO_VERSION:-v0.0.0-20261021193221-ef1732ddf011}"
GOIMPORTS_VERSION="${GOIMPORTS_VERSION:-v0.53.0}"
xsd2go="go run github.com/moov-io/xsd2go/cli/moovio_xsd2go@${XSD2GO_VERSION}"
goimports="go run golang.org/x/tools/cmd/goimports@${GOIMPORTS_VERSION}"

files=($(find ./xsd -name "*.xsd" | sort -u))
for file in "${files[@]}"
do
    $xsd2go convert "$file" github.com/moov-io/fednow20022 gen
done

# Use our custom Time types
go run ./scripts/fixup_imports.go

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
gofmt -w "${files[@]}"
$goimports -w "${files[@]}"

# Set the XML element names that moovio_xsd2go writes as ",any"
go run ./scripts/fixup_xmlnames.go
gofmt -w ./gen
