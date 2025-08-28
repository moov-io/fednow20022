// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: MessageReject_admi_002_001_01
// Base Message: admi.002.001.01
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AQZ5FEeyNX8SYmSuaLA!samples
// Description: FedNow Service Scope

package MessageReject_admi_002_001_01

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/admi_002_001_01"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	RelatedReference     string    `json:"RelatedReference,omitempty"`
	RejectingPartyReason string    `json:"RejectingPartyReason,omitempty"`
	RejectionDateTime    time.Time `json:"RejectionDateTime,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:admi.002.001.01"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &admi_002_001_01.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"RelatedReference", "RejectingPartyReason", "RejectionDateTime",
}
var PathMap = map[string]any{
	"Admi00200101.RltdRef.Ref":     "RelatedReference",
	"Admi00200101.Rsn.RjctgPtyRsn": "RejectingPartyReason",
	"Admi00200101.Rsn.RjctnDtTm":   "RejectionDateTime",
}
