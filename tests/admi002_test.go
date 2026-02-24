package tests

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/moov-io/fednow20022/gen/admi_002_001_01"

	"github.com/stretchr/testify/require"
)

func TestAdmi002(t *testing.T) {
	doc := admi_002_001_01.Document{
		Admi00200101: admi_002_001_01.Admi00200101{
			RltdRef: admi_002_001_01.MessageReference{
				Ref: admi_002_001_01.Max35Text("test"),
			},
		},
	}

	bs, err := xml.Marshal(doc)
	require.NoError(t, err)

	if testing.Verbose() {
		fmt.Printf("\n\n%s\n", string(bs))
	}

	require.Contains(t, string(bs), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.002.001.01">`)
	require.Contains(t, string(bs), "<admi.002.001.01><RltdRef><Ref>test</Ref>")
}
