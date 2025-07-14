// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: ReceiptAcknowledgement_admi.007.001.01
// Base Message: admi.007.001.01
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AIZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package ReceiptAcknowledgement_admi_007_001_01

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/ReceiptAcknowledgement_admi_007_001_01"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId           string                     `json:"message_id,omitempty"`
	CreatedDateTime     time.Time                  `json:"created_date_time,omitempty"`
	RelatedReference    string                     `json:"related_reference,omitempty"`
	MessageName         string                     `json:"message_name,omitempty"`
	RequestHandlingCode common.MessageHandlingCode `json:"request_handling_code,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &ReceiptAcknowledgement_admi_007_001_01.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
	"RelatedReference",
	"MessageName",
	"RequestHandlingCode",
}
var PathMap = map[string]any{
	"RctAck.MsgId.MsgId":       "MessageId",
	"RctAck.MsgId.CreDtTm":     "CreatedDateTime",
	"RctAck.Rpt.RltdRef.Ref":   "RelatedReference",
	"RctAck.Rpt.RltdRef.MsgNm": "MessageName",
	"RctAck.Rpt.ReqHdlg.StsCd": "RequestHandlingCode",
}
