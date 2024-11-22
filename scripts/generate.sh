#!/bin/bash
set -e

moovio_xsd2go convert \
   xsd/xmldsig-core-schema.xsd \
   github.com/moov-io/fednow20022 \
   gen \
   --template-name=internal/templates/fednow20022/xmldsig.tgo \
   --xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"

moovio_xsd2go convert \
   xsd/messages.xsd \
   github.com/moov-io/fednow20022 \
   gen \
   --template-name=internal/templates/fednow20022/model.tgo \
   --template-name=internal/templates/fednow20022/write.tgo \
   --template-name=internal/templates/fednow20022/validate.tgo \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02=admi_004_001_02" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08=pacs_008_001_08"

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
for file in "${files[@]}"
do
    gofmt -w $file
    goimports -w $file
done
