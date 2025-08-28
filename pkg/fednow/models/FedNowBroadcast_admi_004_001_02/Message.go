// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: FedNowBroadcast_admi.004.001.02
// Base Message: admi.004.001.02
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AJZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package FedNowBroadcast_admi_004_001_02

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/admi_004_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	EventCode        common.FundEventType `json:"EventCode,omitempty"`
	EventParameters  []string             `json:"EventParameter,omitempty"`
	EventDescription string               `json:"EventDescription,omitempty"`
	EventTime        time.Time            `json:"EventTime,omitzero"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &admi_004_001_02.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"EventCode", "EventTime",
}
var PathMap = map[string]any{
	"SysEvtNtfctn.EvtInf.EvtCd": "EventCode",
	"SysEvtNtfctn.EvtInf.EvtParam : EventParameters": map[string]any{
		"index": "index",
	},
	"SysEvtNtfctn.EvtInf.EvtDesc": "EventDescription",
	"SysEvtNtfctn.EvtInf.EvtTm":   "EventTime",
}
