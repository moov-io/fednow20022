package tests

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/moov-io/fednow20022/gen/fednow_incoming_external"
	"github.com/moov-io/fednow20022/gen/head_001_001_02"
	"github.com/moov-io/fednow20022/gen/pacs_002_001_10"

	"github.com/stretchr/testify/require"
)

func TestFedNowIncomingEnvelopeMarshalXML(t *testing.T) {
	msg := fednow_incoming_external.FedNowIncoming{
		FedNowIncomingMessage: fednow_incoming_external.FedNowIncomingMessage{
			FedNowPaymentStatus: &fednow_incoming_external.FedNowPaymentStatus{
				AppHdr: head_001_001_02.BusinessApplicationHeaderV02{
					BizMsgIdr: head_001_001_02.Max35Text("test-msg-id"),
					MsgDefIdr: head_001_001_02.Max35Text("pacs.002.001.10"),
				},
				Document: pacs_002_001_10.Document{
					FIToFIPmtStsRpt: pacs_002_001_10.FIToFIPaymentStatusReportV10{
						GrpHdr: pacs_002_001_10.GroupHeader91{
							MsgId: pacs_002_001_10.Max35Text("test-msg-id"),
						},
					},
				},
			},
		},
	}

	bs, err := xml.Marshal(msg)
	require.NoError(t, err)

	if testing.Verbose() {
		fmt.Printf("\n\n%s\n", string(bs))
	}

	require.Contains(t, string(bs), `<FedNowIncoming xmlns="urn:fednow:incoming:v001">`)
	require.Contains(t, string(bs), `<AppHdr xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.02">`)
	require.Contains(t, string(bs), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10">`)
}
