// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: RetrievalRequest_admi.006.001.01
// Base Message: admi.006.001.01
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7AZZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package RetrievalRequest_admi_006_001_01

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/RetrievalRequest_admi_006_001_01"
	"github.com/moov-io/fednow20022/pkg/fednow"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type ResendSearchCriteria struct {
	OriginalMessageNameId string         `json:"original_message_name_id,omitempty"` // Original Message Identification
	BusinessDate          fednow.ISODate `json:"business_date,omitempty"`            // Business Date
	FileReference         string         `json:"file_reference,omitempty"`           // File Reference
	RecipientId           string         `json:"recipient_id,omitempty"`             // Recipient Identification
	RecipientIssuer       string         `json:"recipient_issuer,omitempty"`         // Recipient Issuer
}

type MessageModel struct {
	MessageId       string                 `json:"message_id,omitempty"`
	CreatedDateTime time.Time              `json:"created_date_time,omitempty"`
	Criteriaes      []ResendSearchCriteria `json:"criteriaes,omitempty"`
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:admi.006.001.01"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &RetrievalRequest_admi_006_001_01.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId",
	"CreatedDateTime",
}
var PathMap = map[string]any{
	"RsndReq.MsgHdr.MsgId":   "MessageId",
	"RsndReq.MsgHdr.CreDtTm": "CreatedDateTime",
	"RsndReq.RsndSchCrit : Criteriaes": map[string]any{
		"BizDt":                "BusinessDate",
		"OrgnlMsgNmId":         "OriginalMessageNameId",
		"FileRef":              "FileReference",
		"Rcpt.Id.PrtryId.Id":   "RecipientId",
		"Rcpt.Id.PrtryId.Issr": "RecipientIssuer",
	},
}
