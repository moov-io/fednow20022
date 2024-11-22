#!/bin/bash
set -e

moovio_xsd2go convert \
   xsd/xmldsig-core-schema.xsd \
   github.com/moov-io/fednow20022 \
   gen \
   --template-name=internal/templates/fednow20022/xmldsig.tgo \
   --xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"


files=($(find ./xsd -not -name "FedNow_*" -name "*iso15.xsd" | sort -u))
for file in "${files[@]}"
do
    # file = "./xsd/AccountActivityDetailsReport_camt_052_001_08_20241122_1718_iso15.xsd"
    ns_under=$(echo "$file" | grep -oE '[a-z]{4}_[0-9]{3}_[0-9]{3}_[0-9]{2}')
    ns_dots=$(echo "$ns_under" | tr '_' '.')
    msg=$(basename "$file" .xsd | grep -oE '[a-zA-Z]{1,}_[a-z]{4}_[0-9]{3}_[0-9]{3}_[0-9]{2}')

    moovio_xsd2go convert \
                  "$file" github.com/moov-io/fednow20022 gen \
                  --template-name=internal/templates/fednow20022/model.tgo \
                  --template-name=internal/templates/fednow20022/write.tgo \
                  --template-name=internal/templates/fednow20022/validate.tgo \
                  --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
                  --xmlns-override="urn:iso:std:iso:20022:tech:xsd:${ns_dots}=${msg}" \
                  --xmlns-override="urn:fednow=Document_${msg}"
done



# moovio_xsd2go convert \
#    xsd/AccountActivityDetailsReport_camt_052_001_08_20241122_1718_iso15.xsd \
#    github.com/moov-io/fednow20022 \
#    gen \
#    --template-name=internal/templates/fednow20022/model.tgo \
#    --template-name=internal/templates/fednow20022/write.tgo \
#    --template-name=internal/templates/fednow20022/validate.tgo \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08=camt_052_001_08" \
#    --xmlns-override="urn:fednow=AccountActivityDetailsReport_camt_052_001_08"

# moovio_xsd2go convert \
#    xsd/CustomerCreditTransfer_pacs_008_001_08_20241122_1718_iso15.xsd \
#    github.com/moov-io/fednow20022 \
#    gen \
#    --template-name=internal/templates/fednow20022/model.tgo \
#    --template-name=internal/templates/fednow20022/write.tgo \
#    --template-name=internal/templates/fednow20022/validate.tgo \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08=pacs_008_001_08" \
#    --xmlns-override="urn:fednow=CustomerCreditTransfer_pacs_008_001_08"

# moovio_xsd2go convert \
#    xsd/messages.xsd \
#    github.com/moov-io/fednow20022 \
#    gen \
#    --template-name=internal/templates/fednow20022/model.tgo \
#    --template-name=internal/templates/fednow20022/write.tgo \
#    --template-name=internal/templates/fednow20022/validate.tgo \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.02=head_001_001_02" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08=camt_052_001_08__1" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08=camt_052_001_08__2" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08=camt_052_001_08__3" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.054.001.08=camt_054_001_08" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.060.001.05=camt_060_001_05" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.028.001.09=camt_028_001_09" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08=pacs_008_001_08" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02=admi_004_001_02__1" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.998.001.02=admi_998_001_02" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10=pacs_002_001_10__1" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.011.001.01=admi_011_001_01" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08=pacs_009_001_08" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.029.001.09=camt_029_001_09__1" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.026.001.07=camt_026_001_07" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.002.001.01=admi_002_001_01" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02=admi_004_001_02__2" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10=pacs_002_001_10__2" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10=pacs_004_001_10" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03=pacs_028_001_03" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.007.001.01=admi_007_001_01" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.029.001.09=camt_029_001_09__2" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.055.001.09=camt_055_001_09" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pain.014.001.07=pain_014_001_07" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pain.013.001.07=pain_013_001_07" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.006.001.01=admi_006_001_01" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.029.001.09=camt_029_001_09__3" \
#    --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.056.001.08=camt_056_001_08"

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
for file in "${files[@]}"
do
    gofmt -w $file
    goimports -w $file
done
