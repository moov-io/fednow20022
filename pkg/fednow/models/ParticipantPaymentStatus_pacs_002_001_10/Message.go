// Group: FedNow Service
// Collection: FedNow Service Release 1 - Edition 08/06/2024 (LIVE)
// Usage Guideline: ParticipantPaymentStatus_pacs.002.001.10
// Base Message: pacs.002.001.10
// Date of publication: 06 August 2024
// URL: https://www2.swift.com/mystandards/#/mp/mx/_kijaUKEYEeuB-v2uhrZqSw/_rz7ASZ5FEeyNX8SYmSuaLA
// Description: FedNow Service Scope

package ParticipantPaymentStatus_pacs_002_001_10

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fednow20022/gen/ParticipantPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fednow20022/pkg/fednow/models/common"
)

type MessageModel struct {
	MessageId               string                            `json:"message_id,omitempty"`
	CreatedDateTime         time.Time                         `json:"created_date_time,omitempty"`
	OriginalGroupInfo       common.GroupInformation           `json:"original_group_info,omitempty"`       // Information about the original group
	TransactionInfo         common.TransactionDetailReference `json:"transaction_info,omitempty"`          // Information about the transaction
	TransactionStatus       common.TransactionStatusCode      `json:"transaction_status,omitempty"`        // Status of the transaction
	StatusReasonInformation common.Reason                     `json:"status_reason_information,omitempty"` // Reason for the status
	Agents                  common.RelatedAgents              `json:"agents,omitempty"`                    // Information about the agents involved in the transaction
}

var XLNS = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"
var DataFactory = func() any {
	return &MessageModel{}
}
var DocumentFactory = func() common.ISODocument {
	return &ParticipantPaymentStatus_pacs_002_001_10.Document{
		XMLName: xml.Name{Space: XLNS, Local: "Document"}}
}
var RequireFileds = []string{
	"MessageId", "CreatedDateTime",
	"OriginalGroupInfo.OriginalMessageIdentification",
	"OriginalGroupInfo.OriginalMessageNameIdentification",
	"OriginalGroupInfo.OriginalCreationDateTime",
	"TransactionStatus",
	"Agents.InstructingAgent.PaymentSysCode",
	"Agents.InstructingAgent.PaymentSysMemberId",
	"Agents.InstructedAgent.PaymentSysCode",
	"Agents.InstructedAgent.PaymentSysMemberId",
}
var PathMap = map[string]any{
	"FIToFIPmtStsRpt.GrpHdr.MsgId":                                            "MessageId",
	"FIToFIPmtStsRpt.GrpHdr.CreDtTm":                                          "CreatedDateTime",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId":                      "OriginalGroupInfo.OriginalMessageIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalGroupInfo.OriginalMessageNameIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalGroupInfo.OriginalCreationDateTime",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlInstrId":                                "TransactionInfo.InstructionIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlEndToEndId":                             "TransactionInfo.EndToEndIdentification",
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlUETR":                                   "TransactionInfo.UETR", // Assuming UETR is the same as OrgnlTxId
	"FIToFIPmtStsRpt.TxInfAndSts.OrgnlTxId":                                   "TransactionInfo.TransactionId",
	"FIToFIPmtStsRpt.TxInfAndSts.TxSts":                                       "TransactionStatus",
	"FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.Rsn.Cd":                            "StatusReasonInformation.Code",
	"FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf.AddtlInf":                          "StatusReasonInformation.Proprietary",
	"FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructingAgent.PaymentSysCode",
	"FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructingAgent.PaymentSysMemberId",
	"FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "Agents.InstructedAgent.PaymentSysCode",
	"FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "Agents.InstructedAgent.PaymentSysMemberId",
}
