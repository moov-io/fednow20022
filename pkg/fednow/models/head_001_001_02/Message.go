// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: BusinessApplicationHeader_head.001.001.02
// Base Message: head.001.001.02
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_qR07WZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package head_001_001_02

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/head_001_001_02"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type Related struct {
	FromMemberId                string    `json:"from_member_id,omitempty"` // Identifier of the member sending
	ToMemberId                  string    `json:"to_member_id,omitempty"`   // Identifier of the member receiving
	BusinessMessageIdentifier   string    `json:"business_message_identifier,omitempty"`
	MessageDefinitionIdentifier string    `json:"message_definition_identifier,omitempty"` // Identifier of the message definition
	CreatedDateTime             time.Time `json:"created_date_time,omitempty"`             // Date and
}
type MessageModel struct {
	FromMemberId                string                   `json:"from_member_id,omitempty"`
	ToMemberId                  string                   `json:"to_member_id,omitempty"`
	CreatedDateTime             time.Time                `json:"created_date_time,omitempty"`
	BusinessMessageIdentifier   string                   `json:"business_message_identifier,omitempty"`
	MessageDefinitionIdentifier string                   `json:"message_definition_identifier,omitempty"`
	MarketPractice              common.MarketPractice    `json:"market_practice,omitempty"`
	BusinessProcessingDate      time.Time                `json:"business_processing_date,omitempty"`
	CopyDuplicate               common.CopyDuplicateCode `json:"copy_duplicate,omitempty"` // Indicates if the message is a copy or duplicate
	RelatedInfo                 Related                  `json:"related_info,omitempty"`   // Additional related information
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:head.001.001.02"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &head_001_001_02.AppHdr{
		XMLName: xml.Name{Space: XLNS, Local: "AppHdr"}}
}
var RequireFileds = []string{
	"FromMemberId",
	"ToMemberId",
	"BusinessMessageIdentifier",
	"MessageDefinitionIdentifier",
	"MarketPractice.Registry",
	"MarketPractice.Identification",
	"CreatedDateTime",
}
var PathMap = map[string]any{
	"Fr.FIId.FinInstnId.ClrSysMmbId.MmbId": "FromMemberId",
	"To.FIId.FinInstnId.ClrSysMmbId.MmbId": "ToMemberId",
	"BizMsgIdr":                            "BusinessMessageIdentifier",
	"MsgDefIdr":                            "MessageDefinitionIdentifier",
	"CreDt":                                "CreatedDateTime",
	"MktPrctc.Regy":                        "MarketPractice.Registry",
	"MktPrctc.Id":                          "MarketPractice.Identification",
	"BizPrcgDt":                            "BusinessProcessingDate",
	"CpyDplct":                             "CopyDuplicate",
	"Rltd.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId": "RelatedInfo.FromMemberId",
	"Rltd.To.FIId.FinInstnId.ClrSysMmbId.MmbId": "RelatedInfo.ToMemberId",
	"Rltd.BizMsgIdr": "RelatedInfo.BusinessMessageIdentifier",
	"Rltd.MsgDefIdr": "RelatedInfo.MessageDefinitionIdentifier",
	"Rltd.CreDt":     "RelatedInfo.CreatedDateTime",
}
