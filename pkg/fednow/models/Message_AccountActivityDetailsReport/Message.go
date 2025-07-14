package Message_AccountActivityDetailsReport

import (
	"encoding/xml"

	"github.com/moov-io/fednow20022/pkg/fednow/models/AccountActivityDetailsReport_camt_052_001_08"
	"github.com/moov-io/fednow20022/pkg/fednow/models/head_001_001_02"
)

type MessageModel struct {
	AppHdr                       head_001_001_02.MessageModel                              `json:"app_hdr,omitempty"`
	AccountActivityDetailsReport AccountActivityDetailsReport_camt_052_001_08.MessageModel `json:"data,omitempty"`
}

var XMLAttributes = []xml.Attr{
	{Name: xml.Name{Local: "xmlns:head"}, Value: "urn:iso:std:iso:20022:tech:xsd:head.001.001.02"},
	{Name: xml.Name{Local: "xmlns:aadr"}, Value: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"},
	{Name: xml.Name{Local: "xmlns"}, Value: "urn:fednow"},
}
var HeaderPathMap = head_001_001_02.PathMap
var HeaderDocumentFactory = head_001_001_02.DocumentFactory
var DataPathMap = AccountActivityDetailsReport_camt_052_001_08.PathMap
var DataDocumentFactory = AccountActivityDetailsReport_camt_052_001_08.DocumentFactory

var HeaderWrapper = &head_001_001_02.HeadWrapper{}
var DataWrapper = &AccountActivityDetailsReport_camt_052_001_08.AccountActivityDetailsReportWrapper{}
