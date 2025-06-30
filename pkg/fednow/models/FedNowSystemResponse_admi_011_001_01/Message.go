// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: FedNowSystemResponse_admi.011.001.01
// Base Message: admi.011.001.01
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AMZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope
package FedNowSystemResponse_admi_011_001_01

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/FedNowSystemResponse_admi_011_001_01"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId      string               `json:"message_id,omitempty"`
	EventCode      common.FundEventType `json:"EventCode,omitempty"`
	EventParameter string               `json:"EventParameter,omitempty"`
	EventTime      time.Time            `json:"EventTime,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &FedNowSystemResponse_admi_011_001_01.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId", "EventCode", "EventTime",
}
var PathMap = map[string]any{
	"SysEvtAck.MsgId":            "MessageId",
	"SysEvtAck.AckDtls.EvtCd":    "EventCode",
	"SysEvtAck.AckDtls.EvtParam": "EventParameter",
	"SysEvtAck.AckDtls.EvtTm":    "EventTime",
}
