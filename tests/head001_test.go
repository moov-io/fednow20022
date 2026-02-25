package tests

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/moov-io/fednow20022/gen/head_001_001_02"

	"github.com/stretchr/testify/require"
)

func TestHead001MarshalXML(t *testing.T) {
	hdr := head_001_001_02.BusinessApplicationHeaderV02{
		BizMsgIdr: head_001_001_02.Max35Text("test-msg-id"),
		MsgDefIdr: head_001_001_02.Max35Text("pacs.008.001.08"),
	}

	bs, err := xml.Marshal(hdr)
	require.NoError(t, err)

	if testing.Verbose() {
		fmt.Printf("\n\n%s\n", string(bs))
	}

	require.Contains(t, string(bs), `<AppHdr xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.02">`)
	require.Contains(t, string(bs), "<BizMsgIdr>test-msg-id</BizMsgIdr>")
}
